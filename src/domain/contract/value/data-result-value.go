package domain

// DataResult ...
type DataResult struct {
	Data       interface{}
	StatusCode int
}

// DataInput ...
type DataInput struct {
	Authorization string
	XAppToken     string
	UserInfo      interface{}
	Args          interface{}
}
