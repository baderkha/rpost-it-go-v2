package controller

import (
	"rpost-it-go/pkg/api/service"
	"time"

	"github.com/gofiber/fiber"
)

// SessionAuth : Session based authentication used as the login controller
type SessionAuth struct {
	Basecontroller
	DefaultAuthDurationDays time.Duration
	DomainPrefix            string
}

func (s *SessionAuth) Login(ctx *fiber.Ctx) {
	if ctx.Cookies("session-token") != "" {
		s.OK(ctx, "login success")
		return
	}
	var loginCreds service.AccountLoginJSON
	err := ctx.BodyParser(&loginCreds)
	if err != nil {
		s.BodyCouldNotParse(ctx, err)
		return
	}
	session, err := s.Service().LoginSession(&loginCreds)
	if err != nil {
		s.StatusFromError(ctx, err)
		return
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "session-token",
		Value:    session.ID,
		Secure:   true,
		HTTPOnly: true,
		Domain:   "." + s.DomainPrefix,
		Expires:  time.Now().Add(s.DefaultAuthDurationDays * 24 * time.Hour),
	})
	s.OK(ctx, "login success")
}

func (s *SessionAuth) Logout(ctx *fiber.Ctx) {
	sessionId := ctx.Cookies("session-token")
	if sessionId == "" {
		s.OK(ctx, "logged out")
		return
	}
	s.Service().LogoutSession(sessionId)
	ctx.ClearCookie("session-token")
	s.OK(ctx, "logged out")
}

func (s *SessionAuth) Verify(ctx *fiber.Ctx) {
	sessionId := ctx.Cookies("session-token")
	session, err := s.service.VerifySession(sessionId)
	if err != nil {
		s.StatusFromError(ctx, err)
		return
	}
	// set to context
	ctx.Locals("request-account-id", session.AccountId)
	ctx.Next()
}
