FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/Login-Backend
COPY . $GOPATH/Login-Backend
RUN go mod tidy && go build .

EXPOSE 8080
ENTRYPOINT ["./Login-Backend"]