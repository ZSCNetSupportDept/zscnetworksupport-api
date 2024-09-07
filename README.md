# zscnetworksupport-api
## 介绍
这个仓库是用作网维系统的统一api服务后端，连接网维的主数据库，使用go编写，采用echo这个web框架

目前项目还非常不完善，不要应用到生产环境
## 配置
配置`config.json`里面的内容，仓库里的内容只是演示，要获取生产配置，转到private仓库，nginx配置同理

## 构建
`make build`,然后查看`./make`文件夹
## 部署
计划将本仓库部署到https://api.zsxyww.com

构建需要`go`，`make`程序，运行：

```sh

make build

```

## 使用
把生成的二进制文件放在生产服务器下，指定`config.json`运行即可，你可以设置系统服务来为程序添加开机自启动功能

你还需要配置nginx或者其他反向代理工具，将流量转发到程序所监听的端口

