package gojweto_test

import (
  "net/http";
  "fmt";
  "github.com/jenazads/gojweto";
)

type Login struct{
  Username string `json:"username"`
  Password string `json:"password"`
}

func VerifyLogin(ob Login)(bool,error){
  if ob.Username != "" && ob.Password != "" {
    return true, nil
  } else{
    return false, nil
  }
}

// gojweto: path of priv, pub, nameServer, secretKey, headerAuth, method, bytesmethod, time exp (hours)
//var GojwtObject = gojweto.NewGojwetoOptions("", "", "gojwetoServer", "secretKey", "Auth-gojweto", "HMAC-SHA", "512", 24)
var GojwtObject = gojweto.NewGojwetoOptions(privECDSAKeyPath, pubECDSAKeyPath, "gojwetoServer", "", "Auth-gojweto", "ECDSA", "384", 24)
//var GojwtObject = gojweto.NewGojwetoOptions(privRSAKeyPath, pubRSAKeyPath, "gojwetoServer", "", "Auth-gojweto", "RSA", "256", 24)

// How to use goJweto middleware in gorilla or http/net package
func Example_loginHandler(w http.ResponseWriter, r *http.Request) {
  var ob Login
  ob.Username = r.FormValue("username")
  ob.Password = r.FormValue("password")
  objectid, err := VerifyLogin(ob)
  if err != nil {
    gojweto.JsonResponse(err.Error(),w)
	} else {
    if objectid {
	    tokenString, _ := GojwtObject.CreateToken(ob.Username)
	    dataJSON := gojweto.CredentialsAuth{Token: tokenString, Logged: objectid}
      gojweto.JsonResponse(dataJSON, w)
    } else {
      dataJSON := gojweto.CredentialsNoAuth{Logged: objectid}
      gojweto.JsonResponse(dataJSON, w)
    }
  }
}

// Web to Create a JWT and put in the clients cookie
func Example_setTokenHandler(w http.ResponseWriter, r *http.Request) {
  const indexPage = `<h1>Login</h1>
    <form method="post" action="/login">
      <label for="name">User name</label>
      <input type="text" id="username" name="username">
      <label for="password">Password</label>
      <input type="password" id="password" name="password">
      <button type="submit">Login</button>
    </form>`
  fmt.Fprintf(w, indexPage)
}
