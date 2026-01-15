# Variables
BINARY_NAME=main_wasm.wasm
BUILD_DIR=build
SRC_FILE=cmd/wasm/main_wasm.go

# Déclaration des cibles qui ne sont pas des fichiers
.PHONY: all build clean copy-glue

# Cible par défaut
all: build

# Compilation en WebAssembly
build:
	@echo "Création du dossier de build..."
	mkdir -p $(BUILD_DIR)
	@echo "Compilation de $(SRC_FILE) vers $(BUILD_DIR)/$(BINARY_NAME)..."
	GOOS=js GOARCH=wasm go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_FILE)
	@echo "Compilation terminée avec succès."

# Optionnel : Copier le fichier "glue" JavaScript nécessaire pour exécuter le WASM
copy-glue: build
	@echo "Copie de wasm_exec.js dans $(BUILD_DIR)..."
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" $(BUILD_DIR)/

# Nettoyage du dossier build
clean:
	@echo "Suppression du dossier $(BUILD_DIR)..."
	rm -rf $(BUILD_DIR)