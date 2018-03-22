package gojweto

import "errors"

var(
  ErrInvalidEmptyToken   = errors.New("Invalid Empty Token")
  ErrInvalidAlgorithm    = errors.New("Invalid Algorithm to create token")
  ErrInvalidRSABytes     = errors.New("Invalid RSA len bytes Algorithm")
  ErrInvalidECDSABytes   = errors.New("Invalid ECDSA len bytes Algorithm")
  ErrInvalidHMACHSABytes = errors.New("Invalid HMAC-SHA len bytes Algorithm")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
