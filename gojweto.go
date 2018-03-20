package gojweto

import (
  "os";
  "time";
  "io/ioutil";
  "crypto/rsa";
  "crypto/ecdsa";
  "github.com/dgrijalva/jwt-go";
)

type Gojweto struct {
  pubECDSAPath     string             // pub ECDSA path
  privECDSAPath    string             // priv ECDSA path
  pubECDSAKey      *ecdsa.PublicKey   // pub ECDSA Key
  privECDSAKey     *ecdsa.PrivateKey  // priv ECDSA Key
  pubRSAPath       string             // pub RSA Path
  privRSAPath      string             // priv RSA Path  
  pubRSAKey        *rsa.PublicKey     // pub RSA Key
  privRSAKey       *rsa.PrivateKey    // priv RSA Key
  secretKeyWord    string             // secretKey to encrypt
  headerKeyAuth    string             // headerAuth (in request)
  numHoursDuration time.Duration      // expiration time (hours)
  method           string             // encrypt algorithm
  lenBytes         string             // type of encrypt algorithm (bytes)
  nameServer       string             // claims info
}

func prepareRSAKeys(privRSAPath, pubRSAPath string)(*rsa.PublicKey, *rsa.PrivateKey){
  pwd, _ := os.Getwd()

  verifyBytes, err := ioutil.ReadFile(pwd+pubRSAPath)
  fatal(err)

  verifiedKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  fatal(err)

  signBytes, err := ioutil.ReadFile(pwd+privRSAPath)
  fatal(err)

  signedKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  fatal(err)
  
  return verifiedKey, signedKey
}

func prepareECDSAKeys(privECDSAPath, pubECDSAPath string)(*ecdsa.PublicKey, *ecdsa.PrivateKey){
  pwd, _ := os.Getwd()

  verifyBytes, err := ioutil.ReadFile(pwd+pubECDSAPath)
  fatal(err)

  verifiedKey, err := jwt.ParseECPublicKeyFromPEM(verifyBytes)
  fatal(err)

  signBytes, err := ioutil.ReadFile(pwd+privECDSAPath)
  fatal(err)

  signedKey, err := jwt.ParseECPrivateKeyFromPEM(signBytes)
  fatal(err)
  
  return verifiedKey, signedKey
}

func NewGojweto() (Gojweto){

  return Gojweto{
         secretKeyWord: "Jnzads",
         headerKeyAuth: "Jnzads-JWT",
         numHoursDuration: 1,
         method: "HMAC-SHA",
         lenBytes: "256",
         nameServer: "JnzadsServer"}
}

func NewGojwetoOptions(privKeyPath, pubKeyPath, nameserver, secretkey, headerkey, method, bytes string, hours time.Duration) (Gojweto){
  var verifiedRSAKey   *rsa.PublicKey
  var signedRSAKey     *rsa.PrivateKey
  var verifiedECDSAKey *ecdsa.PublicKey
  var signedECDSAKey   *ecdsa.PrivateKey
  if method == "RSA" {
    verifiedRSAKey, signedRSAKey = prepareRSAKeys(privKeyPath, pubKeyPath)
    return Gojweto{
         pubRSAPath: pubKeyPath,
         privRSAPath: privKeyPath,
         pubRSAKey: verifiedRSAKey,
         privRSAKey: signedRSAKey,
         secretKeyWord: secretkey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}
  } else if method == "ECDSA" {
    verifiedECDSAKey, signedECDSAKey = prepareECDSAKeys(privKeyPath, pubKeyPath)
    return Gojweto{
         pubECDSAPath: pubKeyPath,
         privECDSAPath: privKeyPath,
         pubECDSAKey: verifiedECDSAKey,
         privECDSAKey: signedECDSAKey,
         secretKeyWord: secretkey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}
  } else if method == "HMAC-SHA" {
    return Gojweto{
         secretKeyWord: secretkey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}
  } else {
    return NewGojweto()
  }
}

// hours of duration
func (o *Gojweto) SetNumHoursDuration(hours time.Duration){
  o.numHoursDuration = hours
}

func (o *Gojweto) GetNumHoursDuration()(time.Duration){
  return o.numHoursDuration
}

// Header authorization
func (o *Gojweto) SetHeaderKey(name string){
  o.headerKeyAuth = name
}

func (o *Gojweto) GetHeaderKey()(string){
  return o.headerKeyAuth
}

// Name Server
func (o *Gojweto) SetNameServer(name string){
  o.nameServer = name
}

func (o *Gojweto) GetNameServer()(string){
  return o.nameServer
}

// just for method HMACSHA
func (o *Gojweto) SetSecretKey(name string){
  o.secretKeyWord = name
}

func (o *Gojweto) GetSecretKey()(string){
  return o.secretKeyWord
}

func (o *Gojweto) GetSecretByte()([]byte){
  return []byte(o.secretKeyWord)
}

// just for method RSA
func (o *Gojweto) GetRSAPrivKey()(*rsa.PrivateKey){
  return o.privRSAKey
}

func (o *Gojweto) GetRSAPubKey()(*rsa.PublicKey){
  return o.pubRSAKey
}

// just for method ECDSA
func (o *Gojweto) GetECDSAPrivKey()(*ecdsa.PrivateKey){
  return o.privECDSAKey
}

func (o *Gojweto) GetECDSAPubKey()(*ecdsa.PublicKey){
  return o.pubECDSAKey
}

// path of keys RSA/ECDSA
func (o *Gojweto) SetPubRSAPath(path string)(){
  o.pubRSAPath = path
}

func (o *Gojweto) GetPubRSAPath()(string){
  return o.pubRSAPath
}

func (o *Gojweto) SetPrivRSAPath(path string)(){
  o.privRSAPath = path
}

func (o *Gojweto) GetPrivRSAPath()(string){
  return o.privRSAPath
}

func (o *Gojweto) SetPubECDSAPath(path string)(){
  o.pubECDSAPath = path
}

func (o *Gojweto) GetPubECDSAPath()(string){
  return o.pubECDSAPath
}

func (o *Gojweto) SetPrivECDSAPath(path string)(){
  o.privECDSAPath = path
}

func (o *Gojweto) GetPrivECDSAPath()(string){
  return o.privECDSAPath
}

// change encrypt method
func (o *Gojweto) SetEncryptMethod(method string)(){
  o.method = method
}

func (o *Gojweto) GetEncryptMethod()(string){
  return o.method
}

// change encrypt method
func (o *Gojweto) SetEncryptLenBytes(lenBytes string)(){
  o.lenBytes = lenBytes
}

func (o *Gojweto) GetEncryptLenBytes()(string){
  return o.lenBytes
}
