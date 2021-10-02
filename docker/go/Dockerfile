FROM golang:1.16.5-alpine3.13

RUN apk update && apk add git && apk add tzdata

RUN go get -u "github.com/go-sql-driver/mysql"
RUN go get -u "github.com/jinzhu/gorm"
RUN go get -u "github.com/gin-gonic/gin"
RUN go get -u "github.com/google/uuid"
RUN go get -u "gopkg.in/go-ini/ini.v1"

WORKDIR /go/src/tasks/app
COPY ../../ /go/src/tasks