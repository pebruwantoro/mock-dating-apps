package repository

import "context"

//go:generate mockgen -destination=mocks/repository.go -package=mocks . RepositoryInterface
type RepositoryInterface interface {
	CreateUser(ctx context.Context, input User) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	CreateSwipe(ctx context.Context, input Swipe) (Swipe, error)
	UpdateSwipe(ctx context.Context, input Swipe) (Swipe, error)
	GetSwipeByUserIdAndTargetId(ctx context.Context, userId, targetId string) (Swipe, error)
}
