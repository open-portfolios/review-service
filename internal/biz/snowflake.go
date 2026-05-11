package biz

import "context"

type SnowflakeRepo interface {
	Generate(context.Context) (int64, error)
}

type SnowflakeUsecase struct {
	repo SnowflakeRepo
}

func NewSnowflakeUsecase(repo SnowflakeRepo) *SnowflakeUsecase {
	return &SnowflakeUsecase{
		repo: repo,
	}
}

func (uc *SnowflakeUsecase) Generate(ctx context.Context) (int64, error) {
	return uc.repo.Generate(ctx)
}
