FROM golang:alpine AS build-env

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /src

COPY . /src
RUN go build -o use-gin


FROM alpine

ENV RUNENV dev
WORKDIR /app

COPY --from=build-env /src/use-gin /app/
COPY --from=build-env /src/conf/dev.config.toml /src/conf/config.toml /app/conf/

EXPOSE 8000

#CMD sh -c "while true; do sleep 1; done"
ENTRYPOINT ["./use-gin"]
