package main

import (
	"final_projek_ticketing/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Data source name (DSN) untuk MySQL
	dsn := "root:Lpkia@12345@tcp(127.0.0.1:3307)/db_ticketing?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi ke database:", err)
		return
	}
	// Migrasi  untuk role
	err = db.AutoMigrate(&models.Role{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Role:", err)
		return
	}

	// Melakukan migrasi otomatis untuk User dan Ticket
	// Migrasi  untuk User
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi User:", err)
		return
	}

	fmt.Println("Migrasi User selesai")

	// Migrasi  untuk kategori
	err = db.AutoMigrate(&models.Category{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi kategori:", err)
		return
	}

	// Migrasi  untuk Ticket
	err = db.AutoMigrate(&models.Ticket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Ticket:", err)
		return
	}

	// Migrasi  untuk feedback
	err = db.AutoMigrate(&models.Feedback{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi feedback:", err)
		return
	}

	// Migrasi  untuk image ticket
	err = db.AutoMigrate(&models.ImageTicket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi image ticket:", err)
		return
	}

	// Migrasi  untuk prioritas
	err = db.AutoMigrate(&models.Priority{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Prioritas:", err)
		return
	}

	// Migrasi  untuk assign ticket
	err = db.AutoMigrate(&models.AssignTicket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi assign ticket:", err)
		return
	}

	// Migrasi  untuk solution
	err = db.AutoMigrate(&models.Solution{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Solution:", err)
		return
	}

	fmt.Println("Migrasi  selesai")

}
