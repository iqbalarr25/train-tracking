package repository

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/pkg/helper"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	UserRepositoryInterface interface {
		GetUsersPagination(req model.RequestPagination) (res []model.UserListResponse, count int64, err error)
		GetUserByField(field string, value interface{}) (res *model.User, err error)
		UpdateColumnsByField(req model.User, field string, value interface{}) (err error)
		DeleteUserByField(field string, value interface{}) (err error)
		CreateUser(req model.User) (err error)
	}

	UserRepository struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUsersPagination(req model.RequestPagination) (res []model.UserListResponse, count int64, err error) {
	query := r.DB.Table("users").Where("deleted_at is null")

	if req.Search != "" {
		query = query.Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Where("name ILIKE ? or phone_country_code ilike ?", "%"+req.Search+"%", "%"+req.Search+"%")
		})
	}

	query.Count(&count)

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)
	if req.Sort != "" {
		query = query.Order(req.Sort)
	}
	err = query.Find(&res).Error
	return
}

func (r *UserRepository) GetUserByField(field string, value interface{}) (res *model.User, err error) {
	query := fmt.Sprintf(helper.DynamicFieldConditionQuery, field)
	err = r.DB.Table("users").Where(query, value).Where("deleted_at is null").First(&res).Error
	return
}
func (r *UserRepository) UpdateColumnsByField(req model.User, field string, value interface{}) (err error) {
	query := fmt.Sprintf(helper.DynamicFieldConditionQuery, field)
	err = r.DB.Where(query, value).Updates(&req).Error
	return
}
func (r *UserRepository) DeleteUserByField(field string, value interface{}) (err error) {
	query := fmt.Sprintf(helper.DynamicFieldConditionQuery, field)
	err = r.DB.Table("users").Where(query, value).Update("deleted_at", time.Now()).Error
	return
}
func (r *UserRepository) CreateUser(req model.User) (err error) {
	err = r.DB.Create(&req).Error
	return
}
