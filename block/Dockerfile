FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY block /block
ENTRYPOINT ["/block"]
