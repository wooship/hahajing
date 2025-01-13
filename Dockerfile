FROM golang:1.11.1-alpine3.8

ENV TZ=Asia/Shanghai

WORKDIR "/go/src/hahajing"
ADD hahajing .
ADD favicon.ico .
COPY config ./config
COPY static ./static
EXPOSE 80
CMD ["nohup", "./hahajing", "server"]
