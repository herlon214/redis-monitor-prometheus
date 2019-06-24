################################
# STEP 1 build executable binary
################################
FROM golang:1.12-alpine as builder

WORKDIR $GOPATH/src/github.com/herlon214/redis-monitor-prometheus
COPY . .

# Install git for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Fetch dependencies
RUN GO111MODULE=on go mod download

# Compile with ldflags that will omit all symbol information and make it link statically
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -extldflags "-static"' -o /go/bin/redis-monitor-prometheus ./src/main.go

################################
# STEP 2 final lightweight image
################################
FROM alpine:3.10
RUN apk --update add redis

# Copy the binary from builder
COPY --from=builder /go/bin/redis-monitor-prometheus /usr/bin/redis-monitor-prometheus

# Execute
ENTRYPOINT [ "/usr/bin/redis-monitor-prometheus" ]