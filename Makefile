.PHONY: help
VETARGS?=-all
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

tools: ## Acquire tools for building the project
	@echo "Grabbing tools for the project"
	@go get -u github.com/kardianos/govendor

vet: ## vet runs the Go source code static analysis tool `vet` to find any common errors.
	@echo "go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) $$(ls -d */ | grep -v vendor) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt: ## Run gofmt and changes back to the files
	@gofmt -w $(GOFMT_FILES)

fmtcheck: ## Run a check to show what files gofmt will change
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"
