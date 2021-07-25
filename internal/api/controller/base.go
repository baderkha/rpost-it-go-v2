package controller

import (
	"fmt"
	"rpost-it-go/internal/api/service"
	"rpost-it-go/pkg/util/http"

	"github.com/gofiber/fiber"
)

// Basecontroller : Extend this controller to be able to get abstractions from fiber context
type Basecontroller struct {
	service service.IService
}

// Service : get back the service facade
func (b *Basecontroller) Service() service.IService {
	return b.service
}

// OK : generic ok status
func (b *Basecontroller) OK(ctx *fiber.Ctx, record interface{}) {
	ctx.Status(200).JSON(record)
}

// Created : send back created status
func (b *Basecontroller) Created(ctx *fiber.Ctx, record interface{}) {
	ctx.Status(201).JSON(record)
}

// Updated : sends back updated status
func (b *Basecontroller) Updated(ctx *fiber.Ctx, record interface{}) {
	ctx.Status(200).JSON(record)
}

// Deleted : sends back deleted status
func (b *Basecontroller) Deleted(ctx *fiber.Ctx, model string) {
	ctx.Status(200).SendString(fmt.Sprintf("Deleted %s successfully", model))
}

// BodyCouldNotParse : return this if you're trying to parse json
func (b *Basecontroller) BodyCouldNotParse(ctx *fiber.Ctx, err error) {
	ctx.Status(400).SendString(err.Error())
}

// BadUserInput : sends back a 400
func (b *Basecontroller) BadUserInput(ctx *fiber.Ctx, inputParamKey string) {
	ctx.Status(400).SendString(fmt.Sprintf("Expecting %s to not be empty", inputParamKey))
}

// StatusFromError : send back status encoded in error message
func (b *Basecontroller) StatusFromError(ctx *fiber.Ctx, err error) {
	ctx.Status(http.StatusFromError(err)).SendString(err.Error())
}
