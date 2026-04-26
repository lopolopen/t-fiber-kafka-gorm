.PHONY: test swag wire run gen build name

comma := ,
empty :=
space := $(empty) $(empty)

SWAG_DIRS := ./ \
../../internal/adapters/http/handlers/v1 \
../../internal/adapters/http/dto \
../../internal/applic/query \
../../internal/applic/cmd \
../../internal/applic/result

SWAG_DIRS_COMMA := $(subst $(space),,$(foreach d,$(SWAG_DIRS),$(d)$(comma)))
SWAG_DIRS_COMMA := $(SWAG_DIRS_COMMA:%$(comma)=%)

wire:
	cd ./cmd/api && go tool wire

swag:
	cd ./internal/adapters/http/handlers && go tool swag fmt
	cd ./cmd/api && go tool swag init --dir $(SWAG_DIRS_COMMA)

gen:
	go generate ./...

env ?= local

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

name:
	chmod +x name.sh
	@./name.sh $(org) $(app)