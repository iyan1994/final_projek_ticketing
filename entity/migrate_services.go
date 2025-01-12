package entity

import (
	"final_projek_ticketing/repository"
	"fmt"
)

func MigrationTable() {

	var err error

	// Migrasi  untuk role
	err = repository.Db.AutoMigrate(&Role{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Role:", err)
		return
	}

	fmt.Println("Migrasi Role selesai")

	// Melakukan migrasi otomatis untuk User dan Ticket
	// Migrasi  untuk User
	err = repository.Db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi User:", err)
		return
	}

	fmt.Println("Migrasi User selesai")

	// Migrasi  untuk kategori
	err = repository.Db.AutoMigrate(&Category{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi kategori:", err)
		return
	}
	fmt.Println("Migrasi Category selesai")

	// Migrasi  untuk Ticket
	err = repository.Db.AutoMigrate(&Ticket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Ticket:", err)
		return
	}
	fmt.Println("Migrasi Ticket selesai")

	// Migrasi  untuk feedback
	err = repository.Db.AutoMigrate(&Feedback{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi feedback:", err)
		return
	}

	fmt.Println("Migrasi Feedback selesai")

	// Migrasi  untuk image ticket
	err = repository.Db.AutoMigrate(&ImageTicket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi image ticket:", err)
		return
	}

	fmt.Println("Migrasi Image Ticket selesai")

	// Migrasi  untuk prioritas
	err = repository.Db.AutoMigrate(&Priority{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Prioritas:", err)
		return
	}

	fmt.Println("Migrasi Prioritas selesai")

	// Migrasi  untuk assign ticket
	err = repository.Db.AutoMigrate(&AssignTicket{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi assign ticket:", err)
		return
	}

	fmt.Println("Migrasi Assign Ticket selesai")

	// Migrasi  untuk solution
	err = repository.Db.AutoMigrate(&Solution{})
	if err != nil {
		fmt.Println("Gagal melakukan migrasi Solution:", err)
		return
	}
	fmt.Println("Migrasi Solution selesai")
	fmt.Println("Migrasi  selesai")

}
