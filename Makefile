.PHONY: test swag wire run gen build name

APP_NAME		:= <app-name>
ORG_NAME		:= <org-name>
IMAGE_BASE 		:= ${ORG_NAME}/${APP_NAME}
REGISTRY   		:= docker.io
CONTAINER_TOOL	:= docker
GIT_SHORT_SHA 	:= $(shell git describe --always --dirty 2>/dev/null || echo "unknown")

COMMA := ,
EMPTY :=
SPACE := $(EMPTY) $(EMPTY)

SWAG_DIRS := ./ \
../../internal/adapters/http/handlers/v1 \
../../internal/adapters/http/dto \
../../internal/applic/query \
../../internal/applic/cmd \
../../internal/applic/result \
../../internal/domain/enum

SWAG_DIRS_COMMA := $(subst $(SPACE),,$(foreach d,$(SWAG_DIRS),$(d)$(COMMA)))
SWAG_DIRS_COMMA := $(SWAG_DIRS_COMMA:%$(COMMA)=%)

wire:
	cd ./cmd/api && go tool wire

swag:
	cd ./internal/adapters/http/handlers && go tool swag fmt
	cd ./cmd/api && go tool swag init --dir $(SWAG_DIRS_COMMA)

gen:
	go generate ./...

GOPATH=$(shell go env GOPATH)

proto:
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
	    --proto_path=. \
		--go_out=. \
		--go_opt=paths=source_relative \
		pkg/schema/*.proto

env ?= local

run:
	go mod tidy
	cd ./cmd/api && COMMIT_SHA=${GIT_SHORT_SHA}  go run . -f etc/$(env).yaml

tag 		?= ${GIT_SHORT_SHA}
IMAGE_FULL 	:= $(REGISTRY)/$(IMAGE_BASE):$(tag)
IMAGE_URL 	:= docker://$(IMAGE_FULL)

build:
	${CONTAINER_TOOL} build \
		--build-arg COMMIT_SHA=${GIT_SHORT_SHA} \
		--platform linux/amd64 \
		--tag ${IMAGE_FULL} \
		.

push:
	${CONTAINER_TOOL} push $(IMAGE_FULL) $(IMAGE_URL)

# ===== Can be deleted when the initialization is finished. =====
name:
	@if [ -z "$(org)" ] || [ -z "$(app)" ]; then \
		echo "Usage: make name org=<org-name> app=<app-name>"; \
		exit 1; \
	fi
	chmod +x name.sh
	@./name.sh $(org) $(app)
# ===============================================================