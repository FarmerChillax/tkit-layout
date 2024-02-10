package startup

import (
	"context"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit-layout/internal/model/mysql"
)

func RegisterCallback() map[tkit.CallbackPosition]tkit.CallbackFunc {
	return map[tkit.CallbackPosition]tkit.CallbackFunc{
		tkit.POSITION_NEW: func() error {
			// 初始化数据库
			tkit.Database.Get(context.Background()).AutoMigrate(&mysql.User{})
			tkit.Logger.Info(context.Background(), "init database success")
			return nil
		},
	}
}
