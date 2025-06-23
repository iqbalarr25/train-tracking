package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	UserServiceInterface interface {
		CreateUser(req model.AddUserRequest) (res model.User, err error)
		UpdateUser(req model.UpdateUserRequest, id int) (err error)
		DeleteUser(id int) (err error)
		GetPagination(req model.RequestPagination) (res []model.UserListResponse, count int64, err error)
	}

	UserService struct {
		Repo repository.UserRepositoryInterface
	}
)

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetPagination(req model.RequestPagination) (res []model.UserListResponse, count int64, err error) {
	res, count, err = s.Repo.GetUsersPagination(req)
	return
}

func (s *UserService) CreateUser(req model.AddUserRequest) (res model.User, err error) {
	res = model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err = s.Repo.CreateUser(res)
	if err != nil {
		return model.User{}, err
	}
	return
}

func (s *UserService) UpdateUser(req model.UpdateUserRequest, id int) (err error) {
	userReq := model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	err = s.Repo.UpdateColumnsByField(userReq, "id", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) (err error) {
	err = s.Repo.DeleteUserByField("id", id)
	if err != nil {
		return err
	}
	return nil
}
