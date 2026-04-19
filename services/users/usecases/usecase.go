package usecase

import "context"

type UseCase interface {
	ReadUsers(ctx context.Context)
	ReadUserById(ctx context.Context, id uint64)
	CreateUser(ctx context.Context, payload any)
	UpdateUser(ctx context.Context, id uint64, payload any)
	DeleteUser(ctx context.Context, id uint64)
}

type clientUseCase struct {
	statement any
}

func NewClientUseCase() *clientUseCase {
	return &clientUseCase{
		statement: nil,
	}
}
