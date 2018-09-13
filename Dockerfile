FROM golang:alpine AS build
COPY . /src
RUN cd /src && go build -o gophersays

FROM alpine
LABEL maintainer="d94.zaragoza@gmail.com"
WORKDIR /app
COPY --from=build /src/gophersays /app/
ENTRYPOINT ["./gophersays"]