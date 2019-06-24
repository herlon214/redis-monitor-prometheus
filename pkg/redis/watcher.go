package redis

import (
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	// redisURI is used to connect to redis
	redisURI = os.Getenv("REDIS_URI")
)

// Writer is the common io interface
type Writer interface {
	Write(line []byte) (n int, err error)
}

// Watcher is responsible for reading cmd output and return as string
type Watcher struct {
	Writer Writer
}

// Run spawns a "redis-cli monitor" process
func (w *Watcher) Run() {
	cmd := exec.Command("redis-cli", "-u", redisURI, "monitor")

	mw := io.MultiWriter(w.Writer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
