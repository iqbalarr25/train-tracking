package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
	"TrainTracking/pkg/helper"
	"errors"
	"net/http"
)

type (
	AuthServiceInterface interface {
		Login(request *model.LoginRequest) (*model.LoginResponse, int, error)
		Register(request *model.RegisterRequest) (*model.RegisterResponse, int, error)
	}

	AuthService struct {
		Repository repository.UserRepositoryInterface
	}
)

func NewAuthService(repo repository.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		Repository: repo,
	}
}

func (s *AuthService) Login(request *model.LoginRequest) (*model.LoginResponse, int, error) {
	user, err := s.Repository.GetUserByField("email", request.Email)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New(helper.ErrEmailNotFound)
	}

	if err := user.ComparePassword(request.Password); err != nil {
		return nil, http.StatusBadRequest, errors.New(helper.ErrEmailOrPasswordIsWrong)
	}

	res := model.LoginResponse{
		Id:    user.ID,
		Email: user.Email,
	}

	token, err := helper.CreateToken(user.Email, user.Role)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	res.Token = token
	err = s.Repository.UpdateColumnsByField(
		model.User{
			Token: token,
		},
		"id", user.ID,
	)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &res, http.StatusOK, nil
}

func (s *AuthService) Register(request *model.RegisterRequest) (*model.RegisterResponse, int, error) {
	_, err := s.Repository.GetUserByField("email", request.Email)
	if err == nil {
		return nil, http.StatusBadRequest, errors.New(helper.ErrEmailExists)
	}

	if request.Password != request.ConfirmationPassword {
		return nil, http.StatusBadRequest, errors.New(helper.ErrPasswordConfirmationNotMatch)
	}

	user := &model.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Role:     helper.UserRoleMember,
	}

	err = s.Repository.CreateUser(*user)

	res := model.RegisterResponse{
		Email: user.Email,
		Name:  user.Name,
	}

	token, err := helper.CreateToken(user.Email, user.Role)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	res.Token = token
	err = s.Repository.UpdateColumnsByField(
		model.User{
			Token: token,
		},
		"id", user.ID,
	)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &res, http.StatusOK, nil
}
