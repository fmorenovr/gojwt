package gojweto

import (
  "os";
  "log";
  "crypto/rsa";
)

// CertToken is an Struct to encapsulate username and expires as parameter
type CertToken struct {
  Username string
  Expires int64
}

type Credentials struct {
  Logged bool
  Token  string
}

const (
  privKey    = "/tls-ssl/jwtkeys/rsakey.pem"
  pubKey     = "/tls-ssl/jwtkeys/rsakey.pem.pub"
  secretKey  = "Jnzads"
)

var (
  verifyKey   *rsa.PublicKey
  signingKey  *rsa.PrivateKey
  secretByte  = []byte(secretKey)
  pwd, _      = os.Getwd()
)

func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
