package handlers

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  var ob Login
  ob.Username = r.FormValue("username")
  ob.Password = r.FormValue("password")
  objectid, err := VerifyLogin(ob)
  gojweto.SetSecretKey("your-web")
  gojweto.SetHeaderKey("your-web-Auth")
  if err != nil {
    gojweto.JsonResponse(err.Error(),w)
	} else {
    if objectid {
	    tokenString, _ := gojweto.CreateHS256Token(ob.Username)
	    dataJSON := gojweto.CredentialsAuth{Token: tokenString, Logged: objectid}
      gojweto.JsonResponse(dataJSON, w)
    } else {
      dataJSON := gojweto.CredentialsNoAuth{Logged: objectid}
      gojweto.JsonResponse(dataJSON, w)
    }
  }
}

// create a JWT and put in the clients cookie
func setTokenHandler(w http.ResponseWriter, r *http.Request) {
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
