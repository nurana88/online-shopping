
## up	- Boots up docker dev env
up:
	@echo "$(OK_COLOR)==> Booting up docker environment$(NO_COLOR)"
	docker-compose up -d

stop:
	@printf "$(OK_COLOR)==> Stopping docker containers$(NO_COLOR)\n"
	@docker-compose down

test-unit:
	@printf "$(OK_COLOR)==> Running unit tests$(NO_COLOR)\n"
	@go test -coverprofile=coverage_unit.txt -covermode=atomic

build:
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BINARY)

.PHONY: up stop test-unit build
