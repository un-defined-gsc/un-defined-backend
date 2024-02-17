package response_types

const (

	// HELLO
	HelloPrivateRoute = "routes.private.hello"
	HelloPublicRoute  = "routes.public.hello"
	HelloV1Route      = "routes.v1.hello"

	// VALIDATION
	ValidationErrors = "validation.errors"

	// InternalServer
	InternalServerError = "internal.server.error"

	// FIBER
	FiberNotFound = "fiber.not_found"
	ToManyRequest = "fiber.to_many_request"
	// Request
	ErrUnauthorized = "request.unauthorized"
	RequestSuccess  = "request.success"
	RequestFailed   = "request.failed"

	// USER
	MFARequried     = "user.mfa_requried"
	MailNotVerified = "user.mail_not_verified"

	CaptchaInvalid = "captcha.invalid"
)
