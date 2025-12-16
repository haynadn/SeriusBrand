package main

import (
	"log"
	"os"

	"seriusbrand/backend/internal/db"
	"seriusbrand/backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := db.Migrate(database); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize handlers
	h := handlers.NewHandler(database)

	// Seed default admin user (Task: Implement Admin Auth)
	h.SeedAdminUser()

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/login", h.Login) // New Login Endpoint
		api.POST("/orders", h.CreateOrder)
		api.POST("/orders/:id/proof", h.UploadPaymentProof)
		api.GET("/umkm/:slug", h.GetUmkmPage)
		api.GET("/umkm", h.ListUmkmPages)
	}

	// Admin routes (Protected by Middleware)
	admin := r.Group("/api/admin")
	admin.Use(handlers.AuthMiddleware()) // Apply Middleware
	{
		admin.GET("/orders", h.ListOrders)
		admin.PATCH("/orders/:id", h.UpdateOrderStatus)
		admin.POST("/umkm", h.CreateUmkmPage)
		admin.PUT("/umkm/:id", h.UpdateUmkmPage)
	}

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
