package infrastructure

import (
	"github.com/jinzhu/gorm"
)

// IDB - ref (2)
type IDB interface {
	Connect(driver string, connectionString string) (*gorm.DB, error)
	Get() *gorm.DB
	GetModel(model interface{}) *gorm.DB
}
