FROM golang:alpine
LABEL maintainer="d94.zaragoza@gmail.com"
WORKDIR /go/src/app
COPY . .
RUN go install && go clean
ENTRYPOINT ["app"]