.PHONY: test swag wire run gen build name

env ?= local

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

run:
	go mod tidy
	cd ./cmd/api && go run . -f etc/$(env).yaml

swagdev: swag run

REGISTRY   := docker.io
tag        ?= latest

IMAGE_BASE := <org-name>/<app-name>
IMAGE_FULL := $(IMAGE_BASE):$(tag)
LATEST_IMG := $(IMAGE_BASE):latest

DEST_VERSION := docker://$(REGISTRY)/$(IMAGE_FULL)
DEST_LATEST  := docker://$(REGISTRY)/$(LATEST_IMG)

build:
	-podman manifest rm $(IMAGE_BASE):local
	podman build \
		--platform linux/amd64,linux/arm64 \
		--manifest $(IMAGE_BASE):local \
		.
	podman manifest push $(IMAGE_BASE):local $(DEST_VERSION)
	@if [ "$(tag)" != "latest" ]; then \
		echo "Pushing latest tag..."; \
		podman manifest push $(IMAGE_BASE):local $(DEST_LATEST); \
	fi
	podman manifest rm $(IMAGE_BASE):local


# ===== Can be deleted when the initialization is finished. =====
name:
	@if [ -z "$(org)" ] || [ -z "$(app)" ]; then \
		echo "Usage: make name org=<org-name> app=<app-name>"; \
		exit 1; \
	fi
	chmod +x name.sh
	@./name.sh $(org) $(app)
# ===============================================================