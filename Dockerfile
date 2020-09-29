FROM golang:alpine AS build-env

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /src

COPY . /src
RUN go build -o ginner


FROM alpine:3

ENV RUNENV dev
WORKDIR /app

COPY --from=build-env /src/ginner /app/
COPY --from=build-env /src/conf/dev.config.toml /src/conf/config.toml /app/conf/

EXPOSE 8000

#CMD sh -c "while true; do sleep 1; done"
ENTRYPOINT ["./ginner"]
