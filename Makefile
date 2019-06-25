build:
	@make test
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o dist/linux-amd64/redis-monitor-prometheus src/main.go
	@echo "Build binaries..."
	@zip -r dist/linux-amd64.zip dist/linux-amd64/*
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o dist/darwin-amd64/redis-monitor-prometheus src/main.go
	@zip -r dist/darwin-amd64.zip dist/darwin-amd64/*
	@echo "All done!"
test:
	@echo "Running all */**_test.go files..."
	@find .  -name "*_test.go" -print | xargs -n 1 go test