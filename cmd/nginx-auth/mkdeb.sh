#!/usr/bin/env bash

set -e

VERSION=0.1.3
for ARCH in amd64 arm64; do
    CGO_ENABLED=0 GOARCH=${ARCH} GOOS=linux go build -o rspscale.nginx-auth .

    mkpkg \
        --out=rspscale-nginx-auth-${VERSION}-${ARCH}.deb \
        --name=rspscale-nginx-auth \
        --version=${VERSION} \
        --type=deb \
        --arch=${ARCH} \
        --postinst=deb/postinst.sh \
        --postrm=deb/postrm.sh \
        --prerm=deb/prerm.sh \
        --description="Rspscale NGINX authentication protocol handler" \
        --files=./rspscale.nginx-auth:/usr/sbin/rspscale.nginx-auth,./rspscale.nginx-auth.socket:/lib/systemd/system/rspscale.nginx-auth.socket,./rspscale.nginx-auth.service:/lib/systemd/system/rspscale.nginx-auth.service,./README.md:/usr/share/rspscale/nginx-auth/README.md

    mkpkg \
        --out=rspscale-nginx-auth-${VERSION}-${ARCH}.rpm \
        --name=rspscale-nginx-auth \
        --version=${VERSION} \
        --type=rpm \
        --arch=${ARCH} \
        --postinst=rpm/postinst.sh \
        --postrm=rpm/postrm.sh \
        --prerm=rpm/prerm.sh \
        --description="Rspscale NGINX authentication protocol handler" \
        --files=./rspscale.nginx-auth:/usr/sbin/rspscale.nginx-auth,./rspscale.nginx-auth.socket:/lib/systemd/system/rspscale.nginx-auth.socket,./rspscale.nginx-auth.service:/lib/systemd/system/rspscale.nginx-auth.service,./README.md:/usr/share/rspscale/nginx-auth/README.md
done
