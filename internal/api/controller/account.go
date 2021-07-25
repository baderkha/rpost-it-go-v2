package controller

import (
	"fmt"
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/internal/api/service"
	"rpost-it-go/pkg/util/http"

	"github.com/gofiber/fiber"
)

type Account struct {
	service service.IService
}

func (c *Account) GetById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	record, err := c.service.GetAccountById(id)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(200).JSON(record)
}

func (c *Account) Search(ctx *fiber.Ctx) {
	searchTerm := ctx.Query("search-term")
	if searchTerm == "" {
		ctx.Status(400).SendString("Expecting search-term to not be empty")
		return
	}
	records := c.service.GetAccountByApproximateId(searchTerm)
	ctx.Status(200).JSON(records)
}

func (c *Account) Create(ctx *fiber.Ctx) {
	var account repo.Account
	err := ctx.BodyParser(&account)
	if err != nil {
		ctx.Status(400).SendString(err.Error())
		return
	}
	record, err := c.service.CreateAccount(&account)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(201).JSON(record)
}

func (c *Account) Update(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	var updateAccount repo.Account
	err := ctx.BodyParser(&updateAccount)
	if err != nil {
		ctx.Status(400).SendString(err.Error())
		return
	}
	updateAccount.ID = id
	record, err := c.service.UpdateAccount(&updateAccount)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(200).JSON(record)
}

func (c *Account) Delete(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	err := c.service.DeleteAccount(id)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	deleted := fmt.Sprintf("Account %s Deleted Successfully", id)
	ctx.Status(200).SendString(deleted)
}
