package gojweto

import (
  "log";
  "errors";
)

var(
  GojwetoErrInvalidEmptySecretKey   = errors.New("Invalid Empty Secret Key")
  GojwetoErrInvalidEmptyPrivateKey  = errors.New("Invalid Empty Path Private Key")
  GojwetoErrInvalidEmptyPublicKey   = errors.New("Invalid Empty Path Public Key")
  GojwetoErrInvalidEmptyToken       = errors.New("Invalid Empty Token")
  GojwetoErrInvalidAlgorithm        = errors.New("Invalid Algorithm to create token")
  GojwetoErrInvalidRSABytes         = errors.New("Invalid RSA len bytes Algorithm")
  GojwetoErrInvalidECDSABytes       = errors.New("Invalid ECDSA len bytes Algorithm")
  GojwetoErrInvalidHMACHSABytes     = errors.New("Invalid HMAC-SHA len bytes Algorithm")
  GojwetoErrInvalidToken            = errors.New("Invalid Token")
  GojwetoErrBadFormatToken          = errors.New("Invalid Format Token")
  GojwetoErrTokenExpired            = errors.New("Invalid Token is Expired")
  GojwetoErrNotWorkToken            = errors.New("Invalid Token is not working")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
