FROM golang:1.18.4-alpine3.15 as builder

COPY ./ /tmp/svapi

RUN apk add --no-cache git

RUN cd /tmp/svapi && \
    CGO_ENABLED=0 go build -o ./build/server -tags gingonic,release ./cmd/server

FROM alpine:3.15.4

COPY --from=builder /tmp/svapi/build/server /opt/svapi/server

EXPOSE 4200

CMD [ "/opt/svapi/server" ]
