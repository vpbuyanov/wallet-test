package http

import (
	"github.com/gofiber/fiber/v2"
	"wallet-test/internal/entity"
)

func (routes *Routes) invoice(ctx *fiber.Ctx) error {
	request := new(entity.Request)
	response := new(entity.Response)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseInvoice(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
