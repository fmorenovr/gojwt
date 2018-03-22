package gojweto

import (
  "log";
)

// Struct to encapsulate if the username is not logged
type CredentialsAuth struct {
  Token   string  `json:"Token"`
  Logged  bool    `json:"Logged"`
}

// Struct to encapsulate if the username is not logged
type CredentialsNoAuth struct {
  Logged  bool    `json:"Logged"`
}

// func to evaluate Err
func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
