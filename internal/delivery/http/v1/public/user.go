package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PublicHandler) initUserRoutes(root fiber.Router) {
	root.Post("/login", h.Login)
	root.Post("/register", h.Register)
	root.Post("/recover/send", h.RecoverSend)
	root.Get("/verify/first/:key<alpha>", h.VerifyFirst)
	root.Get("/verify/email/:key<alpha>", h.VerifyEmail)
	root.Post("/recover/:key<alpha>", h.RecoverPassword)

}

// @Tags Auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param login body domains.LoginDTO true "Login"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/login [post]
func (h *PublicHandler) Login(c *fiber.Ctx) error {
	var login domains.LoginDTO
	if err := c.BodyParser(&login); err != nil {
		return err
	}
	resp, err := h.adapter.UsersServices().UsersService().Login(c.Context(), login)
	if err != nil {
		return err
	}
	sess, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	sess.Set("user", resp)
	if err := sess.Save(); err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth
// @Summary Register
// @Description Register
// @Accept json
// @Produce json
// @Param register body domains.RegisterDTO true "Register"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/register [post]
func (h *PublicHandler) Register(c *fiber.Ctx) error {
	var register domains.RegisterDTO
	if err := c.BodyParser(&register); err != nil {
		return err
	}
	err := h.adapter.UsersServices().UsersService().Register(c.Context(), register)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth Verify
// @Summary Verify First
// @Description Kullanıcı ilk kayıt olduğunda gelen maildeki linki tıkladığında çalışır.
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/verify/first/{key} [get]
func (h *PublicHandler) VerifyFirst(c *fiber.Ctx) error {
	key := c.Params("key")
	err := h.adapter.UsersServices().UsersService().VerifyProfile(c.Context(), key)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth Verify
// @Summary Verify Email
// @Description Kullanıcı mailini değiştirdiğinde gelen maildeki linki tıkladığında çalışır.
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/verify/email/{key} [get]
func (h *PublicHandler) VerifyEmail(c *fiber.Ctx) error {
	key := c.Params("key")
	err := h.adapter.UsersServices().UsersService().VerifyEmail(c.Context(), key)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth Recover
// @Summary Recover Send
// @Description Kullanıcı şifresini unuttuğunda çalışır.
// @Accept json
// @Produce json
// @Param email body domains.PasswordRecoverySendDTO true "Email"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/recover/send [post]
func (h *PublicHandler) RecoverSend(c *fiber.Ctx) error {
	var email domains.PasswordRecoverySendDTO
	if err := c.BodyParser(&email); err != nil {
		return err
	}
	err := h.adapter.UsersServices().UsersService().SendRecoveryToken(c.Context(), email)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth Recover
// @Summary Recover Password
// @Description Kullanıcı şifresini unuttuğunda çalışır.
// @Accept json
// @Produce json
// @Param key path string true "Key"
// @Param password body domains.PasswordRecoveryDTO true "Password"
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /public/recover/{key} [post]
func (h *PublicHandler) RecoverPassword(c *fiber.Ctx) error {
	key := c.Params("key")
	var password domains.PasswordRecoveryDTO
	if err := c.BodyParser(&password); err != nil {
		return err
	}
	password.Token = key
	err := h.adapter.UsersServices().UsersService().RecoverPassword(c.Context(), password)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}
