#!/usr/bin/env bash
# generate_cert_key.sh
#
# For generating a development cert and key


DOMAIN="localhost"


GOROOT="$(go env|grep GOROOT|awk -F\" '{print $2}')"

go run "$GOROOT"/src/crypto/tls/generate_cert.go --host="$DOMAIN"

