package service_errors

/*
- Hata kodları ve mesajları burada tanımlanır.
- Kodlar 1000'den başlar ve 1000'er artar.
- Örnek Hatalar için predefined.go dosyasına bakınız.
*/
type ServiceError struct {
	Code    int
	Key     string
	Message string
	err     error
	CodeSec int
}

func (e *ServiceError) Error() string { // Error() fonksiyonu error interface'ini implemente eder. Bu sayede error olarak döndürülebilir.
	if e.err != nil { // Eğer err varsa err'i döndürür.
		return e.err.Error()
	}
	if e.Key != "" { // Key varsa key'i döndürür.
		return e.Key
	}
	return e.Message
}
func NewServiceErrorWithMessage(code int, message string) error {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}
func NewServiceErrorWithKey(code int, key string, codesec ...int) error {
	stat := 500
	if len(codesec) > 0 {
		stat = codesec[0]
	}
	return &ServiceError{
		Code:    code,
		Key:     key,
		CodeSec: stat,
	}
}
func NewServiceErrorWithError(code int, err error) error {
	return &ServiceError{
		Code: code,
		err:  err,
	}
}
