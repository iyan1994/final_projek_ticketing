package controller

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func GenerateTicketNumber(db *gorm.DB) (string, error) {
	var lastTicket struct {
		TicketNumber string `gorm:"column:no_tiket"`
	}

	// Cari nomor tiket terakhir dari tabel
	if err := db.Table("tickets").
		Select("no_tiket").
		Order("id_ticket DESC"). // Ambil tiket terakhir berdasarkan ID
		Limit(1).
		Scan(&lastTicket).Error; err != nil && err != gorm.ErrRecordNotFound {
		return "", fmt.Errorf("failed to fetch last ticket: %v", err)
	}

	// Tentukan nomor tiket baru
	var newTicketNumber string
	if lastTicket.TicketNumber != "" {
		// Ambil nomor dari tiket terakhir
		lastNumber, err := strconv.Atoi(lastTicket.TicketNumber)
		if err != nil {
			return "", fmt.Errorf("failed to parse last ticket number: %v", err)
		}
		newTicketNumber = fmt.Sprintf("%09d", lastNumber+1) // Tambahkan 1 dan format dengan 9 digit
	} else {
		// Jika belum ada tiket, mulai dari 1
		newTicketNumber = "000000001"
	}

	return newTicketNumber, nil
}
