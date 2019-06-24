build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o dist/linux-amd64/kubeless-yaml
	zip -r dist/linux-amd64.zip dist/linux-amd64/*
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o dist/darwin-amd64/kubeless-yaml
	zip -r dist/darwin-amd64.zip dist/darwin-amd64/*