package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pebruwantoro/technical_test_dealls/dto"
	"github.com/pebruwantoro/technical_test_dealls/helper"
	"github.com/pebruwantoro/technical_test_dealls/repository"
)

func (u *Usecase) SignUp(ctx context.Context, request dto.SignUpRequest) (result dto.SignUpResponse, err error) {
	hash, err := helper.HashPassword(request.Password)
	if err != nil {
		return
	}

	var user repository.User
	user, err = u.Repository.CreateUser(ctx, repository.User{
		UUID:      uuid.NewString(),
		Email:     request.Email,
		Username:  request.Username,
		Password:  hash,
		IsPremium: false,
		CreatedAt: time.Now(),
	})

	result.Id = user.UUID

	return
}

func (u *Usecase) Login(ctx context.Context, request dto.LoginRequest) (result dto.LoginResponse, err error) {
	var userModel repository.User
	var token string

	if strings.Contains(request.Identity, "@") {
		userModel, err = u.Repository.GetUserByEmail(ctx, request.Identity)
		if err != nil {
			return
		}
	} else {
		userModel, err = u.Repository.GetUserByUsername(ctx, request.Identity)
		if err != nil {
			return
		}
	}

	if !helper.ValidatePassword(userModel.Password, request.Password) {
		return
	}

	token, err = helper.GenerateJWT(userModel)
	if err != nil {
		return
	}

	result.Id = userModel.UUID
	result.Token = token

	return
}

func (u *Usecase) Swipe(ctx context.Context, request dto.SwipeRequest) (result dto.SwipeResponse, err error) {
	var swipeModel repository.Swipe
	isFirstSwipe := false

	swipeModel, err = u.Repository.GetSwipeByUserIdAndTargetId(ctx, request.UserID, request.TargetID)
	if errors.Is(err, sql.ErrNoRows) {
		isFirstSwipe = true
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	if isFirstSwipe {
		swipeModel, err = u.Repository.CreateSwipe(ctx, repository.Swipe{
			UUID:      uuid.NewString(),
			UserID:    request.UserID,
			TargetID:  request.TargetID,
			Direction: request.Direction,
		})
		if err != nil {
			return
		}

		result.Id = swipeModel.UUID

		return
	}

	swipeModel, err = u.Repository.UpdateSwipe(ctx, repository.Swipe{
		UUID:      swipeModel.UUID,
		UserID:    swipeModel.UserID,
		TargetID:  swipeModel.TargetID,
		Direction: request.Direction,
	})
	if err != nil {
		return
	}

	result.Id = swipeModel.UUID

	return
}
