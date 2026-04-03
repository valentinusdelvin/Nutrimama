package config

import (
	"nutrimama/internal/cron"
	"nutrimama/internal/delivery/http"
	"nutrimama/internal/delivery/http/route"
	"nutrimama/internal/repository"
	"nutrimama/internal/service"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB *gorm.DB
	App *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	
	userRepo := repository.NewUserRepository()
	motherRepo := repository.NewMotherRepository()
	consultantRepo := repository.NewConsultantRepository()
	
	userUseCase := usecase.NewUserUseCase(config.DB, userRepo, motherRepo, consultantRepo)
	userController := http.NewUserController(userUseCase)

	eduToolsRepo := repository.NewEduToolsRepository()
	eduToolsUseCase := usecase.NewEduToolsUseCase(config.DB, eduToolsRepo)
	eduToolsController := http.NewEduToolsController(eduToolsUseCase)

	trackingLogRepo := repository.NewTrackingLogRepository()
	trackingQuestionRepo := repository.NewTrackingQuestionRepository()
	nutritionTrackingRepo := repository.NewNutritionTrackingRepository()
	trackingUseCase := usecase.NewTrackingUseCase(config.DB, trackingLogRepo, trackingQuestionRepo, nutritionTrackingRepo, motherRepo)
	trackingController := http.NewTrackingController(trackingUseCase)

	pushService := service.NewPushService()

	notificationRepo := repository.NewNotificationRepository()
	pushSubscriptionRepo := repository.NewPushSubscriptionRepository()
	notificationUseCase := usecase.NewNotificationUseCase(config.DB, notificationRepo, pushSubscriptionRepo)
	notificationController := http.NewNotificationController(notificationUseCase, pushService)

	userBookmarkRepo := repository.NewUserBookmarkRepository()
	bookmarkUseCase := usecase.NewBookmarkUseCase(config.DB, userBookmarkRepo)
	bookmarkController := http.NewBookmarkController(bookmarkUseCase)

	consultationSessionRepo := repository.NewConsultationSessionRepository()
	consultationUseCase := usecase.NewConsultationUseCase(config.DB, consultationSessionRepo, motherRepo)
	consultationController := http.NewConsultationController(consultationUseCase)

	chatHub := utils.NewChatHub()
	messageRepo := repository.NewMessageRepository()
	chatUseCase := usecase.NewChatUseCase(config.DB, messageRepo, consultationSessionRepo)
	chatController := http.NewChatController(chatUseCase, chatHub)

	foodRepo := repository.NewFoodRepository()
	mealPlanUseCase := usecase.NewMealPlanUseCase(config.DB, foodRepo, nutritionTrackingRepo, motherRepo)
	mealPlanController := http.NewMealPlanController(mealPlanUseCase)

	trackingCron := cron.NewTrackingReminderCron(config.DB, notificationRepo, pushSubscriptionRepo, pushService)
	trackingCron.Start()

	uploadController := http.NewUploadController()

	consultantUseCaseBridge := usecase.NewConsultantUseCase(config.DB, consultantRepo)
	pureConsultantController := http.NewConsultantController(consultantUseCaseBridge)

	routeConfig := &route.RouteConfig{
		App:                    config.App,
		UserController:         userController,
		EduToolsController:     eduToolsController,
		TrackingController:     trackingController,
		NotificationController: notificationController,
		BookmarkController:     bookmarkController,
		ConsultationController: consultationController,
		ChatController:         chatController,
		MealPlanController:     mealPlanController,
		UploadController:       uploadController,
		PureConsultantController: pureConsultantController,
	}
	routeConfig.Setup()
}