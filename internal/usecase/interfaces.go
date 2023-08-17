package usecase

import (
	"github.com/gofiber/fiber/v2"
	"wallet-test/internal/entity"
)

type (
	Repo interface {
		RepoInvoice(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error)
		RepoWithdraw(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error)
		CheckWallet(ctx *fiber.Ctx, request *entity.CheckWalletRequest, response *entity.CheckWalletResponse) (*entity.CheckWalletResponse, error)
	}
)
