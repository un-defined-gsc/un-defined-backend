package otp_serivce

import (
	"bytes"
	"image/png"

	"github.com/pquerna/otp"
	totp "github.com/pquerna/otp/totp"
)

type optService struct {
	issuer string
	period uint
}

func NewOTPService(issuer string, periodSecond uint) *optService {
	return &optService{
		issuer: issuer,
		period: periodSecond,
	}
}

func (t *optService) GenerateOTP(userEmail string) (otpURL string, err error) {
	otp, err := totp.Generate(totp.GenerateOpts{
		Issuer:      t.issuer,
		AccountName: userEmail,
		Period:      t.period,
	})
	if err != nil {
		return
	}
	otpURL = otp.URL()
	return
}

func (t *optService) ValidateOTP(inputKey string, otpURL string) (stat bool, err error) {
	key, err := otp.NewKeyFromURL(otpURL)
	if err != nil {
		return
	}
	stat = totp.Validate(inputKey, key.Secret())
	return
}

func (t *optService) GenerateImage(otpURL string) (image bytes.Buffer, err error) {
	key, err := otp.NewKeyFromURL(otpURL)
	if err != nil {
		return
	}
	img, err := key.Image(200, 200)
	if err != nil {
		return
	}
	err = png.Encode(&image, img)
	return
}
