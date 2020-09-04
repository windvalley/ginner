# Use Gin [![rcard](https://goreportcard.com/badge/github.com/windvalley/use-gin)](https://goreportcard.com/report/github.com/windvalley/use-gin) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/windvalley/use-gin) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](LICENSE)

`Go Gin`脚手架, 帮助用户高效地编写高质量的`Web API`.

## 主要特性

- [x] 项目依赖管理`Go Modules`
- [x] 配置文件组件
  - [x] 采用面向对象配置文件包`toml`
  - [x] 可从命令行加载配置文件
  - [x] 可从系统环境变量加载配置文件
  - [x] 以上两种加载配置文件方式共存时, 命令行方式优先
- [x] 日志记录组件
  - [x] 访问日志文件与错误日志文件分离
  - [x] 自定义日志目录名称、日志保留天数、日志轮转间隔时间、日志格式(`json/txt`)
  - [x] `runmode`为`debug`时日志写入到文件的同时也输出到屏幕上
- [x] 错误码组件, 将返回用户的错误信息进行统一管理
- [x] `Swagger`文档
  - [x] `runmode`为`debug`模式时启用`Swagger`文档
- [x] 路由中间件
  - [x] `JWT`鉴权
    - [x] 支持通过用户名和密码获取`JWT`的场景
    - [x] 支持`API`签名验证的场景, 用户需要提前向服务方申请`appKey`和`appSecret`
  - [x] `API`签名验证
    - [x] `HmacMd5`加密
    - [x] `HmacSha1`加密
    - [x] `HmacSha256`加密
    - [x] `Md5`组合加密
    - [x] `AES`对称加密
    - [x] `RSA`非对称加密
  - [x] `Basic Auth`
  - [x] 跨域资源共享(`CORS`)
  - [x] 请求唯一`ID`记录
  - [x] 访问日志
  - [x] 全局异常(`panic`)捕获
  - [x] `API`性能分析
  - [x] 限流
    - [x] 根据用户`IP`访问频率进行限制
    - [x] 服务端全局访问频率限制
  - [x] 访问限制
    - [x] `IP`白名单
    - [x] 服务端`API`白名单
- [x] 数据库组件
  - [x] 关系数据库
    - [x] `MySQL`
    - [x] `MssSQL`
    - [x] `PostgreSQL`
    - [x] `Sqlite`
  - [x] 缓存
    - [x] `Redis`
    - [x] `RedisCluster`
  - [x] 时序数据库
    - [x] `InfluxDB`
  - [x] 消息队列
    - [x] `Kafka`
- [x] 子项目`demo`
  - [x] 后台常驻进程示例: `cmd/daemonprocess`
- [x] 计划任务`cron`
- [x] 进程内缓存`cache2go`
- [x] 其他小工具(`util`)
  - [x] 进程锁`processlock`: 防止程序被重复执行导致未知错误
  - [x] 发送邮件`gomail`
  - [x] 分页`paginate`
  - [x] 请求客户端`httprequest`
- [x] 支持优雅地重启和停止
- [x] 项目部署
  - [x] 云原生支持, 提供`Dockerfile`
  - [x] 简单的服务管理脚本: `service.sh`
  - [x] `Systemd`
  - [x] `Supervisord`

## 项目部署

首先将当前项目名称`use-gin`, 改成你自己的项目名称:

```bash
# 执行完该脚本后, 当前Go项目的package名称将变更为your-project-name
./change_project_name.sh your-project-name
```

### 常规方式

```bash
./build.sh

# dev
export RUNENV=dev
# production
export RUNENV=prod

./service.sh start
```

### 容器方式

```bash
docker build -t your-project-name .

# dev
docker run --name your-project-name -p80:8000 -d your-project-name

# production
docker run --name your-project-name -p80:8000 -d -e RUNENV=prod your-project-name
```

### Systemd

`CentOS7+`系统下建议使用`Systemd`管理服务.

```bash
sudo cp your-project-name.service /usr/lib/systemd/system/

# 加载新的服务配置文件
sudo systemctl daemon-reload

# 开机自动启动服务
sudo systemctl enable your-project-name

# 生产环境
echo "RUNENV=prod" > conf/systemd.ENV

# 开发环境
echo "RUNENV=dev" > conf/systemd.ENV

# 启动服务
sudo systemctl start your-project-name

# 重启服务
sudo systemctl restart your-project-name

# 关闭服务
sudo systemctl stop your-project-name

# 查看当前服务运行状态
sudo systemctl status your-project-name -l
```

### Supervisord

非`CentOS7+`系统下建议使用`Supervisord`管理服务.

```bash
sudo cp supervisord.conf /etc/

# 重启supervisord, 加载新配置
sudo supervisorctl reload

# 启动服务
sudo supervisorctl start your-project-name

# 停止服务
sudo supervisorctl stop your-project-name

# 重启服务
sudo supervisorctl restart your-project-name

# 查看当前服务运行状态
sudo supervisorctl status your-project-name
```

## 授权许可

本项目采用`MIT`开源授权许可证, 完整的授权说明已放置在[LICENSE](LICENSE)文件中.
