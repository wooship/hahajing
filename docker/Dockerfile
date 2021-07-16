FROM golang:1.11.1-alpine3.8

ENV TZ=Asia/Shanghai

RUN apk --update add git \
#  && mkdir /usr/local/go/src/vendor/github.com \
#  && cd /usr/local/go/src/vendor/github.com \
#  && mkdir PuerkitoBio \
#  && cd PuerkitoBio \
#  && git clone https://github.com/PuerkitoBio/goquery \
#  && cd /usr/local/go/src/vendor/github.com \
#  && mkdir andybalholm \
#  && cd andybalholm \
#  && git clone https://github.com/andybalholm/cascadia \
  && mkdir -p /go/src/github.com/op \
  && cd /go/src/github.com/op \
  && git clone https://github.com/op/go-logging \
  && mkdir -p /go/src/golang.org/x \
  && cd /go/src/golang.org/x \
  && git clone https://github.com/golang/crypto \
  && git clone https://github.com/golang/net \
  && cd /go/src \
  && git clone https://github.com/wooship/hahajing \
  && cd /go/src/hahajing \
  && set GOARCH=amd64 \
  && set GOOS=linux \
  && go build

WORKDIR "/go/src/hahajing"
EXPOSE 80
CMD ["nohup", "./hahajing", "server"]
