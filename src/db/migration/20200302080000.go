package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

// BaseModel define
type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// MigraCreateTableUser - Create table User
func MigraCreateTableUser() (m *gormigrate.Migration) {
	m.ID = "20200302080000"
	m.Migrate = func(tx *gorm.DB) error {
		type User struct {
			BaseModel
			Name  string
			Email string
		}

		return tx.AutoMigrate(&User{}).Error
	}
	m.Rollback = func(tx *gorm.DB) error {
		return tx.DropTable("user").Error
	}

	return
}
