package deps_services

import deps_ports "github.com/un-defined-gsc/un-defined-backend/internal/core/ports/deps"

type depsServices struct {
	// captchaService   deps_ports.ICaptchaService
	hasherService    deps_ports.IHasherService
	mailerService    deps_ports.IEmailService
	otpService       deps_ports.IOTPService
	validatorService deps_ports.IValidatorService
	cencorService    deps_ports.CensorService
	storageservice   deps_ports.IStorageService
}

func NewDepsServices(
	// captchaService deps_ports.ICaptchaService,
	hasherService deps_ports.IHasherService,
	mailerService deps_ports.IEmailService,
	otpService deps_ports.IOTPService,
	validatorService deps_ports.IValidatorService,
	cencorService deps_ports.CensorService,
	storageservice deps_ports.IStorageService,
) deps_ports.IDepsServices {
	return &depsServices{
		// captchaService:   captchaService,
		hasherService:    hasherService,
		mailerService:    mailerService,
		otpService:       otpService,
		validatorService: validatorService,
		cencorService:    cencorService,
		storageservice:   storageservice,
	}
}

// func (s *depsServices) CaptchaService() deps_ports.ICaptchaService {
// 	return s.captchaService
// }

func (s *depsServices) HasherService() deps_ports.IHasherService {
	return s.hasherService
}

func (s *depsServices) MailService() deps_ports.IEmailService {
	return s.mailerService
}

func (s *depsServices) OTPService() deps_ports.IOTPService {
	return s.otpService
}

func (s *depsServices) ValidatorService() deps_ports.IValidatorService {
	return s.validatorService
}
func (s *depsServices) CensorService() deps_ports.CensorService {
	return s.cencorService
}

func (s *depsServices) StorageService() deps_ports.IStorageService {
	return s.storageservice
}
