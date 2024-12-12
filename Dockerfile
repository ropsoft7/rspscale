# Copyright (c) Rspscale Inc & AUTHORS
# SPDX-License-Identifier: BSD-3-Clause

# Note that this Dockerfile is currently NOT used to build any of the published
# Rspscale container images and may have drifted from the image build mechanism
# we use.
# Rspscale images are currently built using https://github.com/tailscale/mkctr,
# and the build script can be found in ./build_docker.sh.
#
#
# This Dockerfile includes all the rspscale binaries.
#
# To build the Dockerfile:
#
#     $ docker build -t ropsoft7/rspscale .
#
# To run the rspscaled agent:
#
#     $ docker run -d --name=rspscaled -v /var/lib:/var/lib -v /dev/net/tun:/dev/net/tun --network=host --privileged ropsoft7/rspscale rspscaled
#
# To then log in:
#
#     $ docker exec rspscaled rspscale up
#
# To see status:
#
#     $ docker exec rspscaled rspscale status


FROM golang:1.23-alpine AS build-env

WORKDIR /go/src/rspscale

COPY go.mod go.sum ./
RUN go mod download

# Pre-build some stuff before the following COPY line invalidates the Docker cache.
RUN go install \
    github.com/aws/aws-sdk-go-v2/aws \
    github.com/aws/aws-sdk-go-v2/config \
    gvisor.dev/gvisor/pkg/tcpip/adapters/gonet \
    gvisor.dev/gvisor/pkg/tcpip/stack \
    golang.org/x/crypto/ssh \
    golang.org/x/crypto/acme \
    github.com/coder/websocket \
    github.com/mdlayher/netlink

COPY . .

# see build_docker.sh
ARG VERSION_LONG=""
ENV VERSION_LONG=$VERSION_LONG
ARG VERSION_SHORT=""
ENV VERSION_SHORT=$VERSION_SHORT
ARG VERSION_GIT_HASH=""
ENV VERSION_GIT_HASH=$VERSION_GIT_HASH
ARG TARGETARCH

RUN GOARCH=$TARGETARCH go install -ldflags="\
      -X scale.ropsoft.cloud/version.longStamp=$VERSION_LONG \
      -X scale.ropsoft.cloud/version.shortStamp=$VERSION_SHORT \
      -X scale.ropsoft.cloud/version.gitCommitStamp=$VERSION_GIT_HASH" \
      -v ./cmd/rspscale ./cmd/rspscaled ./cmd/containerboot

FROM alpine:3.18
RUN apk add --no-cache ca-certificates iptables iproute2 ip6tables

COPY --from=build-env /go/bin/* /usr/local/bin/
# For compat with the previous run.sh, although ideally you should be
# using build_docker.sh which sets an entrypoint for the image.
RUN mkdir /rspscale && ln -s /usr/local/bin/containerboot /rspscale/run.sh
