FROM golang:1.23.2-alpine3.19 AS builder

COPY ./ /build

RUN apk add --no-cache git

RUN cd /build && \
    CGO_ENABLED=0 go build -o ./out/server -tags gingonic,release ./cmd/server

FROM alpine:3.19.4

COPY --from=builder /build/out/server /opt/svapi/server

EXPOSE 4200

CMD [ "/opt/svapi/server" ]
