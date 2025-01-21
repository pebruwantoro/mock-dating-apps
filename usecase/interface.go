package usecase

import (
	"context"

	"github.com/pebruwantoro/technical_test_dealls/dto"
)

//go:generate mockgen -destination=mocks/usecase.go -package=mocks . UsecaseInterface
type UsecaseInterface interface {
	SignUp(ctx context.Context, request dto.SignUpRequest) (dto.SignUpResponse, error)
	Login(ctx context.Context, request dto.LoginRequest) (dto.LoginResponse, error)
	Swipe(ctx context.Context, request dto.SwipeRequest) (dto.SwipeResponse, error)
	PurchasePremiumPackage(ctx context.Context, request dto.PurchasePremiumPackageRequest) (dto.PurchasePremiumPackageResponse, error)
}
