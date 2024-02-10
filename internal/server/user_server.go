package server

import (
	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit-layout/api"
	"github.com/FarmerChillax/tkit-layout/internal/model/mysql"
	"github.com/FarmerChillax/tkit-layout/internal/repository"
	"github.com/FarmerChillax/tkit-layout/pkg/errcode"
	"github.com/FarmerChillax/tkit-layout/service"
	"github.com/gin-gonic/gin"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) Login(ctx *gin.Context, req *api.UserLoginRequest) (*api.UserLoginResponse, error) {
	// 参数校验
	// 这里进行业务参数校验
	// ...

	// 业务逻辑
	svc := service.NewUserService()
	return svc.Login(ctx.Request.Context(), req)
}

func (s *UserServer) Register(ctx *gin.Context, req *api.UserRegisterRequest) (*api.EmptyResponse, error) {
	// 业务校验
	db := tkit.Database.Get(ctx)
	userRepo := repository.NewUserRepo(db)
	exist, err := userRepo.Exist(&mysql.User{
		Username: req.Username,
	})
	if err != nil {
		tkit.Logger.Errorf(ctx.Request.Context(), "check user exist failed: %v", err)
		return nil, errcode.DatabaseError
	}

	if exist {
		return nil, errcode.InvalidParams.WithMsg("用户已存在")
	}

	// 业务逻辑
	svc := service.NewUserService()
	return svc.Register(ctx.Request.Context(), req)
}
