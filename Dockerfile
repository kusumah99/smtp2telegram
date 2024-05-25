FROM golang:1.22-alpine3.20 AS builder

RUN apk add --no-cache git ca-certificates mailcap

WORKDIR /app

COPY . .

# The image should be built with
# --build-arg ST_VERSION=`git describe --tags --always`
ARG ST_VERSION
ARG GOPROXY=direct
RUN CGO_ENABLED=0 GOOS=linux go build \
        -ldflags "-s -w \
            -X main.Version=${ST_VERSION:-UNKNOWN_RELEASE}" \
			-a -o ksa-smtp-to-telegram





FROM alpine:3.18

RUN apk add --no-cache ca-certificates mailcap

COPY --from=builder /app/ksa-smtp-to-telegram /ksa-smtp-to-telegram

USER daemon

ENV ST_SMTP_LISTEN="0.0.0.0:2525"
EXPOSE 2525

ENTRYPOINT ["/ksa-smtp-to-telegram"]
