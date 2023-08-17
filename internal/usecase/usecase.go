package usecase

import (
	"github.com/gofiber/fiber/v2"
	"wallet-test/internal/databases/repos"
	"wallet-test/internal/entity"
)

type UseCase struct {
	repo Repo
}

func NewUseCase(r *repos.Repositories) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (use *UseCase) UseCaseInvoice(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
	response, err := use.repo.RepoInvoice(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, err
	}

	return response, nil
}

func (use *UseCase) UseCaseWithdraw(ctx *fiber.Ctx, request *entity.Request, response *entity.Response) (*entity.Response, error) {
	response, err := use.repo.RepoWithdraw(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}
	return response, nil
}

func (use *UseCase) UseCaseCheckStatusWallet(ctx *fiber.Ctx, request *entity.CheckWalletRequest, response *entity.CheckWalletResponse) (*entity.CheckWalletResponse, error) {
	response, err := use.repo.CheckWallet(ctx, request, response)
	if err != nil {
		response.Status = "Error"
		return response, nil
	}
	return response, nil
}
