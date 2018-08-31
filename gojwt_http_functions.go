package gojwt

import (
  "net/http";
  "encoding/json";
)

// ToJSON return JSON format of elements
func ToJSON(s interface{}) ([]byte, error) {
  response, err := json.MarshalIndent(&s, "", "")
  if err != nil {
    _, ok := err.(*json.UnsupportedTypeError)
    if ok {
      return nil, GojwtErrTriedToMarshal
    } else {
      return nil, GojwtErrInterfaceNotExist
    }
  }
  return response, nil
}

// FromJSON Convert to JSON format the elements
func FromJSON(data []byte) (map[string]interface{}, error) {
  elements := make(map[string]interface{})
  err := json.Unmarshal(data, &elements)
  return elements, err
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

