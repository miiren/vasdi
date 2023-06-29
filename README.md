# 项目目录结构说明

```
┌── view              视图层
├── adapter           适配层
│   ├── controller        http请求控制层
│   │   ├── xxx               xxx业务控制器层
│   │   └── xxx               xxx业务控制器层
│   ├── router            web框架路由管理目录
│   ├── middleware        web框架中间件
│   ├── job               定时任务、脚本启动控制层
│   └── vo     
│       ├── req           入参结构体层
│       └── res           出参结构体层
├── aservice          业务服务层
│   ├── bo                bo结构体层
│   ├── xxx               xxx业务服务层
│   └── xxx               xxx业务服务层
├── dservice          数据服务层
│   ├── po                po结构体层
│   ├── YYY               YYY数据服务层
│   └── YYY               YYY数据服务层
├── inservice         基础设施服务层
│   ├── cnf               
│   │   └── config.go     配置相关
│   ├── code              
│   │   └── errorCode.go  接口code码处理 
│   ├── consts            
│   │   └── redisKey.go   统一定义redisKey静态变量 
│   ├── log               
│   │   └── logger.go     日志组件
│   └── source            资源链接服务
│   │   ├── cache.go          内存缓存
│   │   ├── db.go             MySQL
│   │   └── redis.go          redis
│   └── util              项目util      
├── resource          资源文件目录
├── build.sh          项目构建脚本
└── main.go           项目启动入口

```

# 分层调用说明
