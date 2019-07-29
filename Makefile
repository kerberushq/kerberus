SHELL := /bin/bash
GITCOMMIT=$(shell git describe --tags HEAD)$(shell [[ $$(git status --porcelain) = "" ]] || echo -dirty)
LDFLAGS="-X main.gitCommit=$(GITCOMMIT)"
SDK_RELEASE_VERSION=v0.8.1

IMAGE ?= quay.io/kerberus/kerberus:$(GITCOMMIT)
OPERATOR_NAME ?= "kerberus. operator"

.PHONY: build
build: generate manager image

.PHONY: image
image:
	operator-sdk build $(IMAGE) --image-builder docker

.PHONY: image-push
image-push:
	docker push $(IMAGE)

.PHONY: generate
generate:
	go generate ./...
	./hack/update-codegen.sh

.PHONY: generate-internal
generate-internal:
	./hack/update-codegen-internal.sh

.PHONY: manager
manager:
	go build -o build/_output/bin/kerberus -ldflags ${LDFLAGS} ./cmd/$@ 

.PHONY: up
up:
	export KUBECONFIG=build/admin.kubeconfig
	operator-sdk up local

.PHONY: minikube
minikube:
	export GO111MODULE=on
	export KUBECONFIG=build/admin.kubeconfig
	minikube start

setup:
	curl -JL -o /usr/local/bin/operator-sdk https://github.com/operator-framework/operator-sdk/releases/download/${SDK_RELEASE_VERSION}/operator-sdk-${SDK_RELEASE_VERSION}-x86_64-linux-gnu

# workaround pruning of non go files
# https://github.com/golang/go/issues/26366
# all this below is hack locking to particula gopath config.
# Why this have to be so hard?
versionPath=$(shell go list -f {{.Dir}} k8s.io/code-generator/cmd/client-gen)
gogoprotosTarget:=./vendor/k8s.io/code-generator/

.PHONY: vendor
vendor:
	export GO111MODULE=on
	@go mod tidy -v
	@go mod vendor -v
	@go mod verify
	# work around to save external files in code-generator
	@mkdir -p $(gogoprotosTarget)
	@cp -rf --no-preserve=mode $(versionPath)/../../ $(gogoprotosTarget)
	@chmod 755 -R $(gogoprotosTarget)
	@export GO111MODULE=off
	@go get -v k8s.io/code-generator
