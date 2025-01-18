package main

import (
	"final_projek_ticketing/entity"
	"final_projek_ticketing/middleware"
	"final_projek_ticketing/repository"
	"final_projek_ticketing/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Data source name (DSN) untuk MySQL
	dsn := "root:Lpkia@12345@tcp(host.docker.internal:3307)/db_ticketing?charset=utf8mb4&parseTime=true&loc=Local"
	//dsn := "root:Lpkia@12345@tcp(127.0.0.1:3307)/db_ticketing?charset=utf8mb4&parseTime=true&loc=Local"

	var err error
	repository.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	entity.MigrationTable()

	r := gin.Default()
	r.POST("/user", middleware.AdminAuthMiddleware(repository.Db), service.CreateUserHandler)                                                    // hanya admin
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
