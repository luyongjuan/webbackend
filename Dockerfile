FROM golang:latest

ENV TZ=Asia/Shanghai

COPY ./ /go/src/luyongjuan/webbackend
WORKDIR /go/src/luyongjuan/webbackend
RUN go install ./cmd/...

EXPOSE 8080
ENTRYPOINT  ["app"]
