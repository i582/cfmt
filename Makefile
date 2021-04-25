check:
	@echo "running tests..."
	@go test -count 1 -v ./...
	@echo "everything is OK"

.PHONY: check
