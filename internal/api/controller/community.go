package controller

import (
	"fmt"
	"rpost-it-go/internal/api/repo"
	"rpost-it-go/internal/api/service"
	"rpost-it-go/pkg/util/http"

	"github.com/gofiber/fiber/v2"
)

type Community struct {
	service service.IService
}

func (c *Community) GetById(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	record, err := c.service.GetCommunityById(id)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(200).JSON(record)
}

func (c *Community) SearchCommunities(ctx *fiber.Ctx) {
	searchTerm := ctx.Query("search-term")
	if searchTerm == "" {
		ctx.Status(400).SendString("Expected search-term in the parameter")
		return
	}
	records := c.service.GetCommunityByApproixmateId(searchTerm)
	ctx.Status(200).JSON(records)
}

func (c *Community) GetByAccountId(ctx *fiber.Ctx) {
	accountId := ctx.Params("id")
	if accountId == "" {
		ctx.Status(400).SendString("Expected params to have account id")
		return
	}
	records := c.service.GetCommunitiesByAccountOwnerId(accountId)
	ctx.Status(200).JSON(records)
}

func (c *Community) Create(ctx *fiber.Ctx) {
	var com repo.Community

	err := ctx.BodyParser(&com)
	if err != nil {
		ctx.Status(400).SendString(err.Error())
		return
	}
	com.AccountOwnerId = getAccountId(ctx)
	record, err := c.service.CreateCommunity(&com)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(201).JSON(record)
}

func (c *Community) Update(ctx *fiber.Ctx) {
	var com repo.Community
	id := ctx.Params("id")

	err := ctx.BodyParser(&com)
	if err != nil {
		ctx.Status(400).SendString(err.Error())
		return
	}
	com.AccountOwnerId = getAccountId(ctx)
	record, err := c.service.UpdateCommunity(id, &com)
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(200).JSON(record)
}

func (c *Community) Delete(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	err := c.service.DeleteCommunity(id, getAccountId(ctx))
	if err != nil {
		ctx.Status(http.StatusFromError(err)).SendString(err.Error())
		return
	}
	ctx.Status(200).SendString(fmt.Sprintf("Deleted Community records with id %s successfully !", id))
}
