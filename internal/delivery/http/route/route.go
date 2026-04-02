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
}

func (c *RouteConfig) Setup() {
	SetupGuestRoutes(c)
	SetupProtectedRoutes(c)
}

func SetupGuestRoutes(c *RouteConfig) {
	route := c.App.Group("/api")
	auth := route.Group("/auth")
	auth.Post("/register", c.UserController.Register)
	auth.Post("/login", c.UserController.Login)
	route.Get("/", c.EduToolsController.List)
	route.Get("/:id", c.EduToolsController.Get)
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

	notifications := route.Group("/notifications")
	notifications.Post("/subscribe", c.NotificationController.Subscribe)
	notifications.Get("/", c.NotificationController.GetList)

	bookmarks := route.Group("/bookmarks")
	bookmarks.Post("/toggle", c.BookmarkController.ToggleBookmark)
	bookmarks.Get("/", c.BookmarkController.GetBookmarks)

	consultations := route.Group("/consultations")
	consultations.Post("/book", c.ConsultationController.BookSession)
	consultations.Post("/:id/end", c.ConsultationController.EndSession)

	chat := route.Group("/chat")
	
	// Native Upgrade Firewalls securely routing only socket traffic natively gracefully.
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
