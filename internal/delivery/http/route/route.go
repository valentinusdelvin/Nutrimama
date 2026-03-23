package route

import (
	"nutrimama/internal/delivery/http"
	"nutrimama/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	UserController     *http.UserController
	EduToolsController *http.EduToolsController
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
		// GET endpoints accessible to all authenticated users
	route.Get("/", c.EduToolsController.List)
	route.Get("/:id", c.EduToolsController.Get)
}

func SetupProtectedRoutes(c *RouteConfig) {
	route := c.App.Group("/api")
	route.Use(middleware.AuthMiddleware(c.UserController.UserUseCase))
	route.Get("/profile", c.UserController.Profile)

	edu := route.Group("/edu-tools")
	// Admin only endpoints
	adminMiddleware := middleware.RoleMiddleware("admin")
	edu.Post("/", adminMiddleware, c.EduToolsController.Create)
	edu.Put("/:id", adminMiddleware, c.EduToolsController.Update)
	edu.Delete("/:id", adminMiddleware, c.EduToolsController.Delete)
}
