package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Abstraction for the auth controller , so we can make it portable if we want to switch to jwt
type IAuth interface {
	// Login : login to api
	Login(ctx *fiber.Ctx)
	// Logout of api
	Logout(ctx *fiber.Ctx)
	// Verify credentials
	// requested-account-id , which will bethe index of the account
	Verify(ctx *fiber.Ctx)
}
