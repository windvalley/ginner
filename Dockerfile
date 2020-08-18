FROM golang:1.15

ENV GOPROXY https://goproxy.cn,direct
ENV RUNENV dev

WORKDIR /opt/use-gin
COPY . /opt/use-gin
RUN go mod tidy
RUN go build

EXPOSE 8000

ENTRYPOINT ["./use-gin"]
