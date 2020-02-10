package domain

import "strings"

// User representation persona
type User struct {
	ID    string
	Name  string
	Email string
}

// IsValid check User is valid
func (p User) IsValid() bool {
	if len(strings.TrimSpace(p.Name)) > 0 && len(strings.TrimSpace(p.Email)) > 0 {
		return true
	}
	return false
}
