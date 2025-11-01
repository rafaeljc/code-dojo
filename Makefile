.PHONY: help all install build clean test rebuild fmt

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: clean build install ## Clean, build, and install the CLI
	@echo ""
	@echo "🎉 All done! The dojo CLI is ready to use."
	@if echo "$$PATH" | grep -q "$$(go env GOPATH)/bin"; then \
		echo "✓ You can now run: dojo --help"; \
	else \
		echo "⚠️  Run this to use it: export PATH=\"\$$HOME/go/bin:\$$PATH\" && dojo --help"; \
	fi

install: ## Install dojo CLI to $GOPATH/bin
	@cd tools/dojo-cli && go install ./cmd/dojo
	@echo "✓ dojo CLI installed to $$(go env GOPATH)/bin"
	@if echo "$$PATH" | grep -q "$$(go env GOPATH)/bin"; then \
		echo "✓ You can now run: dojo --help"; \
	else \
		echo "⚠️  Run this to use it: export PATH=\"\$$HOME/go/bin:\$$PATH\" && dojo --help"; \
		echo "   (Or add to ~/.bashrc permanently)"; \
	fi

build: ## Build dojo CLI to bin/dojo
	@mkdir -p bin
	@cd tools/dojo-cli && go build -o ../../bin/dojo ./cmd/dojo
	@echo "✓ Built bin/dojo"

clean: ## Remove built binaries
	@rm -rf bin/
	@echo "✓ Cleaned bin/"

test: ## Run tests
	@cd tools/dojo-cli && go test ./...

rebuild: clean build ## Clean and rebuild local binary
	@echo "✓ Rebuild complete"

fmt: ## Format Go code
	@cd tools/dojo-cli && go fmt ./...
	@echo "✓ Code formatted"
