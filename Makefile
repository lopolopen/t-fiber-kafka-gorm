.PHONY: test swag wire dev

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

dev:
	cd ./cmd/api && go run .

swagdev: swag dev