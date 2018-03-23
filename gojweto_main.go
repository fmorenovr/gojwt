package gojweto_test

import (
  "fmt";
  "github.com/jenazads/gojweto";
)

// main function with hello world in goJweto
func main() {
  var (
    privECDSAKeyPath = "asdsa"
    pubECDSAKeyPath  = ""
    privRSAKeyPath   = ""
    pubRSAKeyPath    = ""
  )
  GojwtObject, err := gojweto.NewGojweto()
  fmt.Println("Example with Default config GoJweto Object: ", GojwtObject, "with error: ", err)
  
  GojwtObject, err = gojweto.NewGojwetoOptions("", "", "", "", "Jnzads-rest-JWT", "HMAC-SHA", "512", 24)
  fmt.Println("Example with empty secret Key and HMAC method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojweto.NewGojwetoOptions(privECDSAKeyPath, pubECDSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "ECDSA", "384", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)

  GojwtObject, err = gojweto.NewGojwetoOptions(privRSAKeyPath, pubRSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "RSA", "256", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)
}
