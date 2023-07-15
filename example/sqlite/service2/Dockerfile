FROM golang:1.19 as builder

ENV CGO_ENABLED=1 GOOS=linux GOPROXY=${GOPROXY:-https://goproxy.cn,direct}

WORKDIR /build

COPY . .

RUN go mod tidy

RUN  go build -o main  -ldflags="-X main.BuildStamp=`date +%Y-%m-%d.%H:%M:%S`" main.go


CMD ["./main"]

EXPOSE 8000



