.PHONY: test doc wire debug clean run gen build name tidy proto push

APP_NAME		:= <app-name>
ORG_NAME		:= <org-name>
IMAGE_BASE 		:= ${ORG_NAME}/${APP_NAME}
REGISTRY   		:= docker.io
GIT_SHORT_SHA 	:= $(shell git describe --always --dirty 2>/dev/null || echo "unknown")

wire:
	@cd ./cmd/api && go tool wire

doc:
	@cd ./internal/adapters/http/handlers && go tool swag fmt
	@go tool swag init -g ./cmd/api/main.go --parseInternal

gen:
	go generate ./...

GOPATH=$(shell go env GOPATH)

proto:
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go \
	    --proto_path=. \
		--go_out=. \
		--go_opt=paths=source_relative \
		pkg/schema/*.proto

MIGRATE_CLI		:= go run -modfile=tools.mod -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate
DATABASE		:=
action			?= 

migrate:
	${MIGRATE_CLI} create -ext sql -dir ./db/migrations -format 20060102_150405 -tz "Asia/Shanghai" ${action}
migrate-up:
	${MIGRATE_CLI} -path ./db/migrations -database ${DATABASE} up
migrate-down:	
	${MIGRATE_CLI} -path ./db/migrations -database ${DATABASE} down

debug:
	docker compose up --no-build

debug-build:
	docker compose up --build

clean:
	docker compose down --volumes --remove-orphans
	docker image prune -f

env ?= local

run:
	go mod tidy
	cd ./cmd/api && APP_ENV=${env} go run -ldflags="-X main.commitSHA=${GIT_SHORT_SHA}" .

tag 		?= ${GIT_SHORT_SHA}
config		?= debug
IMAGE_FULL 	:= $(REGISTRY)/$(IMAGE_BASE):$(tag)
IMAGE_URL 	:= docker://$(IMAGE_FULL)

build:
	docker build \
		--build-arg COMMIT_SHA=${GIT_SHORT_SHA} \
		--build-arg BUILD_MODE=${config} \
		--platform linux/amd64 \
		--tag ${IMAGE_FULL} \
		.

push:
	docker push $(IMAGE_FULL) $(IMAGE_URL)

tidy:
	go mod tidy
	go mod tidy -modfile=tools.mod

# ===== Can be deleted when the initialization is finished. =====
name:
	@if [ -z "$(org)" ] || [ -z "$(app)" ]; then \
		echo "Usage: make name org=<org-name> app=<app-name>"; \
		exit 1; \
	fi
	chmod +x name.sh
	@./name.sh $(org) $(app)
# ===============================================================