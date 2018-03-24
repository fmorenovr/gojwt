package gojwt

import (
  "os";
  "time";
  "io/ioutil";
  "crypto/rsa";
  "crypto/ecdsa";
  "github.com/dgrijalva/jwt-go";
)

// goJwt Struct
type Gojwt struct {
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

// prepare RSA Key pairs function
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

// prepare ECDSA Key pairs function
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

// Create a New GoJwt Instance with HMAC-SHA encrypt method by default
func NewGojwt() (*Gojwt, error){

  return &Gojwt{
         secretKeyWord: "Jnzads",
         headerKeyAuth: "Jnzads-JWT",
         numHoursDuration: 1,
         method: "HMAC-SHA",
         lenBytes: "256",
         nameServer: "JnzadsServer"}, nil
}

// Create a New GoJwt Instance with an encrypt method with parameters as you wish
func NewGojwtOptions(privKeyPath, pubKeyPath, nameserver, secretkey, headerkey, method, bytes string, hours time.Duration) (*Gojwt, error){
  var verifiedRSAKey   *rsa.PublicKey
  var signedRSAKey     *rsa.PrivateKey
  var verifiedECDSAKey *ecdsa.PublicKey
  var signedECDSAKey   *ecdsa.PrivateKey
  
  if method == "RSA" {
    if privKeyPath == "" {
      return nil, GojwtErrInvalidEmptyPrivateKey
    } else if pubKeyPath == "" {
      return nil, GojwtErrInvalidEmptyPublicKey
    }
    verifiedRSAKey, signedRSAKey = prepareRSAKeys(privKeyPath, pubKeyPath)
    return &Gojwt{
         pubRSAPath: pubKeyPath,
         privRSAPath: privKeyPath,
         pubRSAKey: verifiedRSAKey,
         privRSAKey: signedRSAKey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}, nil
  } else if method == "ECDSA" {
    if privKeyPath == "" {
      return nil, GojwtErrInvalidEmptyPrivateKey
    } else if pubKeyPath == "" {
      return nil, GojwtErrInvalidEmptyPublicKey
    }
    verifiedECDSAKey, signedECDSAKey = prepareECDSAKeys(privKeyPath, pubKeyPath)
    return &Gojwt{
         pubECDSAPath: pubKeyPath,
         privECDSAPath: privKeyPath,
         pubECDSAKey: verifiedECDSAKey,
         privECDSAKey: signedECDSAKey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}, nil
  } else if method == "HMAC-SHA" {
    if secretkey == "" {
      return nil, GojwtErrInvalidEmptySecretKey
    }
    return &Gojwt{
         secretKeyWord: secretkey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: method,
         lenBytes: bytes,
         nameServer: nameserver}, nil
  } else {
    return NewGojwt()
  }
}

// set hours of token duration
func (o *Gojwt) SetNumHoursDuration(hours time.Duration){
  o.numHoursDuration = hours
}

// get hours of token duration
func (o *Gojwt) GetNumHoursDuration()(time.Duration){
  return o.numHoursDuration
}

// set Header authorization
func (o *Gojwt) SetHeaderKey(name string){
  o.headerKeyAuth = name
}

// get Header authorization
func (o *Gojwt) GetHeaderKey()(string){
  return o.headerKeyAuth
}

// set Name Server
func (o *Gojwt) SetNameServer(name string){
  o.nameServer = name
}

// get Name Server
func (o *Gojwt) GetNameServer()(string){
  return o.nameServer
}

// just for method HMACSHA
// set secret key word to encrypt using hmac-sha
func (o *Gojwt) SetSecretKey(name string){
  o.secretKeyWord = name
}

// get secret key word to encrypt using hmac-sha
func (o *Gojwt) GetSecretKey()(string){
  return o.secretKeyWord
}

// get secret key word convert in bytes
func (o *Gojwt) GetSecretByte()([]byte){
  return []byte(o.secretKeyWord)
}

// just for method RSA
// get RSA private key
func (o *Gojwt) GetRSAPrivKey()(*rsa.PrivateKey){
  return o.privRSAKey
}

// get RSA public key
func (o *Gojwt) GetRSAPubKey()(*rsa.PublicKey){
  return o.pubRSAKey
}

// just for method ECDSA
// get ECDSA private key
func (o *Gojwt) GetECDSAPrivKey()(*ecdsa.PrivateKey){
  return o.privECDSAKey
}

// get ECDSA public key
func (o *Gojwt) GetECDSAPubKey()(*ecdsa.PublicKey){
  return o.pubECDSAKey
}

// path of keys RSA/ECDSA
// set PATH of RSA Public key
func (o *Gojwt) SetPubRSAPath(path string)(){
  o.pubRSAPath = path
}

// get PATH of RSA Public key
func (o *Gojwt) GetPubRSAPath()(string){
  return o.pubRSAPath
}

// set PATH of RSA Private key
func (o *Gojwt) SetPrivRSAPath(path string)(){
  o.privRSAPath = path
}

// get PATH of RSA Private key
func (o *Gojwt) GetPrivRSAPath()(string){
  return o.privRSAPath
}

// set PATH of ECDSA Public key
func (o *Gojwt) SetPubECDSAPath(path string)(){
  o.pubECDSAPath = path
}

// get PATH of ECDSA Public key
func (o *Gojwt) GetPubECDSAPath()(string){
  return o.pubECDSAPath
}

// set PATH of ECDSA Private key
func (o *Gojwt) SetPrivECDSAPath(path string)(){
  o.privECDSAPath = path
}

// get PATH of ECDSA Private key
func (o *Gojwt) GetPrivECDSAPath()(string){
  return o.privECDSAPath
}

// change encrypt method
func (o *Gojwt) SetEncryptMethod(method string)(){
  o.method = method
}

// get current encrypt method
func (o *Gojwt) GetEncryptMethod()(string){
  return o.method
}

// set bytes of encrypt method
func (o *Gojwt) SetEncryptLenBytes(lenBytes string)(){
  o.lenBytes = lenBytes
}

// get bytes of encrypt method
func (o *Gojwt) GetEncryptLenBytes()(string){
  return o.lenBytes
}
