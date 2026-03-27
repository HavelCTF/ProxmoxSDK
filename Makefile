# -- GLOBAL VARIABLES --
TS_DIR=pkg/typescript
WASM_BINARY=main_wasm.wasm
WASM_MAIN=cmd/wasm/main_wasm.go
WASM_OUT=$(TS_DIR)/dist/$(WASM_BINARY)

# -- CI DETECTION --
# Automatically use 'npm ci' in GitHub Actions, and 'npm install' locally
NPM_CMD=npm install
ifeq ($(CI),true)
	NPM_CMD=npm ci
endif

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

# ==============================================================================
#  MAIN TARGETS
# ==============================================================================

all: build

# La cible 'build' publique génère maintenant le SDK complet et utilisable !
build: deps copy-glue build-ts build-wasm
	$(call print_success,Full SDK build completed successfully)

# ==============================================================================
#  BUILD STEPS
# ==============================================================================

build-wasm:
	$(call print_header,BUILDING WASM)
	@printf "$(CYAN)[1/2]$(NC) $(BOLD)$(BLUE)Creating dist directory...$(NC)"
	@mkdir -p $(TS_DIR)/dist
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[2/2]$(NC) $(BOLD)$(BLUE)Compiling Go WASM directly to $(WASM_OUT)...$(NC)"
	@if GOOS=js GOARCH=wasm go build -o $(WASM_OUT) $(WASM_MAIN); \
	then \
		printf " $(GREEN)[OK]$(NC)\n"; \
	else \
		printf " $(RED)[FAILED]$(NC)\n"; exit 1; \
	fi

copy-glue:
	@printf "$(CYAN)[+]$(NC) $(BOLD)$(BLUE)Copying wasm_exec.js to $(TS_DIR)/src/wasm...$(NC)"
	@mkdir -p $(TS_DIR)/src/wasm
	@if cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(TS_DIR)/src/wasm/; then \
		printf " $(GREEN)[OK]$(NC)\n"; \
	else \
		printf " $(RED)[FAILED]$(NC)\n"; exit 1; \
	fi

build-ts:
	@printf "$(CYAN)[1/2]$(NC) $(BOLD)$(BLUE)Compiling TypeScript...$(NC)\n"
	@cd $(TS_DIR) && npm run build
	@printf "$(CYAN)[2/2]$(NC) $(BOLD)$(BLUE)Copying wasm_exec.js to dist/wasm...$(NC)\n"
	@mkdir -p $(TS_DIR)/dist/wasm
	@cp $(TS_DIR)/src/wasm/wasm_exec.js $(TS_DIR)/dist/wasm/
	@printf " $(GREEN)[OK]$(NC)\n"

# ==============================================================================
#  DEPENDENCIES
# ==============================================================================

deps:
	$(call print_header,DEPENDENCIES)
	@printf "$(CYAN)[1/3]$(NC) $(BOLD)$(BLUE)Downloading Go modules...$(NC)"
	@go mod download || (printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[2/3]$(NC) $(BOLD)$(BLUE)Tidying Go modules...$(NC)"
	@go mod tidy || (printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"
	@printf "$(CYAN)[3/3]$(NC) $(BOLD)$(BLUE)Installing NPM modules ($(NPM_CMD))...$(NC)\n"
	@cd $(TS_DIR) && $(NPM_CMD) || (printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"

# ==============================================================================
#  TESTING & FORMATTING
# ==============================================================================

test:
	$(call print_header,TESTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go test...$(NC)\n"
	@go test -v ./... || exit 1
	$(call print_success,All tests passed)

fmt-check:
	$(call print_header,FORMATTING CHECK)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Checking go fmt...$(NC)"
	@test -z "$$(gofmt -l .)" || (printf " $(RED)[FAILED]$(NC) Unformatted files found\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"

fmt:
	$(call print_header,FORMATTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go fmt...$(NC)"
	@go fmt ./... || (printf " $(RED)[FAILED]$(NC)\n" && exit 1)
	@printf " $(GREEN)[OK]$(NC)\n"

# ==============================================================================
#  LINTING
# ==============================================================================

lint: deps lint-go lint-ts
	$(call print_success,All Linting passed)

lint-go: deps
	$(call print_header,GO LINTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running go vet...$(NC)\n"
	@go vet ./... || exit 1
	$(call print_success,Go linting passed)

lint-ts: deps
	$(call print_header,TYPESCRIPT LINTING)
	@printf "$(CYAN)[1/1]$(NC) $(BOLD)$(BLUE)Running TypeScript linter...$(NC)\n"
	@cd $(TS_DIR) && npm run lint || exit 1
	$(call print_success,TypeScript linting passed)

# ==============================================================================
#  CLEANUP & CI
# ==============================================================================

clean:
	$(call print_header,CLEANUP)
	@rm -rf $(TS_DIR)/dist $(TS_DIR)/node_modules $(TS_DIR)/src/wasm/wasm_exec.js coverage.*
	@printf " $(GREEN)[OK] Clean completed$(NC)\n"

ci: lint fmt-check test build
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
	@printf "  $(CYAN)build$(NC)          Build the full SDK (TS + WASM) (default)\n"
	@printf "  $(CYAN)build-wasm$(NC)     Compile only the Go code to WASM\n"
	@printf "  $(CYAN)build-ts$(NC)       Compile only the TypeScript code\n"
	@printf "  $(CYAN)test$(NC)           Run all Go tests\n"
	@printf "  $(CYAN)fmt-check$(NC)      Check if Go code is formatted properly\n"
	@printf "  $(CYAN)fmt$(NC)            Format Go code\n"
	@printf "  $(CYAN)lint$(NC)           Run Go and TypeScript linters (use lint-go or lint-ts for individual checks)\n"
	@printf "  $(CYAN)deps$(NC)           Download Go modules and NPM dependencies\n"
	@printf "  $(CYAN)clean$(NC)          Remove build artifacts (dist, node_modules)\n"
	@printf "  $(CYAN)ci$(NC)             Run lint, test, and build (for CI/CD)\n"
	@printf "  $(CYAN)help$(NC)           Show this help message\n"

.PHONY: all build build-wasm copy-glue build-ts test fmt-check fmt lint deps clean ci help