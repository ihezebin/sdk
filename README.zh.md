# sdk-go
对Web开发通用库做统一封装和管理。
集成了统一日志库、数据库工具、配置解析和HTTP中间件等。
具体使用示例见各个包README、example目录以及对应的test测试用例文件。

# 目录结构
```
./
├─code
├─configure
├─email
├─example
│  ├─configure
│  └─httpserver
│      ├─config
│      ├─handlers
│      ├─middleware
│      ├─proto
│      ├─routes
│      └─server
├─httpserver
│  ├─middleware
│  └─result
├─jwt
├─logger
│  └─hooks
├─model
│  ├─mongo
│  │  └─test
│  ├─mysql
│  └─redis
│      └─test
├─oss
│  ├─qiniu
│  └─ucloud
└─utils
├─mapper
├─slice
└─timer
```


