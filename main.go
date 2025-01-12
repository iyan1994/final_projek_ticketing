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
	dsn := "root:Lpkia@12345@tcp(127.0.0.1:3307)/db_ticketing?charset=utf8mb4&parseTime=true&loc=Local"

	var err error
	repository.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}

	entity.MigrationTable()

	r := gin.Default()
	r.POST("/user", middleware.AdminAuthMiddleware(repository.Db), service.CreateUserHandler) // hanya admin
	r.POST("/ticket", middleware.AuthMiddleware(repository.Db), service.CreateTicketHandler)  // semua role

	r.POST("/login", service.LoginUserHandler)

	err = r.Run(":8012")
	if err != nil {
		log.Fatalln(err)
		return
	}

}
