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
			-a -o ksa-smtp2telegram





FROM alpine:3.20

RUN apk add --no-cache ca-certificates mailcap

COPY --from=builder /app/ksa-smtp2telegram /ksa-smtp2telegram
COPY --from=builder /app/.env /.env

USER daemon

#ENV SMTP_LISTEN_ADDRESS="0.0.0.0:1025"
#ENV EMAIL_DOMAIN_TELEGRAM="ksatele.gram"	# konfigurasi ini nimpa seting .env
EXPOSE 1025

ENTRYPOINT ["/ksa-smtp2telegram"]
