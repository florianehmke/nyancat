FROM alpine:3.10

RUN apk -U add ca-certificates

EXPOSE 8080

ADD controller /bin/controller

CMD ["controller"]
