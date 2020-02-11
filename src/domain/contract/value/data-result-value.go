package domain

// ResponseData ...
type ResponseData struct {
	Data       interface{}
	StatusCode int
}

// RequestData ...
type RequestData struct {
	Authorization string
	XAppToken     string
	UserInfo      interface{}
	Args          interface{}
}

// InputLogin - parameters login user
type InputLogin struct {
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Pass  string `json:"pass" form:"pass" query:"pass" validate:"required"`
}
