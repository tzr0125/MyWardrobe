参考结构

```
├── cmd
│   ├── apps # 子项目目录
│   │   ├── config.example.yaml # 启动的配置文件样例
│   │   └── server.go # 某子项目启动
│   └── main.go # 项目启动的入口
├── docs # 各种文档
└── pkg  # 项目核心代码包
    ├── conf
    │   └── conf.go # 全局的配置设置
    ├── docs # 各种文档
    ├── logger # 日志的设置
    ├── server # 项目的核心逻辑块， 如果是多个子模块，可以建多个不同的目录，这里示例中建立一个server目录
    │   ├── auth # 认证模块
    │   │   └── acl # 默认的acl认证
    │   ├── controller # 控制器模块
    │   │   ├── controller.go # 一个全局controller的定义
    │   │   └── hello.go # 一个api样例，每种类型的[api接口](https://www.zhihu.com/search?q=api%E6%8E%A5%E5%8F%A3&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3280363968%7D)，建立一个单独的文件，这样条理比较清晰
    │   ├── model # 存储结构的配置，各种数据库存储的model结构配置，如数据库的各个字段的定义等
    │   │   └── hello.go
    │   ├── router # 路由的各种定义
    │   │   ├── [middleware](https://www.zhihu.com/search?q=middleware&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3280363968%7D) # 各种中间件的的定义
    │   │   │   ├── auth.go # 认证
    │   │   │   ├── cache.go # 缓存
    │   │   │   ├── log.go # 日志
    │   │   │   └── ...
    │   │   ├── router.go # 路由全局配置
    │   │   └── routers.go # 各种路由的配置,api代码逻辑主要在这里配置
    │   └── storage # [后端存储](https://www.zhihu.com/search?q=%E5%90%8E%E7%AB%AF%E5%AD%98%E5%82%A8&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A3280363968%7D)的实现
    │       ├── cache # 缓存实现
    │       │   ├── local # 内存
    │       │   └── redis # redis
    │       └── db # 各种数据库的存储
    │           └── mysql
    └── util # 放置各种公用的函数、工具等
        └── util.go
```
