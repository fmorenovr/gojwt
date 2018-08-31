package gojwt

import (
  "log";
  "errors";
)

var(
  GojwtErrInvalidEmptySecretKey   = errors.New("Invalid Empty Secret Key.\n")
  GojwtErrInvalidEmptyPrivateKey  = errors.New("Invalid Empty Path Private Key.\n")
  GojwtErrInvalidEmptyPublicKey   = errors.New("Invalid Empty Path Public Key.\n")
  GojwtErrInvalidEmptyToken       = errors.New("Invalid Empty Token")
  GojwtErrInvalidAlgorithm        = errors.New("Invalid Algorithm to create token.\n")
  GojwtErrInvalidRSABytes         = errors.New("Invalid RSA len bytes Algorithm.\n")
  GojwtErrInvalidECDSABytes       = errors.New("Invalid ECDSA len bytes Algorithm.\n")
  GojwtErrInvalidHMACHSABytes     = errors.New("Invalid HMAC-SHA len bytes Algorithm.\n")
  GojwtErrInvalidToken            = errors.New("Invalid Token")
  GojwtErrBadFormatToken          = errors.New("Invalid Format Token")
  GojwtErrTokenExpired            = errors.New("Invalid Token is Expired.\n")
  GojwtErrNotWorkToken            = errors.New("Invalid Token is not working.\n")
  GojwtErrIsNotPubECDSAKey        = errors.New("Is not an ECDSA Public Key.\n")
  GojwtErrIsNotPrivECDSAKey       = errors.New("Is not an ECDSA Private Key.\n")
  GojwtErrIsNotPubRSAKey          = errors.New("Is not a RSA Public Key.\n")
  GojwtErrIsNotPrivRSAKey         = errors.New("Is not a RSA Private Key.\n")
  GojwtErrTriedToMarshal          = errors.New("Tried to Marshal Invalid Type.\n")
  GojwtErrInterfaceNotExist       = errors.New("Interface passed does not exist.\n")
)

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
