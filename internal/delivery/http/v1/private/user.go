package private

import (
	"github.com/gofiber/fiber/v2"
	"github.com/un-defined-gsc/un-defined-backend/internal/core/domains"
	user_domain "github.com/un-defined-gsc/un-defined-backend/internal/core/domains/user"
	"github.com/un-defined-gsc/un-defined-backend/internal/delivery/http/response_types"
)

func (h *PrivateHandler) initUserRoutes(root fiber.Router) {
	user := root.Group("/user")
	user.Get("/enable/:key<int>", h.EnableSession)
	user.Get("/me", h.GetUserMe)
	user.Get("/me/mfa/qr", h.GetMFAQR)
	// user.Get("/me/email/resend", h.ResendVerifyToken) // Devre dışı bırakıldı şimdilik
	user.Put("/me", h.UpdateUserMe)
	user.Put("/me/password", h.ChangePassword)
	user.Put("/me/email", h.ChangeEmail)
	user.Put("/me/mfa/toggle", h.MFAToggle)
	user.Get("/logout", h.Logout)
}

// @Tags User
// @Summary Get user info
// @Description Get user info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} error_handler.BaseResponse{data=user_domain.User}
// @Router /private/user/me [get]
func (h *PrivateHandler) GetUserMe(c *fiber.Ctx) error {
	user := c.Locals("user").(domains.SessionDTO)
	userdb, err := h.coreAdapter.UsersServices().UsersService().GetMe(c.Context(), *user.UUID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, userdb)
}

// @Tags User
// @Summary Update user info
// @Description Update user info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body user_domain.User true "User"
// @Success 200 {object} error_handler.BaseResponse{data=user_domain.User}
// @Router /private/user/me [put]
func (h *PrivateHandler) UpdateUserMe(c *fiber.Ctx) error {
	var user user_domain.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	usersess := c.Locals("user").(domains.SessionDTO)
	user.UUID = usersess.UUID
	err := h.coreAdapter.UsersServices().UsersService().UpdateMe(c.Context(), user)
	if err != nil {
		return err
	}
	user.Email = usersess.Email
	user.EmailVerified = usersess.EmailVerified
	user.MFAEnabled = usersess.MFAEnabled
	user.LastLogin = usersess.LastLogin
	user.CreatedAt = usersess.CreatedAt
	return h.responseJson(200, response_types.RequestSuccess, user)
}

// @Tags User Password
// @Summary Change user password
// @Description Change user password
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param password body domains.PasswordChangeDTO true "Password"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/user/me/password [put]
func (h *PrivateHandler) ChangePassword(c *fiber.Ctx) error {
	var password domains.PasswordChangeDTO
	if err := c.BodyParser(&password); err != nil {
		return err
	}
	usersess := c.Locals("user").(domains.SessionDTO)
	password.UserId = *usersess.UUID
	err := h.coreAdapter.UsersServices().UsersService().ChangePassword(c.Context(), password)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags User Email
// @Summary Change user email
// @Description Change user email
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param email body domains.EmailCahangeDTO true "Email"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/user/me/email [put]
func (h *PrivateHandler) ChangeEmail(c *fiber.Ctx) error {
	var email domains.EmailCahangeDTO
	if err := c.BodyParser(&email); err != nil {
		return err
	}
	sess, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	usersess := sess.Get("user").(domains.SessionDTO)
	email.UserId = *usersess.UUID
	err = h.coreAdapter.UsersServices().UsersService().ChangeEmail(c.Context(), email)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// Devre dışı bırakıldı şimdilik
//--------------------------------------------
// // @Tags User Email
// // @Summary Resend verify token
// // @Description Resend verify token
// // @Accept json
// // @Produce json
// // @Security ApiKeyAuth
// // @Success 200 {object} error_handler.BaseResponse
// // @Router /private/user/me/email/resend [get]
// func (h *PrivateHandler) ResendVerifyToken(c *fiber.Ctx) error {
// 	usersess := c.Locals("user").(domains.SessionDTO)
// 	err := h.coreAdapter.UsersServices().UsersService().ResendVerifyEmail(c.Context(), domains.EmailDTO{
// 		UserId:    *usersess.UUID,
// 		Email:     usersess.Email,
// 		FirstName: usersess.FirstName,
// 		LastName:  usersess.LastName,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return h.responseJson(200, response_types.RequestSuccess, nil)
// }
//--------------------------------------------

// @Tags User MFA
// @Summary Get MFA QR
// @Description Get MFA QR
// @Accept json
// @Produce png
// @Security ApiKeyAuth
// @Success 200 {file} png
// @Router /private/user/me/mfa/qr [get]
func (h *PrivateHandler) GetMFAQR(c *fiber.Ctx) error {
	usersess := c.Locals("user").(domains.SessionDTO)
	mfa, err := h.coreAdapter.UsersServices().UsersService().GetMFA(c.Context(), *usersess.UUID)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "image/png")
	return c.SendStream(&mfa.KeyImage)
}

// @Tags User MFA
// @Summary Toggle MFA
// @Description Toggle MFA
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/user/me/mfa/toggle [put]
func (h *PrivateHandler) MFAToggle(c *fiber.Ctx) error {
	usersess := c.Locals("user").(domains.SessionDTO)
	status, err := h.coreAdapter.UsersServices().UsersService().MFAToggle(c.Context(), *usersess.UUID)
	if err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, map[string]bool{"mfa_status": status})
}

// @Tags Auth
// @Summary Enable Session
// @Description Enable MFA
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param key path int true "Key"
// @Success 200 {object} error_handler.BaseResponse
// @Router /private/user/enable/{key} [get]
func (h *PrivateHandler) EnableSession(c *fiber.Ctx) error {
	key := c.Params("key")
	sess, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	userses := c.Locals("user").(domains.SessionDTO)
	err = h.coreAdapter.UsersServices().UsersService().EnableSession(c.Context(), &userses, key)
	if err != nil {
		return err
	}
	sess.Set("user", &userses)
	if err := sess.Save(); err != nil {
		return err
	}
	return h.responseJson(200, response_types.RequestSuccess, nil)
}

// @Tags Auth
// @Summary Logout
// @Description Logout
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} error_handler.BaseResponse{}
// @Router /private/user/logout [get]
func (h *PrivateHandler) Logout(c *fiber.Ctx) error { // Service tarafında yapılacak bir şey varsa servisi yazılacak
	sess, err := h.session_store.Get(c)
	if err != nil {
		return err
	}
	sess.Destroy()
	return h.responseJson(200, response_types.RequestSuccess, nil)
}
