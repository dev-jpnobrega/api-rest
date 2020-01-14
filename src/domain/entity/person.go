package domain

import "strings"

// Person representation persona
type Person struct {
	ID   string
	Name string
	Age  int
}

func (p Person) IsValid() bool {
	if len(strings.TrimSpace(p.Name)) > 0 {
		return true
	}
	return false
}
