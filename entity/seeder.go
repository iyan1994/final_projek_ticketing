package entity

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// SeedRoles menambahkan data awal ke tabel roles
func SeedRoles(db *gorm.DB) {
	roles := []Role{
		{IdRole: 1, NameRole: "admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{IdRole: 2, NameRole: "client", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{IdRole: 3, NameRole: "engineer", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, role := range roles {
		var existingRole Role
		result := db.Where("id_role = ?", role.IdRole).First(&existingRole)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				db.Create(&role)
				log.Printf("Role '%s' ditambahkan\n", role.NameRole)
			} else {
				log.Printf("Error cek role: %v\n", result.Error)
			}
		}
	}
}

// SeedRoles menambahkan data awal ke tabel categori
func SeedCategories(db *gorm.DB) {
	categoris := []Category{
		{IdCategory: 1, NameCategory: "Request", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{IdCategory: 2, NameCategory: "Problem", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{IdCategory: 3, NameCategory: "Corective", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, categori := range categoris {
		var existingCategori Category
		result := db.Where("id_category = ?", categori.IdCategory).First(&existingCategori)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				db.Create(&categori)
				log.Printf("Categoris '%s' ditambahkan\n", categori.NameCategory)
			} else {
				log.Printf("Error cek Categoris: %v\n", result.Error)
			}
		}
	}
}
