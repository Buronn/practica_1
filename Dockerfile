FROM registry.gitlab.com/eit-udp/atraccion-talentos/devops/deployment:gorm-chromedp as builder
ADD . /build
RUN go build -o app main.go

FROM frolvlad/alpine-glibc
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates chromium\
        && update-ca-certificates 2>/dev/null || true
COPY --from=builder /build/ /app/
EXPOSE 3000
ENTRYPOINT ["/app/app"]
