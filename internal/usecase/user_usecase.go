package usecase

import (
	// "errors"

	"errors"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/model/converter"
	"nutrimama/internal/repository"
	"nutrimama/internal/utils"

	bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB                   *gorm.DB
	UserRepository       *repository.UserRepository
	MotherRepository     *repository.MotherRepository
	ConsultantRepository *repository.ConsultantRepository
}

func NewUserUseCase(
	db *gorm.DB, 
	userRepository *repository.UserRepository,
	motherRepository *repository.MotherRepository,
	consultantRepository *repository.ConsultantRepository,
) *UserUseCase {
	return &UserUseCase{
		DB:                   db,
		UserRepository:       userRepository,
		MotherRepository:     motherRepository,
		ConsultantRepository: consultantRepository,
	}
}

func (u *UserUseCase) Create(req model.RegisterRequest) (*model.UserResponse, error) {
	c := u.DB.Begin()
	existingUser, _ := u.UserRepository.FindByEmail(c, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)
	user := entity.User{
		Email:      req.Email,
		Password:   req.Password,
		PhoneNumber: req.PhoneNumber,
		Role:      req.Role,
	}

	if err := u.UserRepository.Create(c, &user); err != nil {
		return nil, err
	}

	if req.Role == "mother" {
		if err := u.MotherRepository.Create(c, &entity.Mother{UserID: user.ID, FullName: req.Name}); err != nil {
			return nil, err
		}
	} else if req.Role == "consultant" {
		if err := u.ConsultantRepository.Create(c, &entity.Consultant{UserID: user.ID, FullName: req.Name}); err != nil {
			return nil, err
		}
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	if err := c.Commit().Error; err != nil {
		return nil, err
	}

	return converter.UserToResponse(&user, token, req.Name), nil

}

func (u *UserUseCase) Login(email string, password string) (*model.UserResponse, error) {
	db := u.DB
	user, err := u.UserRepository.FindByEmail(db, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	var name string
	if user.Role == "mother" {
		mother, _ := u.MotherRepository.FindByUserId(db, user.ID)
		if mother != nil { name = mother.FullName }
	} else if user.Role == "consultant" {
		consultant, _ := u.ConsultantRepository.FindByUserId(db, user.ID)
		if consultant != nil { name = consultant.FullName }
	}

	return converter.UserToResponse(user, token, name), nil
}

func (u *UserUseCase) UpdateMotherProfile(userId int, req model.EditMotherProfileRequest) (*entity.Mother, error) {
	c := u.DB.Begin()

	mother, err := u.MotherRepository.FindByUserId(c, userId)
	isNew := false
	if err != nil {
		isNew = true
		mother = &entity.Mother{UserID: userId}
	}

	mother.FullName = req.FullName
	if req.BloodType != nil {
		mother.BloodType = req.BloodType
	}
	if req.BirthDate != nil {
		mother.BirthDate = req.BirthDate
	}
	if req.MedicalHistory != nil {
		mother.MedicalHistory = req.MedicalHistory
	}
	if req.Height != nil {
		mother.Height = req.Height
	}

	if isNew {
		err = u.MotherRepository.Create(c, mother)
	} else {
		err = u.MotherRepository.Update(c, mother)
	}

	if err != nil {
		c.Rollback()
		return nil, err
	}

	if err := c.Commit().Error; err != nil {
		return nil, err
	}

	return mother, nil
}

func (u *UserUseCase) UpdateConsultantProfile(userId int, req model.EditConsultantProfileRequest) (*entity.Consultant, error) {
	c := u.DB.Begin()

	consultant, err := u.ConsultantRepository.FindByUserId(c, userId)
	isNew := false
	if err != nil {
		isNew = true
		consultant = &entity.Consultant{UserID: userId}
	}

	consultant.FullName = req.FullName
	if req.FacilityName != nil {
		consultant.FacilityName = req.FacilityName
	}

	if isNew {
		err = u.ConsultantRepository.Create(c, consultant)
	} else {
		err = u.ConsultantRepository.Update(c, consultant)
	}

	if err != nil {
		c.Rollback()
		return nil, err
	}

	if err := c.Commit().Error; err != nil {
		return nil, err
	}

	return consultant, nil
}