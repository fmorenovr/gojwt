package gojwt_test

import (
  "net/http";
  "fmt";
  "github.com/jenazads/gojwt";
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

// How to use goJwt middleware in gorilla or http/net package
func Example_loginHandler(w http.ResponseWriter, r *http.Request) {
  // gojwt: path of priv, pub, nameServer, secretKey, headerAuth, method, bytesmethod, time exp (hours)
  // This var must be global and could be modify using Sets and Gets methods
  //GojwtObject, _ := gojwt.NewGojwtOptions("", "", "gojwetoServer", "secretKey", "Auth-gojweto", "HMAC-SHA", "512", 24)
  GojwtObject, _ := gojwt.NewGojwtOptions(privECDSAKeyPath, pubECDSAKeyPath, "gojwetoServer", "", "Auth-gojweto", "ECDSA", "384", 24)
  //GojwtObject, _ := gojwt.NewGojwtOptions(privRSAKeyPath, pubRSAKeyPath, "gojwetoServer", "", "Auth-gojweto", "RSA", "256", 24)
  var ob Login
  ob.Username = r.FormValue("username")
  ob.Password = r.FormValue("password")
  objectid, err := VerifyLogin(ob)
  if err != nil {
    gojwt.JsonResponse(err.Error(),w)
	} else {
    if objectid {
	    tokenString, _ := GojwtObject.CreateToken(ob.Username)
	    dataJSON := gojwt.CredentialsAuth{Token: tokenString, Logged: objectid}
      gojwt.JsonResponse(dataJSON, w)
    } else {
      dataJSON := gojwt.CredentialsNoAuth{Logged: objectid}
      gojwt.JsonResponse(dataJSON, w)
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
