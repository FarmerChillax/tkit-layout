package service

import (
	"context"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit-layout/api"
	"github.com/FarmerChillax/tkit-layout/internal/model/mysql"
	"github.com/FarmerChillax/tkit-layout/internal/repository"
	"github.com/FarmerChillax/tkit-layout/pkg/errcode"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Login(ctx context.Context, req *api.UserLoginRequest) (*api.UserLoginResponse, error) {
	return &api.UserLoginResponse{
		Token: "fake-token",
	}, nil
}

func (us *UserService) Register(ctx context.Context, req *api.UserRegisterRequest) (*api.EmptyResponse, error) {
	db := tkit.Database.Get(ctx)
	userRepo := repository.NewUserRepo(db)
	err := userRepo.Create(&mysql.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		tkit.Logger.Errorf(ctx, "create user failed: %v", err)
		return nil, errcode.DatabaseError
	}

	return &api.EmptyResponse{}, nil
}
