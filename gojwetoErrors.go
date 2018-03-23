package gojweto

import (
  "log";
  "errors";
)

var(
  ErrInvalidEmptySecretKey   = errors.New("Invalid Empty Secret Key")
  ErrInvalidEmptyPrivateKey  = errors.New("Invalid Empty Path Private Key")
  ErrInvalidEmptyPublicKey   = errors.New("Invalid Empty Path Public Key")
  ErrInvalidEmptyToken       = errors.New("Invalid Empty Token")
  ErrInvalidAlgorithm        = errors.New("Invalid Algorithm to create token")
  ErrInvalidRSABytes         = errors.New("Invalid RSA len bytes Algorithm")
  ErrInvalidECDSABytes       = errors.New("Invalid ECDSA len bytes Algorithm")
  ErrInvalidHMACHSABytes     = errors.New("Invalid HMAC-SHA len bytes Algorithm")
  ErrInvalidToken            = errors.New("Invalid Token")
  ErrBadFormatToken          = errors.New("Invalid Format Token")
  ErrTokenExpired            = errors.New("Invalid Token is Expired")
  ErrNotWorkToken            = errors.New("Invalid Token is not working")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
