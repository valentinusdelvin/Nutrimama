package usecase

import (
	"errors"
	"fmt"
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/repository"
	"time"

	"gorm.io/gorm"
)

type ConsultationUseCase struct {
	DB             *gorm.DB
	SessionRepo    *repository.ConsultationSessionRepository
	MotherRepo     *repository.MotherRepository
}

func NewConsultationUseCase(db *gorm.DB, sRepo *repository.ConsultationSessionRepository, mRepo *repository.MotherRepository) *ConsultationUseCase {
	return &ConsultationUseCase{
		DB:          db,
		SessionRepo: sRepo,
		MotherRepo:  mRepo,
	}
}

type BookSessionRequest struct {
	ConsultantID  int    `json:"consultant_id"`
	StartTime     string `json:"start_time"` // Format: "HH:MM"
	DurationHours int    `json:"duration_hours"`
}

func (u *ConsultationUseCase) BookSession(userId int, req BookSessionRequest) (*entity.ConsultationSession, error) {
	c := u.DB.Begin()

	mother, err := u.MotherRepo.FindByUserId(c, userId)
	if err != nil {
		c.Rollback()
		return nil, errors.New("only registered mothers can book consultations")
	}

	var consultant entity.Consultant
	if err := c.Where("consultant_id = ?", req.ConsultantID).First(&consultant).Error; err != nil {
		c.Rollback()
		return nil, errors.New("consultant not found")
	}

	if req.DurationHours < 1 {
		c.Rollback()
		return nil, errors.New("duration must be at least 1 hour")
	}

	// Lock scheduling specifically to exactly today dynamically.
	now := time.Now()
	loc := time.FixedZone("Asia/Jakarta", 7*3600) // standardizing timezone logic safely
	today := now.In(loc)

	reqTimeStr := fmt.Sprintf("%04d-%02d-%02d %s", today.Year(), today.Month(), today.Day(), req.StartTime)
	layout := "2006-01-02 15:04"
	startTime, err := time.ParseInLocation(layout, reqTimeStr, loc)
	if err != nil {
		c.Rollback()
		return nil, errors.New("invalid start time format, expected exactly HH:MM")
	}

	// if startTime.Before(now) {
	// 	c.Rollback()
	// 	return nil, errors.New("cannot book a past schedule today")
	// }

	endTime := startTime.Add(time.Duration(req.DurationHours) * time.Hour)

	// if endTime.YearDay() != today.YearDay() || endTime.Year() != today.Year() {
	// 	c.Rollback()
	// 	return nil, errors.New("consultation sessions cannot overlap past midnight into tomorrow")
	// }

	timeStartStr := startTime.Format("15:04:05")
	hourEndStr := endTime.Format("15:04:05")
	todayDateStr := today.Format("2006-01-02")

	overlaps, err := u.SessionRepo.FindOverlapping(c, consultant.ConsultantID, todayDateStr, timeStartStr, hourEndStr)
	if err != nil {
		c.Rollback()
		return nil, err
	}
	if overlaps > 0 {
		c.Rollback()
		return nil, errors.New("this consultant is already booked during this exact timeframe")
	}

	totalPrice := consultant.HourlyRate * float64(req.DurationHours)

	session := entity.ConsultationSession{
		MotherID:      mother.MotherID,
		ConsultantID:  consultant.ConsultantID,
		SessionDate:   today,
		TimeStart:     timeStartStr,
		HourEnd:       hourEndStr,
		TotalPrice:    totalPrice,
		PaymentStatus: "paid", // TODO: PLACEHOLDER FOR FUTURE MIDTRANS/STRIPE PAYMENT MODULE LOGIC!
		Status:        "scheduled",
	}

	if err := u.SessionRepo.Create(c, &session); err != nil {
		c.Rollback()
		return nil, err
	}

	c.Commit()
	return &session, nil
}

func (u *ConsultationUseCase) EndSession(userId int, userRole string, consultationId int) (*entity.ConsultationSession, error) {
	c := u.DB

	var session entity.ConsultationSession
	if err := c.Where("consultation_id = ?", consultationId).First(&session).Error; err != nil {
		return nil, errors.New("consultation session not found")
	}

	if userRole == "mother" {
		var mother entity.Mother
		if err := c.Where("user_id = ?", userId).First(&mother).Error; err != nil {
			return nil, errors.New("unauthorized mapping mother structure securely")
		}
		if session.MotherID != mother.MotherID {
			return nil, errors.New("unauthorized to end this session")
		}
		session.Status = "ended"
		if err := c.Save(&session).Error; err != nil {
			return nil, err
		}
		return &session, nil

	} else if userRole == "consultant" {
		var consultant entity.Consultant
		if err := c.Where("user_id = ?", userId).First(&consultant).Error; err != nil {
			return nil, errors.New("unauthorized mapping consultant structure securely")
		}
		if session.ConsultantID != consultant.ConsultantID {
			return nil, errors.New("unauthorized to end this session")
		}
		
		now := time.Now()
		loc := time.FixedZone("Asia/Jakarta", 7*3600)
		today := now.In(loc)

		hourEndStr := fmt.Sprintf("%04d-%02d-%02d %s", session.SessionDate.Year(), session.SessionDate.Month(), session.SessionDate.Day(), session.HourEnd)
		layout := "2006-01-02 15:04:05" 
		
		endTime, err := time.ParseInLocation(layout, hourEndStr, loc)
		if err != nil {
			return nil, errors.New("failed to parse session end time limit natively")
		}

		if today.Before(endTime) {
			return nil, errors.New("consultants cannot terminate sessions before the scheduled time structurally expires")
		}

		session.Status = "ended"
		if err := c.Save(&session).Error; err != nil {
			return nil, err
		}
		return &session, nil

	}

	return nil, errors.New("invalid role mapping securely")
}

func (u *ConsultationUseCase) GetInboxList(userId int, userRole string) ([]model.InboxItemResponse, error) {
	var resolvedRoleID int
	c := u.DB

	if userRole == "mother" {
		var mother entity.Mother
		if err := c.Where("user_id = ?", userId).First(&mother).Error; err != nil {
			return nil, errors.New("unauthorized mapping mother structure securely")
		}
		resolvedRoleID = mother.MotherID
	} else if userRole == "consultant" {
		var consultant entity.Consultant
		if err := c.Where("user_id = ?", userId).First(&consultant).Error; err != nil {
			return nil, errors.New("unauthorized mapping consultant structure securely")
		}
		resolvedRoleID = consultant.ConsultantID
	} else {
		return nil, errors.New("invalid role mapping payload")
	}

	sessions, err := u.SessionRepo.FindByRole(c, userRole, resolvedRoleID)
	if err != nil {
		return nil, err
	}

	loc := time.FixedZone("Asia/Jakarta", 7*3600)
	now := time.Now().In(loc)

	var inboxList []model.InboxItemResponse
	for _, s := range sessions {
		displayStatus := s.Status
		if s.Status == "scheduled" {
			dateStr := s.SessionDate.Format("2006-01-02")
			layout := "2006-01-02 15:04:05"

			startDT, errS := time.ParseInLocation(layout, dateStr+" "+s.TimeStart, loc)
			endDT, errE := time.ParseInLocation(layout, dateStr+" "+s.HourEnd, loc)

			if errS == nil && errE == nil {
				if now.After(startDT) && now.Before(endDT) {
					displayStatus = "active"
				}
			}
		}

		item := model.InboxItemResponse{
			ConsultationID: s.ConsultationID,
			SessionDate:    s.SessionDate,
			TimeStart:      s.TimeStart,
			HourEnd:        s.HourEnd,
			Status:         displayStatus,
		}

		if userRole == "mother" {
			item.PartnerName = s.Consultant.FullName
			item.PartnerRole = "consultant"
		} else {
			item.PartnerName = s.Mother.FullName
			item.PartnerRole = "mother"
		}

		inboxList = append(inboxList, item)
	}

	return inboxList, nil
}

