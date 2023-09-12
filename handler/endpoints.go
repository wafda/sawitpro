package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetProfile(ctx echo.Context, params generated.GetProfileParams) error {

	resp, err := s.Repository.GetProfile(ctx.Request().Context(), repository.GetProfileParams{
		Token: params.Token,
	})
	if err != nil {
		return ctx.JSON(http.StatusForbidden, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PutProfile(ctx echo.Context, params generated.PutProfileParams) error {
	var req generated.User
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	resp, err := s.Repository.UpdateProfile(ctx.Request().Context(), repository.PutProfileParams{
		Token: params.Token,
	}, repository.User{
		FullName:     req.FullName,
		PhoneNumbers: req.PhoneNumbers,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) RegisterUser(ctx echo.Context) error {
	var req generated.RegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	resp, err := s.Repository.RegisterUser(ctx.Request().Context(), repository.RegisterRequest{
		FullName:             &req.FullName,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		PhoneNumbers:         &req.PhoneNumbers,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) LoginUser(ctx echo.Context) error {
	var req generated.LoginRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := s.Repository.Login(ctx.Request().Context(), repository.LoginRequest{
		Password:     req.Password,
		PhoneNumbers: req.PhoneNumbers,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}
