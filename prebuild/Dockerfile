FROM golang:alpine
RUN mkdir /build
WORKDIR /build
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres
RUN go get github.com/chromedp/chromedp
RUN go get github.com/chromedp/cdproto/runtime
RUN go get github.com/minio/minio-go/v7