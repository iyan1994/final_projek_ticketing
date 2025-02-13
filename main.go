package main

import (
	"final_projek_ticketing/entity"
	"final_projek_ticketing/middleware"
	"final_projek_ticketing/repository"
	"final_projek_ticketing/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbOptions := os.Getenv("DB_OPTIONS")

	// Construct DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbOptions)

	repository.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	entity.MigrationTable()
	// Jalankan Seeder
	entity.SeedRoles(repository.Db)
	entity.SeedCategories(repository.Db)

	r := gin.Default()
	r.POST("/user", middleware.AuthMiddleware(repository.Db), service.CreateUserHandler)                                                         // semua role
	r.GET("/user/getclient", middleware.AdminAuthMiddleware(repository.Db), service.GetClientHandler)                                            // hanya admin
	r.GET("/user/getengineer", middleware.AdminAuthMiddleware(repository.Db), service.GetEngineerHandler)                                        // hanya admin
	r.DELETE("/user/delete/:id", middleware.AdminAuthMiddleware(repository.Db), service.DeleteUserHandler)                                       // hanya admin
	r.POST("/ticket", middleware.AuthMiddleware(repository.Db), service.CreateTicketHandler)                                                     // semua role
	r.POST("/assignticket", middleware.AdminAuthMiddleware(repository.Db), service.CreateAssignTicketHandler)                                    // hanya admin
	r.PUT("/assignticket/start/:id", middleware.EngineerAuthMiddleware(repository.Db), service.StartTicketHandler)                               // hanya engineer
	r.PUT("/assignticket/closed/:id", middleware.EngineerAuthMiddleware(repository.Db), service.ClosedTicketHandler)                             // hanya engineer
	r.POST("/assignticket/solution/:id", middleware.EngineerAuthMiddleware(repository.Db), service.UploadSolutionHandler)                        // hanya engineer
	r.GET("/assignticket/myassignticket", middleware.EngineerAuthMiddleware(repository.Db), service.ViewAssignTicketEngineerHandler)             // hanya engineer
	r.GET("/assignticket/myassignticketbyid/:id", middleware.EngineerAuthMiddleware(repository.Db), service.ViewAssignTicketEngineerByIdHandler) // hanya engineer

	r.POST("/login", service.LoginUserHandler)                                                                           // semua role
	r.GET("/ticket/myticket", middleware.AuthMiddleware(repository.Db), service.ViewMyTicketHandler)                     // semua role
	r.GET("/ticket/allticket", middleware.AdminAuthMiddleware(repository.Db), service.ViewAllTicketHandler)              // semua role
	r.GET("/ticket/viewticketbyid/:id", middleware.AuthMiddleware(repository.Db), service.ViewTicketByIdHandler)         // semua role
	r.POST("/ticket/uploadimageticket/:id", middleware.AuthMiddleware(repository.Db), service.UploadImagesTicketHandler) // semua role
	r.POST("/ticket/feedback/:id", middleware.AuthMiddleware(repository.Db), service.CreateFeedbackHandler)              // semua role
	r.GET("/ticket/viewfeedback/:id", middleware.AuthMiddleware(repository.Db), service.ViewFeedbackByIdTicket)          // semua role
	r.GET("/ticket/viewsolution/:id", middleware.AuthMiddleware(repository.Db), service.ViewSolutionByIdHandler)         // semua role
	r.GET("/ticket/viewticketimage/:id", middleware.AuthMiddleware(repository.Db), service.ViewImageTicketByIdTicket)    // semua role

	r.GET("/ticket/image/:filename", service.DownloadImageTicket)     // tidak ada role
	r.GET("/solution/image/:filename", service.DownloadImageSolution) // tidak ada role

	err = r.Run(":8083")
	if err != nil {
		log.Fatalln(err)
		return
	}

}
