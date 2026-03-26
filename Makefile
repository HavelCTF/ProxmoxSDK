# -- GLOBAL VARIABLES --
BUILD_DIR=build

# -- TYPESCRIPT VARIABLES --
TS_DIR=pkg/typescript

# -- WASM VARIABLES --
WASM_BINARY=main_wasm.wasm
WASM_MAIN=cmd/wasm/main_wasm.go

# -- COLORS --
GREEN=\033[0;32m
BLUE=\033[0;34m
YELLOW=\033[0;33m
CYAN=\033[0;36m
RED=\033[0;31m
NC=\033[0m
BOLD=\033[1m

# -- BOX DRAWING --
define print_header
	@printf "\n$(BOLD)$(CYAN)+-------------------------------------------+$(NC)\n"
	@printf "$(BOLD)$(CYAN)|$(NC) %-42s$(BOLD)$(CYAN)|$(NC)\n" "$(1)"
	@printf "$(BOLD)$(CYAN)+-------------------------------------------+$(NC)\n\n"
endef

define print_success
	@printf "\n$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"
	@printf "$(BOLD)$(GREEN)|$(NC) [OK] %-37s$(BOLD)$(GREEN)|$(NC)\n" "$(1)"
	@printf "$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"
endef

define print_error
	@printf "\n$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"
	@printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "$(1)"
	@printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"
endef

# ==============================================================================
#  MAIN TARGETS
# ==============================================================================

all: build copy-glue build-ts
	$(call print_success,Build completed successfully)

# ==============================================================================
#  BUILD
# ==============================================================================

build:
	$(call print_header,BUILD)
	@printf "$(CYAN)[1/2]$(NC) $(BOLD)$(BLUE)Creating build directory...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[2/2]$(NC) $(BOLD)$(BLUE)Compiling $(WASM_MAIN) to $(BUILD_DIR)/$(WASM_BINARY)...$(NC)"
	@if GOOS=js GOARCH=wasm go build -o $(BUILD_DIR)/$(WASM_BINARY) $(WASM_MAIN); \
	then \
		printf " $(GREEN)[OK]$(NC)\n"; \
	else \
		printf " $(RED)[FAILED]$(NC)\n"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "WASM compilation failed"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		exit 1; \
	fi

copy-glue: build
	@printf "$(CYAN)[+]$(NC) $(BOLD)$(BLUE)Copying wasm_exec.js to $(BUILD_DIR)...$(NC)"
	@if cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(BUILD_DIR)/; then \
		printf " $(GREEN)[OK]$(NC)\n"; \
	else \
		printf " $(RED)[FAILED]$(NC)\n"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "Failed to copy wasm_exec.js"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		exit 1; \
	fi

build-ts:
	@printf "$(CYAN)[1/2]$(NC) $(BOLD)$(BLUE)Installing NPM dependencies...$(NC)\n"
	@cd $(TS_DIR) && npm install
	@printf "$(CYAN)[2/2]$(NC) $(BOLD)$(BLUE)Compiling TypeScript...$(NC)\n"
	@cd $(TS_DIR) && npm run build
	@printf " $(GREEN)[OK]$(NC)\n"


# ==============================================================================
#  TESTING
# ==============================================================================

test:
	$(call print_header,TESTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go test...$(NC)\n"
	@if go test -v ./... ; \
	then \
		printf "\n$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(GREEN)|$(NC) [OK] %-37s$(BOLD)$(GREEN)|$(NC)\n" "All tests passed"; \
		printf "$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"; \
	else \
		printf "\n$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "Some tests failed"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		exit 1; \
	fi

test-coverage:
	$(call print_header,TEST COVERAGE)
	@printf "$(CYAN)[1/2]$(NC) $(BOLD)$(BLUE)Running tests with coverage...$(NC)\n"
	@if go test -covermode=count -coverpkg=./... -coverprofile=coverage.out -v ./... ; then \
		printf " $(GREEN)[OK]$(NC)\n"; \
	else \
		printf "\n$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "Tests failed during coverage"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		exit 1; \
	fi
	@printf "$(CYAN)[2/2]$(NC) $(BOLD)$(BLUE)Generating coverage report...$(NC)"
	@go tool cover -html=coverage.out -o coverage.html || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	$(call print_success,Coverage report: coverage.html)

# ==============================================================================
#  LINTING & FORMATTING
# ==============================================================================

lint:
	$(call print_header,LINTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go vet...$(NC)\n"
	@if go vet ./... 2>&1; \
	then \
		printf "\n$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(GREEN)|$(NC) [OK] %-37s$(BOLD)$(GREEN)|$(NC)\n" "Linting passed"; \
		printf "$(BOLD)$(GREEN)+-------------------------------------------+$(NC)\n"; \
	else \
		printf "\n$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		printf "$(BOLD)$(RED)|$(NC) [FAILED] %-33s$(BOLD)$(RED)|$(NC)\n" "Linting errors found"; \
		printf "$(BOLD)$(RED)+-------------------------------------------+$(NC)\n"; \
		exit 1; \
	fi

fmt-check:
	$(call print_header,FORMATTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Checking go fmt...$(NC)"
	@test -z "$$(gofmt -l .)" || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	$(call print_success,Code formatted)

fmt:
	$(call print_header,FORMATTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go fmt...$(NC)"
	@go fmt ./... || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	$(call print_success,Code formatted)

# ==============================================================================
#  DEPENDENCIES
# ==============================================================================

deps:
	$(call print_header,DEPENDENCIES)
	@printf "$(CYAN)[1/3]$(NC) $(BOLD)$(BLUE)Downloading Go modules...$(NC)"
	@go mod download || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[2/3]$(NC) $(BOLD)$(BLUE)Tidying Go modules...$(NC)"
	@go mod tidy || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[3/3]$(NC) $(BOLD)$(BLUE)Installing NPM modules...$(NC)\n"
	@cd $(TS_DIR) && npm install || \
		(printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	$(call print_success,Dependencies installed)

# ==============================================================================
#  CLEANUP
# ==============================================================================

clean:
	$(call print_header,CLEANUP)
	@printf "$(CYAN)[1/3]$(NC) $(BOLD)$(YELLOW)Removing $(BUILD_DIR) directory...$(NC)"
	@rm -rf $(BUILD_DIR)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[2/3]$(NC) $(BOLD)$(YELLOW)Removing coverage files...$(NC)"
	@rm -f coverage.out coverage.html
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[3/3]$(NC) $(BOLD)$(YELLOW)Removing node_modules...$(NC)"
	@rm -rf $(TS_DIR)/node_modules $(TS_DIR)/dist
	@printf " $(GREEN)[OK]$(NC)\n"
	$(call print_success,Clean completed)

# ==============================================================================
#  CI/CD
# ==============================================================================

ci: lint test build
	$(call print_success,CI pipeline completed successfully)

# ==============================================================================
#  HELP
# ==============================================================================

help:
	@printf "\n$(BOLD)$(CYAN)+-------------------------------------------+$(NC)\n"
	@printf "$(BOLD)$(CYAN)|$(NC) %-42s$(BOLD)$(CYAN)|$(NC)\n" "ProxmoxSDK Build System"
	@printf "$(BOLD)$(CYAN)+-------------------------------------------+$(NC)\n\n"
	@printf "$(BOLD)Usage:$(NC) make $(CYAN)<target>$(NC)\n\n"
	@printf "$(BOLD)Targets:$(NC)\n"
	@printf "  $(CYAN)all$(NC)            Build the WASM binary (default)\n"
	@printf "  $(CYAN)build$(NC)          Compile Go to WASM\n"
	@printf "  $(CYAN)test$(NC)           Run all tests\n"
	@printf "  $(CYAN)test-coverage$(NC)  Run tests with coverage report\n"
	@printf "  $(CYAN)lint$(NC)           Run go vet linter\n"
	@printf "  $(CYAN)fmt$(NC)            Format code with go fmt\n"
	@printf "  $(CYAN)deps$(NC)           Download Go modules and NPM dependencies\n"
	@printf "  $(CYAN)clean$(NC)          Remove build artifacts (including node_modules)\n"
	@printf "  $(CYAN)ci$(NC)             Run lint, test, and build (for CI/CD)\n"
	@printf "  $(CYAN)help$(NC)           Show this help message\n"

.PHONY: all build copy-glue build-ts test test-coverage fmt lint deps clean ci help