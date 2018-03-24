package gojwt

import (
  "log";
  "errors";
)

var(
  GojwtErrInvalidEmptySecretKey   = errors.New("Invalid Empty Secret Key")
  GojwtErrInvalidEmptyPrivateKey  = errors.New("Invalid Empty Path Private Key")
  GojwtErrInvalidEmptyPublicKey   = errors.New("Invalid Empty Path Public Key")
  GojwtErrInvalidEmptyToken       = errors.New("Invalid Empty Token")
  GojwtErrInvalidAlgorithm        = errors.New("Invalid Algorithm to create token")
  GojwtErrInvalidRSABytes         = errors.New("Invalid RSA len bytes Algorithm")
  GojwtErrInvalidECDSABytes       = errors.New("Invalid ECDSA len bytes Algorithm")
  GojwtErrInvalidHMACHSABytes     = errors.New("Invalid HMAC-SHA len bytes Algorithm")
  GojwtErrInvalidToken            = errors.New("Invalid Token")
  GojwtErrBadFormatToken          = errors.New("Invalid Format Token")
  GojwtErrTokenExpired            = errors.New("Invalid Token is Expired")
  GojwtErrNotWorkToken            = errors.New("Invalid Token is not working")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
