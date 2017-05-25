FROM alpine:3.3

ADD rabbitmq-worker /usr/local/bin/rabbitmq-worker

CMD ["/usr/local/bin/rabbitmq-worker"]
