package startup

import (
	"github.com/FarmerChillax/tkit-layout/internal/router"
	"github.com/gin-gonic/gin"
)

func Register(mux *gin.Engine) error {

	// 注册用户路由组
	if err := router.RegisterUserRouter(mux); err != nil {
		return err
	}

	return nil
}
