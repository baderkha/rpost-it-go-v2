package controller

import (
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/internal/api/service"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber"
)

// SessionAuth : Session based authentication used as the login controller
type SessionAuth struct {
	Basecontroller
	DefaultAuthDurationDays time.Duration
	DomainPrefix            string
	IsLocalHost             bool
}

func (s *SessionAuth) writeSessionToCookie(ctx *fiber.Ctx, session *repo.Session) {

	cookie := &fiber.Cookie{
		Name:     "session-token",
		Value:    session.ID,
		HTTPOnly: true,
		Secure:   !s.IsLocalHost,
		Expires:  session.ExpiryTime,
	}

	cookie.Domain = "." + s.DomainPrefix

	ctx.Cookie(cookie)
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
	s.writeSessionToCookie(ctx, session)
	s.OK(ctx, "login success")
}

// Logout : logs you out and deletes the session as well as the cookie
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

// Verify : ontop of the base implementation , the session also extends itself
// everytime a request is made with anything that has this middleware attached
func (s *SessionAuth) Verify(ctx *fiber.Ctx) {
	sessionId := ctx.Cookies("session-token")
	spew.Dump(sessionId)

	session, err := s.service.VerifySession(sessionId)
	spew.Dump(err)
	if err != nil {
		s.StatusFromError(ctx, err)
		return
	}

	s.refreshSession(ctx, session)
	// set to context
	ctx.Locals("request-account-id", session.AccountId)
	ctx.Next()
}

// refreshSession : refreshes the session for a token on the backened
func (s *SessionAuth) refreshSession(ctx *fiber.Ctx, session *repo.Session) {
	session, _ = s.service.RefreshSession(session) // keep refreshing your token , every time you hit a route endpoint
	ctx.ClearCookie("session-token")
	s.writeSessionToCookie(ctx, session)
}
