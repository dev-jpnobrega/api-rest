package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// User representation persona
type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// IsValid check User is valid
func (p User) IsValid() bool {
	if len(strings.TrimSpace(p.Name)) > 0 && len(strings.TrimSpace(p.Email)) > 0 {
		return true
	}
	return false
}
