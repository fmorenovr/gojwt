package gojweto

import (
  "net/http";
  "encoding/json";
)

// Write in JSON Format
func JsonResponse(response interface{}, w http.ResponseWriter) {
  json, err := json.Marshal(response)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write(json)
}

// middleware gojweto
func MiddlewareGoJwetoHeaders(pageHandler, noAuthHandler http.HandlerFunc, o Gojweto) (http.HandlerFunc) {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get(o.GetHeaderKey())
    valid, _, _ := o.ValidateToken(tokenString)
    if !valid {
      noAuthHandler(w, r)
      return
    } else {
      pageHandler(w, r)
      return
    }
  })
}

