package session

import (
	"encoding/gob"

	"github.com/ProjectMonWeb/API-Service/internal/core/domains"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func NewSessionStore(storage ...fiber.Storage) *session.Store {
	// Bunun yeri burası mı bilmiyorum.
	gob.Register(domains.SessionDTO{})
	if len(storage) <= 0 {
		storage = append(storage, session.ConfigDefault.Storage)
	}
	return session.New(session.Config{
		CookieSecure:   true,
		CookieHTTPOnly: true,
		Storage:        storage[0],
	})
}
