package conf

import (
  "net/http";
  "github.com/jenazads/gojweto";
)

// create a JWT and put in the clients cookie
func setToken(username string, w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  gojweto.SetSecretKey("Jnzads-web")
  tokenString, _ := gojweto.CreateHS256Token(username)
  dataJSON := gojweto.Credentials{Token: tokenString, Logged: true}
  gojweto.JsonResponse(dataJSON, w)
}
