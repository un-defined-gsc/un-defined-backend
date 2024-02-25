package service_errors

import "net/http"

/*
- Hata kodları ve mesajları burada tanımlanır.
Kodlar 0000'den başlar ve 1000'er artar.
Kodların Açıklaması
  - 0000: Başarılı işlemler
  - 1000: Kullanıcı tarafı ile ilgili hatalar
 	 + 1001: Kullanıcı adı veya şifre hatalı
 	 + 1002: Kullanıcı deaktif
	 + 1003: Kullanıcı banlı

  - 2000: Sistem ile ilgili hatalar
 	 + 2001: MFA tipi hatalı

  - 3000: Veritabanı ile ilgili hatalar

  - 4000: Dış servisler ile ilgili hatalar
	  + 4001: MFA sistemi çalışmıyor


  - 5000: Diğer hatalar

Mesajlar istenilen dile göre burada tanımlanır.
Yada i18n paketi için key tanımlanır.
*/

var (

	// 1000: Kullanıcı tarafı ile ilgili hatalar 400
	ErrInvalidUsernameOrPassword = NewServiceErrorWithKey(1001, "invalid_username_or_password", http.StatusUnauthorized) // Kullanıcı adı veya şifre hatalı
	ErrUserIsDisabled            = NewServiceErrorWithKey(1002, "user_is_disabled", http.StatusUnauthorized)             // Kullanıcı deaktif
	ErrUserIsBanned              = NewServiceErrorWithKey(1003, "user_is_banned", http.StatusUnauthorized)               // Kullanıcı banlı
	ErrUserEmailNotVerified      = NewServiceErrorWithKey(1004, "user_email_not_verified", http.StatusUnauthorized)      // Kullanıcı emaili doğrulanmamış
	ErrTokenInvalid              = NewServiceErrorWithKey(1005, "token_invalid", http.StatusBadRequest)                  // Token geçersiz
	ErrToManyRequest             = NewServiceErrorWithKey(1006, "to_many_request", http.StatusTooManyRequests)           // Çok fazla istek
	ErrInvalidOldPassword        = NewServiceErrorWithKey(1007, "invalid_old_password", http.StatusBadRequest)           // Eski şifre hatalı
	ErrInvalidOTP                = NewServiceErrorWithKey(1008, "invalid_otp", http.StatusBadRequest)                    // OTP hatalı
	ErrMFANotEnabled             = NewServiceErrorWithKey(1009, "mfa_not_enabled", http.StatusPreconditionRequired)      // MFA aktif değil
	ErrMonitorNotFound           = NewServiceErrorWithKey(1010, "monitor_not_found", http.StatusNotFound)                // Monitor bulunamadı
	ErrMonitorAlreadyExists      = NewServiceErrorWithKey(1011, "monitor_already_exists", http.StatusConflict)           // Monitor zaten var
	ErrCaptchaInvalid            = NewServiceErrorWithKey(1012, "captcha_invalid", http.StatusBadRequest)
	ErrLikedAlreadyExists        = NewServiceErrorWithKey(1013, "like_already_exists", http.StatusConflict) // Monitor zaten var
	ErrLikeNotFound              = NewServiceErrorWithKey(1014, "like_not_found", http.StatusNotFound)      // Monitor zaten var
	// Captcha hatalı

	// 2000: Sistem ile ilgili hatalar 500
	ErrInvalidMFAType = NewServiceErrorWithKey(2001, "invalid_mfa_type") // MFA tipi hatalı

	// 3000: Veritabanı ile ilgili hatalar 500
	ErrDataNotFound    = NewServiceErrorWithKey(3001, "data_not_found")   // 3001 Veri bulunamadı
	ErrDataDuplication = NewServiceErrorWithKey(3002, "data_duplication") // 3002 Veri çoğaltılamaz

	// 4000: Dış servisler ile ilgili hatalar 500
	ErrMFANotWorking = NewServiceErrorWithKey(4001, "mfa_not_working") // MFA sistemi çalışmıyor
	ErrGRPCDataError = NewServiceErrorWithKey(4002, "grpc_data_error") // GRPC ile ilgili hata
	// 5000: Diğer hatalar

)
