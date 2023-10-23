package types

type PlatformType struct {
	Platform string `json:"platform" binding:"required"`
}

type LoginRequestType struct {
	Key      string `json:"key" binding:"required"`
	Password string `json:"password"`
	Platform string `json:"platform"`
}

type RegisterRequestType struct {
	Username    *string
	FirstName   string `binding:"required"`
	LastName    string `binding:"required"`
	Password    string `binding:"required"`
	Platform    string `binding:"required"`
	Email       *string
	PhoneNumber *string
}

type ValidateEmail2FAType struct {
	Email string `json:"email" binding:"required"`
	PlatformType
}

type GetOAuthTokenWithPhoneNumberType struct {
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	PlatformType
}
