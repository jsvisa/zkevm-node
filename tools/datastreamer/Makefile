# Check dependencies
# Check for Go
.PHONY: check-go
check-go:
	@which go > /dev/null || (echo "Error: Go is not installed" && exit 1)

# Targets that require the checks
generate-file: check-go
reprocess: check-go
decode-entry-offline: check-go
decode-l2block-offline: check-go
decode-entry: check-go
decode-l2block: check-go
truncate: check-go

arguments := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

.PHONY: generate-file
generate-file: ## Runs the tool to populate the binary file
	go run main.go generate -cfg config/tool.config.toml

.PHONY: decode-entry
decode-entry: ## Runs the tool to decode a given entry number
	go run main.go decode-entry -cfg config/tool.config.toml -entry $(arguments)

.PHONY: decode-l2block
decode-l2block: ## Runs the tool to decode a given L2 block
	go run main.go decode-l2block -cfg config/tool.config.toml -l2block $(arguments)

.PHONY: decode-entry-offline
decode-entry-offline: ## Runs the offline tool to decode a given entry number
	go run main.go decode-entry-offline -cfg config/tool.config.toml -entry $(arguments)

.PHONY: decode-l2block-offline
decode-l2block-offline: ## Runs the offline tool to decode a given L2 block
	go run main.go decode-l2block-offline -cfg config/tool.config.toml -l2block $(arguments)

.PHONY: truncate
truncate: ## Runs the offline tool to truncate the stream file
	go run main.go truncate -cfg config/tool.config.toml -entry $(arguments)

# .PHONY: reprocess
reprocess: ## Runs the tool to reprocess the information in the stream up to a given l2 block
	go run main.go reprocess -cfg config/tool.config.toml -genesis ../test/config/test.genesis.config.json -l2block $(arguments)

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.
.DEFAULT_GOAL := help

.PHONY: help
help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT: 
	@echo ""

