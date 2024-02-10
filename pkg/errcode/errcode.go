package errcode

import "github.com/FarmerChillax/tkit"

var (
	// 客户端错误（4xx）
	InvalidParams             = tkit.NewError(400, 10000001, "入参错误")
	NotFound                  = tkit.NewError(404, 10000002, "找不到")
	UnauthorizedAuthNotExist  = tkit.NewError(401, 10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = tkit.NewError(401, 10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = tkit.NewError(401, 10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = tkit.NewError(401, 10000006, "鉴权失败，Token 生成失败")

	// 服务端错误（5xx）
	ServerError     = tkit.NewError(500, 20000001, "服务内部错误")
	TooManyRequests = tkit.NewError(500, 20000002, "请求过多")

	// 自定义错误
	DatabaseError = tkit.NewError(500, 30000001, "数据库错误")
)
