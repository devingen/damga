.PHONY: build
build:
	go build -o ./bin/damga .

.PHONY: build-to-tools
build-to-tools:
	go build -o "$DEVINGEN_REPOS/tools/bin/damga" .

# Define the output directory
OUTPUT_DIR := release

# Check if VERSION is provided
ifndef VERSION
$(error VERSION is not defined. Use 'make archive-release VERSION=<version>' to provide it.)
endif

# Build for multiple platforms with version
.PHONY: build-release
build-release:
	mkdir -p $(OUTPUT_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/damga-linux-amd64-$(VERSION) # Linux
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/damga-windows-amd64-$(VERSION).exe # Windows
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/damga-darwin-amd64-$(VERSION) # macOS

# Archive the binaries into .tar.gz and .zip files
.PHONY: archive-release
archive-release: clean-release build-release
	tar -czvf $(OUTPUT_DIR)/damga-linux-amd64-$(VERSION).tar.gz -C $(OUTPUT_DIR) damga-linux-amd64-$(VERSION)
	zip $(OUTPUT_DIR)/damga-windows-amd64-$(VERSION).zip $(OUTPUT_DIR)/damga-windows-amd64-$(VERSION).exe
	tar -czvf $(OUTPUT_DIR)/damga-darwin-amd64-$(VERSION).tar.gz -C $(OUTPUT_DIR) damga-darwin-amd64-$(VERSION)

.PHONY: clean-release
clean-release:
	rm -rf $(OUTPUT_DIR)