FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go-gin-backend-admin
COPY . $GOPATH/src/go-gin-backend-admin
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./go-gin-backend-admin"]