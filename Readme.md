# golang + JWT = goJwt (GoJweto)

goJwt (Golang for JSON Web Token) is a Golang implementation for REST service security.  
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/goJwt).

## JWT

JWT (JSON Web Token) is a standard to make secure a connection in a compact URL-safe means of representing claims to be transferred between two parties.  
See more info [here](https://jwt.io).

## goJwt

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

      go get github.com/jenazads/gojwt/

* Then, you should use for differents implements in Go.
        
    * First, Create a HMAC_SHA gojwt object, specifying nameServer, headerAuth in request, secretKey, bytes, and expiration time (in hours).
    
            var GojwtObject, _ = gojwt.NewGojwtHMAC_SHA("JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "512", 24)
    
    * Or a RSA/ECDA Object, specifying nameServer, headerAuth in request, privKeypath, pubKeyPath, bytes, and expiration time (in hours).
    
            var GojwtObject, _ = gojwt.NewGojwtRSA("JnzadsServer", "Jnzads-rest-JWT", privKeyPath, pubKeyPath, "384", 24)
            var GojwtObject, _ = gojwt.NewGojwtECDSA("JnzadsServer", "Jnzads-rest-JWT", privKeyPath, pubKeyPath, "256", 24)
    
        
    * Then, generate the token string specifyind a nameserver and username:
      
            tokenString, _ := GojwtObject.CreateToken(Username)

    * Using in Go net/http package:
      
      * Add `examples/goJwtHandler.go` in your controllers directory.
      
      * Then, in your muxServe add:
      
        ```go
          muxHttp.HandleFunc("/setToken", setTokenHandler)
          muxHttp.HandleFunc("/login", LoginHandler)
          muxHttp.HandleFunc("/profile", gojwt.MiddlewareGojwtHeaders(WithAuthHandler, NoAuthHandler))
        ```

    * Using in BeeGo:
    
      * Add `examples/goJwtBeeGoController.go` in your controllers directory.
        
      * And, in other controllers, add your new controller instead beegoController.
      
        ```go
            import (
              "encoding/json";
              "restfulapi-beego/models";
              //"github.com/astaxie/beego";
            )

            type AlertController struct {
	            //beego.Controller
	            GoJwtController
            }
        ```
