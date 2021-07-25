package controller

import (
	"rpost-it-go/internal/api/service"
	"time"

	"github.com/gofiber/fiber"
)

// SessionAuth : Session based authentication used as the login controller
type SessionAuth struct {
	Basecontroller
	DefaultAuthDurationDays time.Duration
	DomainPrefix            string
	IsLocalHost             bool
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
	cookie := &fiber.Cookie{
		Name:     "session-token",
		Value:    session.ID,
		HTTPOnly: true,
		Secure:   !s.IsLocalHost,
		Expires:  time.Now().Add(s.DefaultAuthDurationDays * 24 * time.Hour),
	}

	cookie.Domain = "." + s.DomainPrefix

	ctx.Cookie(cookie)
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
