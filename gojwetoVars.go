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

