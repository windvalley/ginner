# Ginner [![Go Report Card](https://goreportcard.com/badge/github.com/windvalley/ginner)](https://goreportcard.com/report/github.com/windvalley/ginner) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/7b817bd0effe4d02bbea489ba0541edb)](https://www.codacy.com/gh/windvalley/ginner/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=windvalley/ginner&amp;utm_campaign=Badge_Grade) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=windvalley_ginner&metric=alert_status)](https://sonarcloud.io/dashboard?id=windvalley_ginner) ![Go](https://github.com/windvalley/ginner/workflows/Go/badge.svg) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/windvalley/ginner) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](LICENSE)

Based on the Gin framework, this project integrates commonly used web components and tools, and demonstrates a reasonable code organization structure to help users develop web back-end applications efficiently.

[简体中文](README_ZH.md)

## Features

- [x] Go Modules
- [x] Configuration
  - [x] Object-oriented configuration
  - [x] Load configuration files from command line parameters
  - [x] Load configuration files from system environment variables
  - [x] Load configuration files from command line parameters takes precedence
- [x] Logger
  - [x] Separation of error log and access log
  - [x] The log format, log retention time, log rotation interval, log directory name can be customized
  - [x] If the runmode is debug, the log will be output to the screen at the same time
- [x] Error code system
  - [x] For centralized managing error response information
- [x] Swagger
  - [x] If the runmode is debug, swagger will be automatically enabled
- [x] Middlewares
  - [x] JWT Authentication
    - [x] Login: get JWT token by username and password
    - [x] API Signature Authentication: Users need to apply for `appKey` and `appSecret` in advance
  - [x] API Signature Authentication
    - [x] HmacMd5
    - [x] HmacSha1
    - [x] HmacSha256
    - [x] Md5
    - [x] AES
    - [x] RSA
  - [x] Basic Auth
  - [x] CORS
  - [x] RequestID(TraceID)
  - [x] Access log
  - [x] User operation audit
  - [x] Global panic catch and recover
  - [x] Pprof
  - [x] Limiter
    - [x] Request rate limiter based on client ip
    - [x] Global request rate limiter
  - [x] ACL
    - [x] IP allowlist
    - [x] Server API allowlist
- [x] Databases
  - [x] Relation Database(GORM)
    - [x] MySQL
    - [x] MssSQL
    - [x] PostgreSQL
    - [x] Sqlite
  - [x] NoSQL
    - [x] MongoDB
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
- [x] Go cache: cache2go
- [x] Utils
  - [x] Processlock: avoid errors caused by repeated execution of the program(process)
  - [x] Live reloading gin server in development phase
  - [x] Pagination
  - [x] Redis mutex
  - [x] HTTP request
  - [x] Gomail
- [x] Graceful restart or stop gin server
- [x] Health check when the server starts
- [x] SSL support
- [x] Deployment
  - [x] Dockerfile
  - [x] Systemd
  - [x] Supervisord
  - [x] Simple manage script: service.sh

## Deployment

First you need to change the current project name to your own:

```bash
git clone git@github.com:windvalley/ginner.git
cd ginner

# This will change the current project name `ginner` to your own project name.
./change_project_name.sh your-project-name
```

### Normal

```bash
./build.sh

# development
export RUNENV=dev
# production
export RUNENV=prod

# start
./service.sh start

# restart
./service.sh restart

# graceful reload
./service.sh reload

# graceful stop
./service.sh stop

# check status
./service.sh status
```

### Docker

```bash
# build image
docker build -t your-project-name .

# dev
docker run --name your-project-name -p80:8000 -d your-project-name

# production
docker run --name your-project-name -p80:8000 -d -e RUNENV=prod your-project-name
```

### Systemd

```bash
sudo cp your-project-name.service /usr/lib/systemd/system/

# reload the new service config file
sudo systemctl daemon-reload

# autostart after rebooting
sudo systemctl enable your-project-name

# production
echo "RUNENV=prod" > conf/systemd.ENV

# development
echo "RUNENV=dev" > conf/systemd.ENV

# start
sudo systemctl start your-project-name

# restart
sudo systemctl restart your-project-name

# graceful stop
sudo systemctl stop your-project-name

# check status
sudo systemctl status your-project-name -l
```

### Supervisord

```bash
sudo cp supervisord.conf /etc/

# supervisord restart
sudo supervisorctl reload

# start
sudo supervisorctl start your-project-name

# graceful stop
sudo supervisorctl stop your-project-name

# restart
sudo supervisorctl restart your-project-name

# check status
sudo supervisorctl status your-project-name
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
