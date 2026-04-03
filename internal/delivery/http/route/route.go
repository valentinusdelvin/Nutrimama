package route

import (
	"nutrimama/internal/delivery/http"
	"nutrimama/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type RouteConfig struct {
	App                    *fiber.App
	UserController         *http.UserController
	EduToolsController     *http.EduToolsController
	TrackingController     *http.TrackingController
	NotificationController *http.NotificationController
	BookmarkController     *http.BookmarkController
	ConsultationController *http.ConsultationController
	ChatController         *http.ChatController
	MealPlanController     *http.MealPlanController
	UploadController       *http.UploadController
	PureConsultantController *http.ConsultantController
}

func (c *RouteConfig) Setup() {
	c.App.Static("/daris/uploads", "./public/uploads")

	SetupGuestRoutes(c)
	SetupProtectedRoutes(c)
}

func SetupGuestRoutes(c *RouteConfig) {
	route := c.App.Group("/api")
	
	route.Post("/upload", c.UploadController.UploadImage)
	
	route.Get("/consultants", c.PureConsultantController.List)
	route.Get("/notifications/vapid-public-key", c.NotificationController.GetVapidPublicKey)

	auth := route.Group("/auth")
	auth.Post("/register", c.UserController.Register)
	auth.Post("/login", c.UserController.Login)
	route.Get("/edu-tools", c.EduToolsController.List)
	route.Get("/edu-tools/:slug", c.EduToolsController.Get)
}

func SetupProtectedRoutes(c *RouteConfig) {
	route := c.App.Group("/api")
	route.Use(middleware.AuthMiddleware(c.UserController.UserUseCase))
	route.Get("/profile", c.UserController.Profile)
	route.Put("/profile/edit", c.UserController.EditProfile)

	edu := route.Group("/edu-tools")
	// Admin endpoint
	adminMiddleware := middleware.RoleMiddleware("admin")
	edu.Post("/", adminMiddleware, c.EduToolsController.Create)
	edu.Put("/:id", adminMiddleware, c.EduToolsController.Update)
	edu.Delete("/:id", adminMiddleware, c.EduToolsController.Delete)

	tracking := route.Group("/tracking")
	tracking.Post("/", c.TrackingController.SubmitTracking)
	tracking.Get("/", c.TrackingController.GetScores)
	tracking.Get("/questions", c.TrackingController.GetQuestions)

	notifications := route.Group("/notifications")
	notifications.Post("/subscribe", c.NotificationController.Subscribe)
	notifications.Get("/", c.NotificationController.GetList)
	notifications.Post("/test", c.NotificationController.SendTest)

	bookmarks := route.Group("/bookmarks")
	bookmarks.Post("/toggle", c.BookmarkController.ToggleBookmark)
	bookmarks.Get("/", c.BookmarkController.GetBookmarks)

	mealPlans := route.Group("/meal-plans")
	mealPlans.Post("/generate-ai", c.MealPlanController.GenerateAIMealPlan)
	mealPlans.Get("/", c.MealPlanController.GetMealPlan)

	consultations := route.Group("/consultations")
	consultations.Get("/", c.ConsultationController.GetInbox)
	consultations.Post("/book", c.ConsultationController.BookSession)
	consultations.Post("/:id/end", c.ConsultationController.EndSession)

	chat := route.Group("/chat")
	
	chat.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	
	chat.Get("/ws/:id", websocket.New(c.ChatController.HandleWebSocket))
	
	chat.Post("/send", c.ChatController.SendMessage)
	chat.Get("/history/:id", c.ChatController.GetHistory)
}
