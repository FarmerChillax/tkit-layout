package router

import (
	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit-layout/internal/server"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(mux *gin.Engine) error {
	userGroup := mux.Group("/user")
	{
		userServer := server.NewUserServer()
		userGroup.GET("/login", tkit.Wrap(userServer.Login))
		userGroup.POST("/register", tkit.Wrap(userServer.Register))
	}

	return nil
}
