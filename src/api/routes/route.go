package routes

import (
	"go-fitbyte/src/pkg/activity"
	"go-fitbyte/src/pkg/auth"
	"go-fitbyte/src/pkg/user"
	"go-fitbyte/src/pkg/userfile"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, v *viper.Viper, db *gorm.DB, services Services) {
	// API v1 group
	api := app.Group("/api/v1")

	// Init JWT Manager (24 jam expired)
	jwtManager := auth.NewJWTManager(v.GetString("JWT_SECRET"), 24*time.Hour)

	AuthRouter(api, services.AuthService, jwtManager)
	ProfileRouter(api, services.ProfileService, jwtManager)
	UserfileRouter(api, services.ProfileService, services.UploadFileService, jwtManager)
	ActivityRouter(api, services.ActivityService, jwtManager)
	// Register all routers
	// AuthRouter(api, services.AuthService) // Future router
	// UserRouter(api, services.UserService) // Future router
	// --- Health check route for Kubernetes probes ---
	app.Get("/healthz", func(c *fiber.Ctx) error {
		sqlDB, err := db.DB() // get underlying *sql.DB from GORM
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Database connection error")
		}

		if err := sqlDB.Ping(); err != nil { // try pinging the DB
			return c.Status(fiber.StatusInternalServerError).SendString("Database not reachable")
		}

		return c.SendStatus(fiber.StatusOK) // 200 OK if DB is fine
	})
}

// Services struct holds all service dependencies
type Services struct {
	AuthService       auth.Service
	ActivityService   activity.Service
	ProfileService    user.Service
	UploadFileService userfile.Service
	// AuthService auth.Service // Future service
	// UserService user.Service // Future service
}
