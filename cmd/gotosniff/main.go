package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/tangfeixiong/go-to-docker/cmd"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	basename := filepath.Base(os.Args[0])
	command := cmd.CommandFor(basename)
	// command := cmd.NewRootCommand(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
