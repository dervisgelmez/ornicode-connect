package services

import (
	"fmt"
	core "ornicode-connect/initializers/packages"
	"ornicode-connect/src/model"
	types "ornicode-connect/src/types/request"
	"ornicode-connect/src/types/response"
	"ornicode-connect/src/utility"
	"time"
)

type UserService struct{}

func (UserService) Generate2FATokenForPhoneNumber(request types.GetOAuthTokenWithPhoneNumberType) *response.OAuthTokenResponse {
	otp := utility.GenerateDigitToken(6)
	expiredAt := time.Now().Add(time.Minute * time.Duration(3))
	platform := core.DB.First(&model.UserPlatform{Name: request.Platform})
	fmt.Println(platform)

	core.DB.Create(model.UserOtp{
		Key:       request.PhoneNumber,
		Type:      "phoneNumber",
		Otp:       otp,
		ExpiredAt: &expiredAt,
	})

	responseType := new(response.OAuthTokenResponse)
	responseType.Token = otp
	responseType.ExpiredAt = expiredAt

	return responseType
}
