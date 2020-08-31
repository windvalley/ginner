use-gin
===

`Go Gin`脚手架, 帮助用户高效地编写高质量的`Web API`.


Features
===

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
- [x] 支持优雅地重启和停止
    - [x] 提供服务管理脚本`service.sh`
- [x] 云原生支持, 提供`Dockerfile`
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
    - [x] 跨域资源共享(`CORS`)
    - [x] 请求唯一`ID`记录
    - [x] 访问日志
    - [x] 全局异常(`panic`)捕获
    - [x] `IP`限流
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
    - [x] 进程锁, 防止进程被重复执行导致未知错误
    - [x] 发送邮件`gomail`
    - [x] 分页
    - [x] 请求客户端`httprequest`

Deployment
===

## Normal
```bash
./build.sh

# dev
export RUNENV=dev
# production
export RUNENV=prod

./service.sh start
```

## Docker
```bash
docker build -t use-gin .

# dev
docker container run --name use-gin -p80:8000 -d use-gin

# production
docker container run --name use-gin -p80:8000 -d -e RUNENV=prod use-gin
```

