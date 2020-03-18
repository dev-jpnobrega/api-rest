package main

// Settings representation configs application
type Settings struct {
	ConnectionString string
}

func (s *Settings) Init() {
	s.ConnectionString = "postgres://postgres:postgres@localhost/profile?sslmode=disable"
}
