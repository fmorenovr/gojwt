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
  pubKeyPath       string             // pub key path (RSA/ECDSA)
  privKeyPath      string             // priv key path (RSA/ECDSA)
  pubECDSAKey      *ecdsa.PublicKey   // pub ECDSA Key
  privECDSAKey     *ecdsa.PrivateKey  // priv ECDSA Key
  pubRSAKey        *rsa.PublicKey     // pub RSA Key
  privRSAKey       *rsa.PrivateKey    // priv RSA Key
  secretKeyWord    string             // secretKey (HMAC-SHA)
  headerKeyAuth    string             // headerAuth (in http-request)
  numHoursDuration time.Duration      // expiration time (hours)
  method           string             // encrypt algorithm
  lenBytes         string             // type of encrypt algorithm (bytes)
  nameServer       string             // claims info
}

// prepare RSA Key pairs function
func prepareRSAKeys(privRSAPath, pubRSAPath string)(*rsa.PublicKey, *rsa.PrivateKey, error){
  pwd, _ := os.Getwd()

  verifyBytes, err := ioutil.ReadFile(pwd+pubRSAPath)
  if err != nil{
    return &rsa.PublicKey{}, &rsa.PrivateKey{}, GojwtErrInvalidEmptyPublicKey
  }

  verifiedKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  if err != nil{
    return &rsa.PublicKey{}, &rsa.PrivateKey{}, GojwtErrIsNotPubRSAKey
  }

  signBytes, err := ioutil.ReadFile(pwd+privRSAPath)
  if err != nil{
    return &rsa.PublicKey{}, &rsa.PrivateKey{}, GojwtErrInvalidEmptyPrivateKey
  }

  signedKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  if err != nil{
    return &rsa.PublicKey{}, &rsa.PrivateKey{}, GojwtErrIsNotPrivRSAKey
  }
  
  return verifiedKey, signedKey, nil
}

// prepare ECDSA Key pairs function
func prepareECDSAKeys(privECDSAPath, pubECDSAPath string)(*ecdsa.PublicKey, *ecdsa.PrivateKey, error){
  pwd, _ := os.Getwd()

  verifyBytes, err := ioutil.ReadFile(pwd+pubECDSAPath)
  if err != nil{
    return &ecdsa.PublicKey{}, &ecdsa.PrivateKey{}, GojwtErrInvalidEmptyPublicKey
  }

  verifiedKey, err := jwt.ParseECPublicKeyFromPEM(verifyBytes)
  if err != nil{
    return &ecdsa.PublicKey{}, &ecdsa.PrivateKey{}, GojwtErrIsNotPubECDSAKey
  }

  signBytes, err := ioutil.ReadFile(pwd+privECDSAPath)
  if err != nil{
    return &ecdsa.PublicKey{}, &ecdsa.PrivateKey{}, GojwtErrInvalidEmptyPrivateKey
  }

  signedKey, err := jwt.ParseECPrivateKeyFromPEM(signBytes)
  if err != nil{
    return &ecdsa.PublicKey{}, &ecdsa.PrivateKey{}, GojwtErrIsNotPrivECDSAKey
  }
  
  return verifiedKey, signedKey, nil
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

// Create a New GoJwt Instance with HMAC-SHA method
func NewGojwtHMAC_SHA(nameserver, headerkey, secretkey, lenbytes string, hours time.Duration) (*Gojwt, error){
  if secretkey == "" {
    return nil, GojwtErrInvalidEmptySecretKey
  }
  return &Gojwt{
         secretKeyWord: secretkey,
         headerKeyAuth: headerkey,
         numHoursDuration: hours,
         method: "HMAC-SHA",
         lenBytes: lenbytes,
         nameServer: nameserver}, nil
}

// Create a New GoJwt Instance with ECDSA method
func NewGojwtECDSA(nameserver, headerkey, privKeyPath, pubKeyPath, lenbytes string, hours time.Duration) (*Gojwt, error){
  var verifiedECDSAKey *ecdsa.PublicKey
  var signedECDSAKey   *ecdsa.PrivateKey
  if privKeyPath == "" {
    return nil, GojwtErrInvalidEmptyPrivateKey
  } else if pubKeyPath == "" {
    return nil, GojwtErrInvalidEmptyPublicKey
  }
  verifiedECDSAKey, signedECDSAKey, err := prepareECDSAKeys(privKeyPath, pubKeyPath)
  if err != nil{
    return nil, err
  }
  return &Gojwt{
       pubKeyPath: pubKeyPath,
       privKeyPath: privKeyPath,
       pubECDSAKey: verifiedECDSAKey,
       privECDSAKey: signedECDSAKey,
       headerKeyAuth: headerkey,
       numHoursDuration: hours,
       method: "ECDSA",
       lenBytes: lenbytes,
       nameServer: nameserver}, nil
}

// Create a New GoJwt Instance with RSA method
func NewGojwtRSA(nameserver, headerkey, privKeyPath, pubKeyPath, lenbytes string, hours time.Duration) (*Gojwt, error){
  var verifiedRSAKey   *rsa.PublicKey
  var signedRSAKey     *rsa.PrivateKey
  
  if privKeyPath == "" {
    return nil, GojwtErrInvalidEmptyPrivateKey
  } else if pubKeyPath == "" {
    return nil, GojwtErrInvalidEmptyPublicKey
  }
  verifiedRSAKey, signedRSAKey, err := prepareRSAKeys(privKeyPath, pubKeyPath)
  if err != nil{
    return nil, err
  }
  return &Gojwt{
       pubKeyPath: pubKeyPath,
       privKeyPath: privKeyPath,
       pubRSAKey: verifiedRSAKey,
       privRSAKey: signedRSAKey,
       headerKeyAuth: headerkey,
       numHoursDuration: hours,
       method: "RSA",
       lenBytes: lenbytes,
       nameServer: nameserver}, nil
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
// set PATH of Public key
func (o *Gojwt) SetPubKeyPath(path string)(){
  o.pubKeyPath = path
}

// get PATH of Public key
func (o *Gojwt) GetPubKeyPath()(string){
  return o.pubKeyPath
}

// set PATH of Private key
func (o *Gojwt) SetPrivKeyPath(path string)(){
  o.privKeyPath = path
}

// get PATH of Private key
func (o *Gojwt) GetPrivKeyPath()(string){
  return o.privKeyPath
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
