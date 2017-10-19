# goJweto

goJweto (Golang JSON Web Token) is a Golang implementation for REST service security.

JWT (JSON Web Token) is a Golang implementation.  
* First, You should create your RSA key pairs.  
  Create `/tls-ssl/jwtkeys/` directory in your root path of your project:

      cd jwt/keys
      openssl genrsa -out rsakey.pem 2048
      openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub

* Next, Install it:

      go get github.com/dgrijalva/jwt-go/

* Once installed, You should download my library:

      go get github.com/jenazads/gojweto/

* Then, you should use for differents Web Frameworks in Go.

    * Using in Go net/http package:
      
      * Add `examples/goJwetoHandler.go` in your controllers directory.
      
      * Then, in your tokenHandler add:
      
        ```go
          muxHttp.HandleFunc("/settoken", setToken)
          muxHttp.HandleFunc("/profile", gojweto.MiddlewareGoJwetoHeaders(protectedProfile, NotFoundHandler,"Jnzads-web-JWT"))
        ```

    * Using in BeeGo:
    
      * Add `examples/goJwetoBeeGoController.go` in your controllers directory.
        
      * Then, in your loginController add:
      
            tokenString, _ = gojweto.CreateRS256Token(Username)
        
      * And, in other controllers, add your new controller instead beegoController.
      
        ```go
            import (
              "encoding/json";
              "restfulapi-beego/models";
              //"github.com/astaxie/beego";
            )

            type AlertController struct {
	            //beego.Controller
	            GoJwetoController
            }
        ```
