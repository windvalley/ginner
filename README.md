# Use Gin [![rcard](https://goreportcard.com/badge/github.com/windvalley/use-gin)](https://goreportcard.com/report/github.com/windvalley/use-gin) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/windvalley/use-gin) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/windvalley/use-gin/master/LICENSE)

Using Go `Gin` to develop high quality applications(Web API) efficiently.

[简体中文](README_ZH.md)

## Features
- [x] Go Modules
- [x] Configuration
    - [x] BurntSushi/toml
    - [x] Load config from command line parameters
    - [x] Load from system environment variable RUNENV: `prod/dev`
    - [x] If load config from CLI params failed, then load config from system environment variable RUNENV
- [x] Logger
    - [x] Create `access.log` and `error.log`
    - [x] Define log dir name, log save days, log rotate interval, log format(json/txt)
    - [x] When runmode is debug, log can also be output to screen
- [x] Error code system, for central managing the error response messages
- [x] Graceful restart or stop
    - [x] Manage script: service.sh
- [x] Dockerfile
- [x] Swagger
    - [x] Auto enable swagger when runmode is debug
- [x] Middleware
    - [x] JWT Authentication
        - [x] Login: get JWT token by username and password
        - [x] API Signature Authentication: user need to apply `appKey` and `appSecret` in advance
    - [x] API Signature Authentication
        - [x] HmacMd5
        - [x] HmacSha1
        - [x] HmacSha256
        - [x] Md5
        - [x] AES
        - [x] RSA
    - [x] CORS
    - [x] X-Request-Id
    - [x] Accesslog
    - [x] Global catch panic
    - [x] IP Limiter
    - [x] ACL
        - [x] IP allow list
        - [x] Server API allow list
- [x] Databases
    - [x] Relation Database
        - [x] MySQL
        - [x] MssSQL
        - [x] PostgreSQL
        - [x] Sqlite
    - [x] Cache
        - [x] Redis
        - [x] RedisCluster
    - [x] Time Series Database
        - [x] InfluxDB
    - [x] MQ
        - [x] Kafka
- [x] Subproject demo
    - [x] cmd/daemonprocess
- [x] Crontab: cron
- [x] Go Cache: cache2go
- [x] Utils: util
    - [x] processlock, for avoiding errors caused by repeated execution
    - [x] gomail
    - [x] pagination
    - [x] httprequest

## Deployment

### Normal
```bash
./build.sh

# dev
export RUNENV=dev
# production
export RUNENV=prod

./service.sh start
```

### Docker
```bash
docker build -t use-gin .

# dev
docker container run --name use-gin -p80:8000 -d use-gin

# production
docker container run --name use-gin -p80:8000 -d -e RUNENV=prod use-gin
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
