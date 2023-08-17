package http

import (
	"github.com/gofiber/fiber/v2"
	"wallet-test/internal/entity"
)

func (routes *Routes) wallet(ctx *fiber.Ctx) error {
	request := new(entity.CheckWalletRequest)
	response := new(entity.CheckWalletResponse)

	err := ctx.BodyParser(request)
	if err != nil {
		return err
	}

	response, err = routes.Usecase.UseCaseCheckStatusWallet(ctx, request, response)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
