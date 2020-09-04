# Use Gin [![rcard](https://goreportcard.com/badge/github.com/windvalley/use-gin)](https://goreportcard.com/report/github.com/windvalley/use-gin) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/windvalley/use-gin) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](LICENSE)

Using Go `Gin` to develop high quality applications(Web API) efficiently.

[简体中文](README_ZH.md)

## Features

- [x] Go Modules
- [x] Configuration
  - [x] toml
  - [x] Load config from command line parameters
  - [x] Load from system environment variable RUNENV: `prod/dev`
  - [x] If load config from CLI params failed, then load config from system environment variable RUNENV
- [x] Logger
  - [x] Create `access.log` and `error.log`
  - [x] Define log dir name, log save days, log rotate interval, log format(json/txt)
  - [x] When runmode is debug, log can also be output to screen
- [x] Error code system, for central managing the error response messages
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
  - [x] Basic Auth
  - [x] CORS
  - [x] X-Request-Id
  - [x] Accesslog
  - [x] Global catch panic
  - [x] Pprof
  - [x] Limiter
    - [x] Request rate limiter based on client ip
    - [x] Global request limiter
  - [x] ACL
    - [x] IP allowlist
    - [x] Server API allowlist
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
- [x] Util
  - [x] processlock: avoiding errors caused by repeated execution
  - [x] gomail
  - [x] paginate
  - [x] httprequest
- [x] Graceful restart or stop
- [x] Deployment
  - [x] Dockerfile
  - [x] Simple manage script: service.sh
  - [x] Systemd
  - [x] Supervisord

## Deployment

Change project name first:

```bash
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

./service.sh start
```

### Docker

```bash
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

# stop
sudo supervisorctl stop your-project-name

# restart
sudo supervisorctl restart your-project-name

# check status
sudo supervisorctl status your-project-name
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
