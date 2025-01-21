package usecase

import "github.com/pebruwantoro/technical_test_dealls/repository"

type Usecase struct {
	Repository repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) *Usecase {
	return &Usecase{
		Repository: repo,
	}
}
