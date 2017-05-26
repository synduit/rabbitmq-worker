FROM alpine:3.3

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ADD rabbitmq-worker /usr/local/bin/rabbitmq-worker

CMD ["/usr/local/bin/rabbitmq-worker"]
