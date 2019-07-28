FROM alpine:3.10

RUN apk -U add ca-certificates

EXPOSE 8080

ADD miner /bin/miner

CMD ["controller"]
