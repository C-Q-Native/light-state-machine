package runtime

import (
	"context"
	_ "encoding/base64"
	_ "encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"net/http"

	logger "github.com/google/logger"
)

type SMRuntime struct {
	ctx           context.Context
	cancel        context.CancelFunc
	runtimeConfig *Config
}

func NewSMRuntime() *SMRuntime {
	ctx, cancel := context.WithCancel(context.Background())

	config := &Config{
		HTTPPort: 80,
	}

	return &SMRuntime{
		ctx:           ctx,
		cancel:        cancel,
		runtimeConfig: config,
	}
}

func (r *SMRuntime) Run() error {
	start := time.Now()

	d := time.Since(start).Milliseconds()
	logger.Infof("state machine initialized. Status: Running. Init Elapsed %vms", d)

	err := r.initRuntime()
	if err != nil {
		return err
	}

	return nil
}

func (r *SMRuntime) Shutdown() {
	os.Exit(0)
}

func (r *SMRuntime) initRuntime() error {
	// Start HTTP Server
	err := r.startHTTPServer(r.runtimeConfig.HTTPPort)
	if err != nil {
		logger.Fatalf("failed to start HTTP server: %s", err)
	}
	return nil
}

func (r *SMRuntime) startHTTPServer(port int) error {

	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		return err
	}
	return nil
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
