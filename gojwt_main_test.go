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
  
  GojwtObject, err = gojwt.NewGojwtOptions("", "", "", "", "Jnzads-rest-JWT", "HMAC-SHA", "512", 24)
  fmt.Println("Example with empty secret Key and HMAC method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojwt.NewGojwtOptions(privECDSAKeyPath, pubECDSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "ECDSA", "384", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojwt.NewGojwtOptions(privRSAKeyPath, pubRSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "RSA", "256", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)
}

