FROM golang:alpine
LABEL maintainer="d94.zaragoza@gmail.com"
WORKDIR /go/src/app
COPY . .
ENTRYPOINT ["go", "run", "main.go"]