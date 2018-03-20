package gojweto

import (
  "log";
)

// CertToken is an Struct to encapsulate username and expires as parameter
type CredentialsAuth struct {
  Token   string  `json:"Token"`
  Logged  bool    `json:"Logged"`
}

type CredentialsNoAuth struct {
  Logged  bool    `json:"Logged"`
}

func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
