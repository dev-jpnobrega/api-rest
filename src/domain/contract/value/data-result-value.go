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

// InputLogin - parameters login user
type InputLogin struct {
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Pass  string `json:"pass" form:"pass" query:"pass"`
}
