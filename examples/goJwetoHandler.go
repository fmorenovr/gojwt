package conf

import (
  "fmt";
  "net/http";
  "github.com/jenazads/gojweto";
)

// create a JWT and put in the clients cookie
func setToken(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  gojweto.SetSecretKey("beagons-web")
  tokenString, _ := gojweto.CreateHS256Token("localhost:9000")
  dataJSON := gojweto.Credentials{Token: tokenString, Logged: true}
  gojweto.JsonResponse(dataJSON, w)
}

func protectedProfile(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world protected")
}
