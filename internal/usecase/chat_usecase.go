package usecase

import (
	"errors"
	"fmt"
	"nutrimama/internal/entity"
	"nutrimama/internal/repository"
	"time"

	"gorm.io/gorm"
)

type ChatUseCase struct {
	DB          *gorm.DB
	MessageRepo *repository.MessageRepository
	SessionRepo *repository.ConsultationSessionRepository
}

func NewChatUseCase(db *gorm.DB, mRepo *repository.MessageRepository, sRepo *repository.ConsultationSessionRepository) *ChatUseCase {
	return &ChatUseCase{
		DB:          db,
		MessageRepo: mRepo,
		SessionRepo: sRepo,
	}
}

func (u *ChatUseCase) ValidateSessionAccess(userId int, userRole string, consultationId int, requiresActive bool) (*entity.ConsultationSession, error) {
	var session entity.ConsultationSession
	if err := u.DB.Where("consultation_id = ?", consultationId).First(&session).Error; err != nil {
		return nil, errors.New("consultation session not found")
	}

	if userRole == "mother" && session.MotherID != userId {
		return nil, errors.New("unauthorized to access this chat room")
	}
	if userRole == "consultant" && session.ConsultantID != userId {
		return nil, errors.New("unauthorized to access this chat room")
	}

	if !requiresActive {
		return &session, nil
	}

	if session.Status == "ended" {
		return nil, errors.New("this consultation session is officially terminated")
	}
	if session.PaymentStatus != "paid" {
		return nil, errors.New("cannot initiate chat for an unpaid session")
	}

	now := time.Now()
	loc := time.FixedZone("Asia/Jakarta", 7*3600)
	today := now.In(loc)

	startStr := fmt.Sprintf("%04d-%02d-%02d %s", session.SessionDate.Year(), session.SessionDate.Month(), session.SessionDate.Day(), session.TimeStart)
	endStr := fmt.Sprintf("%04d-%02d-%02d %s", session.SessionDate.Year(), session.SessionDate.Month(), session.SessionDate.Day(), session.HourEnd)
	
	layout := "2006-01-02 15:04:05" 
	startTime, _ := time.ParseInLocation(layout, startStr, loc)
	endTime, _ := time.ParseInLocation(layout, endStr, loc)

	if today.Before(startTime) {
		return nil, errors.New("this session has not actively started yet")
	}
	if today.After(endTime) {
		return nil, errors.New("this session has organically expired based on the scheduled blocks")
	}

	return &session, nil
}

type SendMessageRequest struct {
	ConsultationID int    `json:"consultation_id"`
	Message        string `json:"message"`
}

func (u *ChatUseCase) SendMessage(userId int, userRole string, req SendMessageRequest) (*entity.Message, error) {
	_, err := u.ValidateSessionAccess(userId, userRole, req.ConsultationID, true)
	if err != nil {
		return nil, err
	}

	msg := entity.Message{
		ConsultationID: &req.ConsultationID,
		Message:        req.Message,
		Status:         "sent",
	}

	if userRole == "mother" {
		msg.MotherID = &userId
	} else {
		msg.ConsultantID = &userId
	}

	if err := u.MessageRepo.Create(u.DB, &msg); err != nil {
		return nil, err
	}

	// TODO: WEBSOCKET EVENT EMIT PLACEHOLDER
	// Upon successful REST DB Save, we will later inject WsHub.BroadcastToRoom(req.ConsultationID, msg) 
	// heavily optimizing realtime delivery flawlessly!

	return &msg, nil
}

func (u *ChatUseCase) GetHistory(userId int, userRole string, consultationId int) ([]entity.Message, error) {
	_, err := u.ValidateSessionAccess(userId, userRole, consultationId, false)
	if err != nil {
		return nil, err
	}

	return u.MessageRepo.FindByConsultationId(u.DB, consultationId)
}
