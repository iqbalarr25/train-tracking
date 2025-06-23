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
		LoginAdmin(request *model.LoginAdminRequest) (*model.LoginResponse, int, error)
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

func (s *AuthService) LoginAdmin(request *model.LoginAdminRequest) (*model.LoginResponse, int, error) {
	user, err := s.Repository.GetUserByField("email", request.Email)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New(helper.ErrEmailNotFound)
	}

	if err := user.ComparePassword(request.Password); err != nil {
		return nil, http.StatusBadRequest, errors.New(helper.ErrEmailOrPasswordIsWrong)
	}

	if user.Role != helper.UserRoleAdmin {
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
