# zscnetworksupport-api
## 介绍
这个仓库是用作网维系统的统一api服务后端，连接网维的主数据库，使用go编写，采用echo这个web框架

目前项目还非常不完善，不要应用到生产环境
## 配置
在主目录下配置`config.json`里面的内容，仓库里的内容只是演示，要获取生产配置，转到private仓库，nginx配置同理

## 部署
计划将本仓库部署到https://api.zsxyww.com

构建需要`go`，`make`程序，运行：

```sh

make build

```

