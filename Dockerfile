FROM golang:1.24.1-alpine3.21 AS builder

COPY ./ /build

RUN apk add --no-cache git

RUN cd /build && \
    CGO_ENABLED=0 go build -o ./out/server -tags gingonic,release ./cmd/server

FROM alpine:3.21.3

COPY --from=builder /build/out/server /opt/svapi/server

EXPOSE 4200

CMD [ "/opt/svapi/server" ]
