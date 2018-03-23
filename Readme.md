# golang + JWT = goJweto

goJweto (Golang for JSON Web Token) is a Golang implementation for REST service security.  
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/goJweto).

## JWT

JWT (JSON Web Token) is a standard to make secure a connection in a compact URL-safe means of representing claims to be transferred between two parties.  
See more info [here](https://jwt.io).

## goJweto

* First, You should create your RSA key pairs.  
  Create `/tls-ssl/jwtkeys/` directory in your root path of your project:

      cd jwt/keys
      openssl genrsa -out rsakey.pem 2048
      openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub

* Or You should create your ECDSA key pairs.  
  Create `/tls-ssl/jwtkeys/` directory in your root path of your project:

    * First, select a curve list:
    
          openssl ecparam -list_curves

    * Then, select secp256r1 or secp384r1:

          cd jwt/keys
          openssl ecparam -genkey -name secp384r1 | sed -e '1,3d' > ecdsakey.pem
          openssl ec -in ecdsakey.pem -pubout > ecdsakey.pem.pub

* Next, You should download my library:

      go get github.com/jenazads/gojweto/

* Then, you should use for differents Web Frameworks in Go.
        
    * First, Create a gojweto object, specifying privKeypath, pubKeyPath, nameServer, secretKey, headerAuth in request, algorithm, bytes, and expiration time (in hours).
    
            var GojwtObject, _ = gojweto.NewGojwetoOptions("", "", "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "HMAC-SHA", "512", 24)
            var GojwtObject. _ = gojweto.NewGojwetoOptions(privECDSAKeyPath, pubECDSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "ECDSA", "384", 24)
            var GojwtObject. _ = gojweto.NewGojwetoOptions(privRSAKeyPath, pubRSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "RSA", "256", 24)
    
        
    * Then, generate the token string specifyind a nameserver and username:
      
            tokenString, _ := GojwtObject.CreateToken(Username)

    * Using in Go net/http package:
      
      * Add `examples/goJwetoHandler.go` in your controllers directory.
      
      * Then, in your muxServe add:
      
        ```go
          muxHttp.HandleFunc("/setToken", setTokenHandler)
          muxHttp.HandleFunc("/login", LoginHandler)
          muxHttp.HandleFunc("/profile", gojweto.MiddlewareGoJwetoHeaders(WithAuthHandler, NoAuthHandler))
        ```

    * Using in BeeGo:
    
      * Add `examples/goJwetoBeeGoController.go` in your controllers directory.
        
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
