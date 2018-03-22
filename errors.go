package gojweto

import (
  "errors"
)
  
var(
  ErrInvalidEmptyToken   = errors.New("Invalid Empty Token")
  ErrInvalidAlgorithm    = errors.New("Invalid Algorithm to create token")
  ErrInvalidRSABytes     = errors.New("Invalid RSA len bytes Algorithm")
  ErrInvalidECDSABytes   = errors.New("Invalid ECDSA len bytes Algorithm")
  ErrInvalidHMACHSABytes = errors.New("Invalid HMAC-SHA len bytes Algorithm")
  ErrInvalidToken        = errors.New("Invalid Token")
  ErrBadFormatToken      = errors.New("Invalid Format Token")
  ErrTokenExpired        = errors.New("Invalid Token is Expired")
  ErrNotWorkToken        = errors.New("Invalid Token is not working")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
