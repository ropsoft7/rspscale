IMAGE_REPO ?= ropsoft7/rspscale
SYNO_ARCH ?= "x86_64"
SYNO_DSM ?= "7"
TAGS ?= "latest"

PLATFORM ?= "flyio" ## flyio==linux/amd64. Set to "" to build all platforms.

vet: ## Run go vet
	./tool/go vet ./...

tidy: ## Run go mod tidy
	./tool/go mod tidy

lint: ## Run golangci-lint
	./tool/go run github.com/golangci/golangci-lint/cmd/golangci-lint run

updatedeps: ## Update depaware deps
	# depaware (via x/tools/go/packages) shells back to "go", so make sure the "go"
	# it finds in its $$PATH is the right one.
	PATH="$$(./tool/go env GOROOT)/bin:$$PATH" ./tool/go run github.com/tailscale/depaware --update \
		scale.ropsoft.cloud/cmd/rspscaled \
		scale.ropsoft.cloud/cmd/rspscale \
		scale.ropsoft.cloud/cmd/derper \
		scale.ropsoft.cloud/cmd/k8s-operator \
		scale.ropsoft.cloud/cmd/stund

depaware: ## Run depaware checks
	# depaware (via x/tools/go/packages) shells back to "go", so make sure the "go"
	# it finds in its $$PATH is the right one.
	PATH="$$(./tool/go env GOROOT)/bin:$$PATH" ./tool/go run github.com/tailscale/depaware --check \
		scale.ropsoft.cloud/cmd/rspscaled \
		scale.ropsoft.cloud/cmd/rspscale \
		scale.ropsoft.cloud/cmd/derper \
		scale.ropsoft.cloud/cmd/k8s-operator \
		scale.ropsoft.cloud/cmd/stund

buildwindows: ## Build rspscale CLI for windows/amd64
	GOOS=windows GOARCH=amd64 ./tool/go install scale.ropsoft.cloud/cmd/rspscale scale.ropsoft.cloud/cmd/rspscaled

build386: ## Build rspscale CLI for linux/386
	GOOS=linux GOARCH=386 ./tool/go install scale.ropsoft.cloud/cmd/rspscale scale.ropsoft.cloud/cmd/rspscaled

buildlinuxarm: ## Build rspscale CLI for linux/arm
	GOOS=linux GOARCH=arm ./tool/go install scale.ropsoft.cloud/cmd/rspscale scale.ropsoft.cloud/cmd/rspscaled

buildwasm: ## Build rspscale CLI for js/wasm
	GOOS=js GOARCH=wasm ./tool/go install ./cmd/tsconnect/wasm ./cmd/rspscale/cli

buildplan9:
	GOOS=plan9 GOARCH=amd64 ./tool/go install ./cmd/rspscale ./cmd/rspscaled

buildlinuxloong64: ## Build rspscale CLI for linux/loong64
	GOOS=linux GOARCH=loong64 ./tool/go install scale.ropsoft.cloud/cmd/rspscale scale.ropsoft.cloud/cmd/rspscaled

buildmultiarchimage: ## Build (and optionally push) multiarch docker image
	./build_docker.sh

check: staticcheck vet depaware buildwindows build386 buildlinuxarm buildwasm ## Perform basic checks and compilation tests

staticcheck: ## Run staticcheck.io checks
	./tool/go run honnef.co/go/tools/cmd/staticcheck -- $$(./tool/go list ./... | grep -v tempfork)

kube-generate-all: kube-generate-deepcopy ## Refresh generated files for Rspscale Kubernetes Operator
	./tool/go generate ./cmd/k8s-operator

# Rspscale operator watches Connector custom resources in a Kubernetes cluster
# and caches them locally. Caching is done implicitly by controller-runtime
# library (the middleware used by Rspscale operator to create kube control
# loops). When a Connector resource is GET/LIST-ed from within our control loop,
# the request goes through the cache. To ensure that cache contents don't get
# modified by control loops, controller-runtime deep copies the requested
# object. In order for this to work, Connector must implement deep copy
# functionality so we autogenerate it here.
# https://github.com/kubernetes-sigs/controller-runtime/blob/v0.16.3/pkg/cache/internal/cache_reader.go#L86-L89
kube-generate-deepcopy: ## Refresh generated deepcopy functionality for Rspscale kube API types
	./scripts/kube-deepcopy.sh

spk: ## Build synology package for ${SYNO_ARCH} architecture and ${SYNO_DSM} DSM version
	./tool/go run ./cmd/dist build synology/dsm${SYNO_DSM}/${SYNO_ARCH}

spkall: ## Build synology packages for all architectures and DSM versions
	./tool/go run ./cmd/dist build synology

pushspk: spk ## Push and install synology package on ${SYNO_HOST} host
	echo "Pushing SPK to root@${SYNO_HOST} (env var SYNO_HOST) ..."
	scp rspscale.spk root@${SYNO_HOST}:
	ssh root@${SYNO_HOST} /usr/syno/bin/synopkg install rspscale.spk

publishdevimage: ## Build and publish rspscale image to location specified by ${REPO}
	@test -n "${REPO}" || (echo "REPO=... required; e.g. REPO=ghcr.io/${USER}/rspscale" && exit 1)
	@test "${REPO}" != "ropsoft7/rspscale" || (echo "REPO=... must not be ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "ghcr.io/ropsoft7/rspscale" || (echo "REPO=... must not be ghcr.io/ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "rspscale/k8s-operator" || (echo "REPO=... must not be rspscale/k8s-operator" && exit 1)
	@test "${REPO}" != "ghcr.io/rspscale/k8s-operator" || (echo "REPO=... must not be ghcr.io/rspscale/k8s-operator" && exit 1)
	TAGS="${TAGS}" REPOS=${REPO} PLATFORM=${PLATFORM} PUSH=true TARGET=client ./build_docker.sh

publishdevoperator: ## Build and publish k8s-operator image to location specified by ${REPO}
	@test -n "${REPO}" || (echo "REPO=... required; e.g. REPO=ghcr.io/${USER}/rspscale" && exit 1)
	@test "${REPO}" != "ropsoft7/rspscale" || (echo "REPO=... must not be ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "ghcr.io/ropsoft7/rspscale" || (echo "REPO=... must not be ghcr.io/ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "rspscale/k8s-operator" || (echo "REPO=... must not be rspscale/k8s-operator" && exit 1)
	@test "${REPO}" != "ghcr.io/rspscale/k8s-operator" || (echo "REPO=... must not be ghcr.io/rspscale/k8s-operator" && exit 1)
	TAGS="${TAGS}" REPOS=${REPO} PLATFORM=${PLATFORM} PUSH=true TARGET=k8s-operator ./build_docker.sh

publishdevnameserver: ## Build and publish k8s-nameserver image to location specified by ${REPO}
	@test -n "${REPO}" || (echo "REPO=... required; e.g. REPO=ghcr.io/${USER}/rspscale" && exit 1)
	@test "${REPO}" != "ropsoft7/rspscale" || (echo "REPO=... must not be ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "ghcr.io/ropsoft7/rspscale" || (echo "REPO=... must not be ghcr.io/ropsoft7/rspscale" && exit 1)
	@test "${REPO}" != "rspscale/k8s-nameserver" || (echo "REPO=... must not be rspscale/k8s-nameserver" && exit 1)
	@test "${REPO}" != "ghcr.io/rspscale/k8s-nameserver" || (echo "REPO=... must not be ghcr.io/rspscale/k8s-nameserver" && exit 1)
	TAGS="${TAGS}" REPOS=${REPO} PLATFORM=${PLATFORM} PUSH=true TARGET=k8s-nameserver ./build_docker.sh

.PHONY: sshintegrationtest
sshintegrationtest: ## Run the SSH integration tests in various Docker containers
	@GOOS=linux GOARCH=amd64 ./tool/go test -tags integrationtest -c ./ssh/tailssh -o ssh/tailssh/testcontainers/tailssh.test && \
	GOOS=linux GOARCH=amd64 ./tool/go build -o ssh/tailssh/testcontainers/rspscaled ./cmd/rspscaled && \
	echo "Testing on ubuntu:focal" && docker build --build-arg="BASE=ubuntu:focal" -t ssh-ubuntu-focal ssh/tailssh/testcontainers && \
	echo "Testing on ubuntu:jammy" && docker build --build-arg="BASE=ubuntu:jammy" -t ssh-ubuntu-jammy ssh/tailssh/testcontainers && \
	echo "Testing on ubuntu:mantic" && docker build --build-arg="BASE=ubuntu:mantic" -t ssh-ubuntu-mantic ssh/tailssh/testcontainers && \
	echo "Testing on ubuntu:noble" && docker build --build-arg="BASE=ubuntu:noble" -t ssh-ubuntu-noble ssh/tailssh/testcontainers && \
	echo "Testing on alpine:latest" && docker build --build-arg="BASE=alpine:latest" -t ssh-alpine-latest ssh/tailssh/testcontainers

help: ## Show this help
	@echo "\nSpecify a command. The choices are:\n"
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-20s\033[m %s\n", $$1, $$2}'
	@echo ""
.PHONY: help

.DEFAULT_GOAL := help
