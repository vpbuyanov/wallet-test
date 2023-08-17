package http

import (
	"github.com/gofiber/fiber/v2"
	"wallet-test/internal/usecase"
)

type Routes struct {
	Usecase *usecase.UseCase
}

func (routes *Routes) RegisterRoutes(app fiber.Router) {
	app.Post("/invoice", routes.invoice)
	app.Post("/withdraw", routes.withdraw)
	app.Post("/wallet", routes.wallet)
}
