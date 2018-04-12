package gojwt

import (
  "net/http";
  "encoding/json";
)

// parse any to json
func ParseToJSON(s interface{})(responde []byte){
  response, err := json.MarshalIndent(&s, "", "")
  if err != nil {
    _, ok := err.(*json.UnsupportedTypeError)
    if ok {
      return []byte("Tried to Marshal Invalid Type.")
    } else {
      return []byte("Interface passed does not exist.")
    }
  }
  return response
}

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

// middleware gojwt
func MiddlewareGojwtHeaders(pageHandler, noAuthHandler http.HandlerFunc, o *Gojwt) (http.HandlerFunc) {
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

