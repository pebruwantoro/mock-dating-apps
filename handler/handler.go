package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pebruwantoro/technical_test_dealls/dto"
	"github.com/pebruwantoro/technical_test_dealls/helper"
)

func (s *Server) SignUp(c echo.Context) error {
	ctx := context.Background()
	req := dto.SignUpRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	result, err := s.Usecase.SignUp(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusCreated, dto.BaseResponse{
		Success: true,
		Message: "success sign up",
		Data:    result,
	})
}

func (s *Server) Login(c echo.Context) error {
	ctx := context.Background()
	req := dto.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	result, err := s.Usecase.Login(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: "success login",
		Data:    result,
	})
}

func (s *Server) Swipe(c echo.Context) error {
	ctx := context.Background()
	req := dto.SwipeRequest{}

	claims, ok := c.Get("claims").(*helper.Claims)
	if !ok || claims == nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized: Invalid or missing claims")
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: "error binding the request body",
		})
	}

	req.IsPremium = claims.IsPremium
	req.UserID = claims.UUID

	result, err := s.Usecase.Swipe(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: "success swipe",
		Data:    result,
	})
}

func (s *Server) PurchasePremiumPackage(c echo.Context) error {
	ctx := context.Background()
	req := dto.PurchasePremiumPackageRequest{}

	claims, ok := c.Get("claims").(*helper.Claims)
	if !ok || claims == nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized: Invalid or missing claims")
	}

	req.UserID = claims.UUID

	result, err := s.Usecase.PurchasePremiumPackage(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success: false,
			Message: fmt.Sprintf("error: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: "success purchase premium package",
		Data:    result,
	})
}
