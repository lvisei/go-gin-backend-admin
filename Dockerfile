FROM golang:alpine AS builder
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go-gin-backend-admin
COPY . $GOPATH/src/go-gin-backend-admin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-backend-admin .

FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /go-gin-backend-admin
COPY --from=builder go/src/go-gin-backend-admin/config ./config
COPY --from=builder go/src/go-gin-backend-admin/go-gin-backend-admin .
EXPOSE 8000
CMD ["./go-gin-backend-admin"]