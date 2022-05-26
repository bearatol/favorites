PROJECTNAME=$(shell basename "$(PWD)")
DIR=$(shell pwd)
PROTO=$(shell pwd)/proto/favorites
GOLANGCI_LINT=$(shell which golangci-lint run)
DATABASE="postgres://favorites:favorites@localhost:5444/favorites?sslmode=disable"

.PHONY: migrations
## run: start the project
run: migrations
	@go run $(DIR)/cmd/favorites/main.go -config_file="config.local.yml" -config_dir="$(DIR)/configs"

local-pg:
	@mkdir -p $(DIR)/db_data
	@docker-compose -f $(DIR)/docker/docker-compose.local.yaml up -d

## lint: start linter for the project
lint:
	@echo " > Start golang linter"
	@$(GOLANGCI_LINT) run --go=1.17 --timeout 5m0s

go-test:
	@echo " > Testing ..."
	@go test -v $(DIR)/...

go-mod-download:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod download

go-mod-tidy:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

migrations:
	@goose -dir $(DIR)/migrations postgres $(DATABASE) up

## protoc: generate go files from a protobuf
protoc:
	rm -rf $(PROTO)/gen && mkdir $(PROTO)/gen && mkdir $(PROTO)/gen/swagger
	protoc -I $(PROTO) \
	--go_out $(PROTO)/gen --go_opt paths=source_relative \
	--go-grpc_out $(PROTO)/gen --go-grpc_opt paths=source_relative \
	--grpc-gateway_out $(PROTO)/gen --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt allow_delete_body=true \
	--openapiv2_out $(PROTO)/gen/swagger \
	--openapiv2_opt logtostderr=true \
	--openapiv2_opt allow_delete_body=true \
	favorites.proto

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo