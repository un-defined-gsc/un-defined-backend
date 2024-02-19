package deps_ports

import (
	"bytes"
	"context"
)

type IDepsServices interface {
	HasherService() IHasherService
	ValidatorService() IValidatorService
	MailService() IEmailService
	OTPService() IOTPService
	CensorService() CensorService
	// CaptchaService() ICaptchaService
}

// type ICaptchaService interface {
// 	New() string
// 	GetImageBytes(id string) (buf bytes.Buffer, err error)
// 	Verify(id, value string) bool
// }

type IHasherService interface {
	HashPassword(password string) (hashedPassword string, err error)
	CompareHashAndPassword(hashedPassword string, password string) (ok bool, err error)
}

type IOTPService interface { // Geliştirilecek şuan biraz belirsiz duruyor
	GenerateOTP(userEmail string) (otpURL string, err error)
	ValidateOTP(inputKey string, otpURL string) (stat bool, err error)
	GenerateImage(otpURL string) (image bytes.Buffer, err error)
}
type IValidatorService interface {
	ValidateStruct(s interface{}) (err error)
}

type IEmailService interface {
	SendRegisterVerifyMail(ctx context.Context, name, surname, expdate, to, link string) (err error)
	SendMailChangeVerify(ctx context.Context, name, surname, expdate, to, link string) (err error)
	SendNotifyOldMail(ctx context.Context, name, surname, changedate, to string, changed bool) (err error)
	SendPasswordRecoveryMail(ctx context.Context, name, surname, expdate, to, link string) (err error)
	SendFeedbackSuccessMail(ctx context.Context, name, surname, feedback, created_at, to string) (err error) // Eklenecekler var
	SendMails(ctx context.Context, to []string, subject, body string) (err error)
}

type CensorService interface {
	CensorText(textAddrs ...*string) (err error)
}
