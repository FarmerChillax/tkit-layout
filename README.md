# tkit-layout

[tkit](https://github.com/FarmerChillax/tkit) 的示例项目

在线文档：https://farmerchillax.github.io/tkit/


## 各层级的作用
- server层: 用于参数校验、参数准备（从 metadata、request 或 context 等地方获取）
- service层: 处理主要流程代码、处理各类错误与记录相关日志
```shell
.
├── api // 下面维护了接口 schema（对应微服务使用的proto文件以及根据它们所生成的go文件）
├── internal
│   ├── model // 内部模型存储
│   │   ├── dto // 
│   │   └── mysql // MySQL 数据模型
│   ├── repository // 仓储层
│   ├── router // 用于路由注册
│   └── server // 用于参数校验、参数准备（从 metadata、request 或 context 等地方获取）
├── pkg
│   ├── errcode // 错误码存放
│   └── utils // 工具包
├── service // 处理主要流程代码。入参默认是正确的，无需处理的（也许应该放到 internal 下？）
└── startup // 程序启动依赖, 用于各种初始化（如 全局变量、配置、路由注册等）
```