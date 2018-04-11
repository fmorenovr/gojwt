package gojwt_test

import (
  "fmt";
  "github.com/jenazads/gojwt";
)

// main function with hello world in goJwt
func Example_createGoJwtObjects() {
  var (
    privECDSAKeyPath = "asdsa"
    pubECDSAKeyPath  = ""
    privRSAKeyPath   = ""
    pubRSAKeyPath    = ""
  )
  GojwtObject, err := gojwt.NewGojwt()
  fmt.Println("Example with Default config GoJwt Object: ", GojwtObject, "with error: ", err)
  
  GojwtObject, err = gojwt.NewGojwtHMAC_SHA("JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "512", 24)
  fmt.Println("Example with empty secret Key and HMAC method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojwt.NewGojwtRSA("JnzadsServer", "Jnzads-rest-JWT", privKeyPath, pubKeyPath, "384", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojwt.NewGojwtECDSA("JnzadsServer", "Jnzads-rest-JWT", privKeyPath, pubKeyPath, "256", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)
}

