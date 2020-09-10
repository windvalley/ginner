# Use Gin [![rcard](https://goreportcard.com/badge/github.com/windvalley/use-gin)](https://goreportcard.com/report/github.com/windvalley/use-gin) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/38d3eff769c14fecb01e91160e143727)](https://www.codacy.com/manual/windvalley/use-gin?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=windvalley/use-gin&amp;utm_campaign=Badge_Grade) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/windvalley/use-gin) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](LICENSE)

Using Go `Gin` to develop high quality applications(Web API) efficiently.

[简体中文](README_ZH.md)

## Features

- [x] Go Modules
- [x] Configuration
  - [x] toml
  - [x] Load config from command line parameters
  - [x] Load config from system environment variable
  - [x] If load config from CLI params failed, then load config from system environment variable
- [x] Logger
  - [x] Separation of error log and access log
  - [x] Customize log dirname, log save days, log rotate interval, log format(json/txt)
  - [x] When runmode is debug, log can also be output to screen
- [x] Error code system: for central managing the error response messages
- [x] Swagger
  - [x] Auto enable swagger when runmode is debug
- [x] Middlewares
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
  - [x] Basic Auth
  - [x] CORS
  - [x] RequestID(TraceID)
  - [x] Access log
  - [x] Global catch panic
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
  - [x] processlock: avoiding errors caused by repeated execution
  - [x] gomail
  - [x] paginate
  - [x] http request
  - [x] live reloading the server in development stage
- [x] Graceful restart or stop
- [x] Server health check when server started
- [x] SSL Support
- [x] Deployment
  - [x] Dockerfile
  - [x] Simple manage script: service.sh
  - [x] Systemd
  - [x] Supervisord

## Deployment

Change project name first:

```bash
git clone git@github.com:windvalley/use-gin.git
cd use-gin

# This will be change use-gin(current package name) to your own name.
./change_project_name.sh your-project-name
```

### Normal

```bash
./build.sh

# dev
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
