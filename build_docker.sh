#!/usr/bin/env sh
#
# This script builds Rspscale container images using
# github.com/tailscale/mkctr.
# By default the images will be tagged with the current version and git
# hash of this repository as produced by ./cmd/mkversion.
# This is the image build mechanim used to build the official Rspscale
# container images.

set -eu

# Use the "go" binary from the "tool" directory (which is github.com/tailscale/go)
export PATH="$PWD"/tool:"$PATH"

eval "$(./build_dist.sh shellvars)"

DEFAULT_TARGET="client"
DEFAULT_TAGS="v${VERSION_SHORT},v${VERSION_MINOR}"
DEFAULT_BASE="rspscale/alpine-base:3.18"
# Set a few pre-defined OCI annotations. The source annotation is used by tools such as Renovate that scan the linked
# Github repo to find release notes for any new image tags. Note that for official Rspscale images the default
# annotations defined here will be overriden by release scripts that call this script.
# https://github.com/opencontainers/image-spec/blob/main/annotations.md#pre-defined-annotation-keys
DEFAULT_ANNOTATIONS="org.opencontainers.image.source=https://github.com/ropsoft7/rspscale/blob/main/build_docker.sh,org.opencontainers.image.vendor=Rspscale"

PUSH="${PUSH:-false}"
TARGET="${TARGET:-${DEFAULT_TARGET}}"
TAGS="${TAGS:-${DEFAULT_TAGS}}"
BASE="${BASE:-${DEFAULT_BASE}}"
PLATFORM="${PLATFORM:-}" # default to all platforms
# OCI annotations that will be added to the image.
# https://github.com/opencontainers/image-spec/blob/main/annotations.md
ANNOTATIONS="${ANNOTATIONS:-${DEFAULT_ANNOTATIONS}}"

case "$TARGET" in
  client)
    DEFAULT_REPOS="ropsoft7/rspscale"
    REPOS="${REPOS:-${DEFAULT_REPOS}}"
    go run github.com/tailscale/mkctr \
      --gopaths="\
        scale.ropsoft.cloud/cmd/rspscale:/usr/local/bin/rspscale, \
        scale.ropsoft.cloud/cmd/rspscaled:/usr/local/bin/rspscaled, \
        scale.ropsoft.cloud/cmd/containerboot:/usr/local/bin/containerboot" \
      --ldflags="\
        -X scale.ropsoft.cloud/version.longStamp=${VERSION_LONG} \
        -X scale.ropsoft.cloud/version.shortStamp=${VERSION_SHORT} \
        -X scale.ropsoft.cloud/version.gitCommitStamp=${VERSION_GIT_HASH}" \
      --base="${BASE}" \
      --tags="${TAGS}" \
      --gotags="ts_kube,ts_package_container" \
      --repos="${REPOS}" \
      --push="${PUSH}" \
      --target="${PLATFORM}" \
      --annotations="${ANNOTATIONS}" \
      /usr/local/bin/containerboot
    ;;
  k8s-operator)
    DEFAULT_REPOS="rspscale/k8s-operator"
    REPOS="${REPOS:-${DEFAULT_REPOS}}"
    go run github.com/tailscale/mkctr \
      --gopaths="scale.ropsoft.cloud/cmd/k8s-operator:/usr/local/bin/operator" \
      --ldflags="\
        -X scale.ropsoft.cloud/version.longStamp=${VERSION_LONG} \
        -X scale.ropsoft.cloud/version.shortStamp=${VERSION_SHORT} \
        -X scale.ropsoft.cloud/version.gitCommitStamp=${VERSION_GIT_HASH}" \
      --base="${BASE}" \
      --tags="${TAGS}" \
      --gotags="ts_kube,ts_package_container" \
      --repos="${REPOS}" \
      --push="${PUSH}" \
      --target="${PLATFORM}" \
      --annotations="${ANNOTATIONS}" \
      /usr/local/bin/operator
    ;;
  k8s-nameserver)
    DEFAULT_REPOS="rspscale/k8s-nameserver"
    REPOS="${REPOS:-${DEFAULT_REPOS}}"
    go run github.com/tailscale/mkctr \
      --gopaths="scale.ropsoft.cloud/cmd/k8s-nameserver:/usr/local/bin/k8s-nameserver" \
      --ldflags=" \
        -X scale.ropsoft.cloud/version.longStamp=${VERSION_LONG} \
        -X scale.ropsoft.cloud/version.shortStamp=${VERSION_SHORT} \
        -X scale.ropsoft.cloud/version.gitCommitStamp=${VERSION_GIT_HASH}" \
      --base="${BASE}" \
      --tags="${TAGS}" \
      --gotags="ts_kube,ts_package_container" \
      --repos="${REPOS}" \
      --push="${PUSH}" \
      --target="${PLATFORM}" \
      --annotations="${ANNOTATIONS}" \
      /usr/local/bin/k8s-nameserver
    ;;
  *)
    echo "unknown target: $TARGET"
    exit 1
    ;;
esac
