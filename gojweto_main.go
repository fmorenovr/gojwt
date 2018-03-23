package gojweto

import (
  "fmt";
  //"github.com/jenazads/gojweto";
)

var (
  privECDSAKeyPath = "asdsa"
  pubECDSAKeyPath  = ""
  privRSAKeyPath   = ""
  pubRSAKeyPath    = ""
)

func main() {
  //GojwtObject, err := gojweto.NewGojweto()
  GojwtObject, err := NewGojweto()
  fmt.Println("Example with Default config GoJweto Object: ", GojwtObject, "with error: ", err)
  
  //GojwtObject, err = gojweto.NewGojwetoOptions("", "", "", "", "Jnzads-rest-JWT", "HMAC-SHA", "512", 24)
  GojwtObject, err = NewGojwetoOptions("", "", "", "", "Jnzads-rest-JWT", "HMAC-SHA", "512", 24)
  fmt.Println("Example with empty secret Key and HMAC method: ", GojwtObject, "with error: ", err)
  //GojwtObject, err = gojweto.NewGojwetoOptions(privECDSAKeyPath, pubECDSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "ECDSA", "384", 24)
  GojwtObject, err = NewGojwetoOptions(privECDSAKeyPath, pubECDSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "ECDSA", "384", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)
  //GojwtObject, err = gojweto.NewGojwetoOptions(privRSAKeyPath, pubRSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "RSA", "256", 24)
  GojwtObject, err = NewGojwetoOptions(privRSAKeyPath, pubRSAKeyPath, "JnzadsServer", "jnzads-rest", "Jnzads-rest-JWT", "RSA", "256", 24)
  fmt.Println("Example with empty secret Key and RSA/ECDSA method: ", GojwtObject, "with error: ", err)
}
