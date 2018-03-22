package gojweto

import (
  "fmt";
  "time";
  "errors";
  "github.com/dgrijalva/jwt-go";
)

// JWT schema of the data it will store
type Claims struct {
  NameServer   string `json:"nameServer,omitempty"`
  jwt.StandardClaims
}

func (o *Gojweto) CreateToken(username string) (tokenString string, err error) {
  method := o.GetEncryptMethod()
  lenByte := o.GetEncryptLenBytes()
  if method == "RSA" {
    tokenString, err = o.createRSAToken(lenByte, username)
  } else if method == "ECDSA" {
    tokenString, err = o.createECDSAToken(lenByte, username)
  } else if method == "HMAC-SHA" {
    tokenString, err = o.createHMACSHAToken(lenByte, username)
  } else {
    return "", ErrInvalidAlgorithm
  }
  return tokenString, err
}

func (o *Gojweto) ValidateToken(tokenString string) (isValid bool, username string, err error) {
  method := o.GetEncryptMethod()
  if method == "RSA" || method == "ECDSA" {
    isValid, username, err = o.validateECD_RSAToken(tokenString)
  } else if method == "HMAC-SHA" {
    isValid, username, err = o.validateHMACSHAToken(tokenString)
  } else {
    return false, "", ErrInvalidEmptyToken
  }
  return isValid, username, err
}

// Create token with RSA algorithm
func (o *Gojweto) createRSAToken(lenBytes, username string) (string, error) {
  var token *jwt.Token
  // Create the Claims
  claims := Claims{
    o.GetNameServer(),
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * o.GetNumHoursDuration()).Unix(), //time.Unix(c.ExpiresAt, 0),
      Issuer:    username,
      IssuedAt: time.Now().Unix(),
    },
  }
  if lenBytes == "256" {
    token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
  } else if lenBytes == "384"{
    token = jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
  } else if lenBytes == "512"{
    token = jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
  } else {
    return "", ErrInvalidRSABytes
  }
  tokenString, err := token.SignedString(o.GetRSAPrivKey())
  fatal(err)
  return tokenString, err
}

// Create token with ECDSA algorithm
func (o *Gojweto) createECDSAToken(lenBytes, username string) (string, error) {
  var token *jwt.Token
  // Create the Claims
  claims := Claims{
    o.GetNameServer(),
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * o.GetNumHoursDuration()).Unix(), //time.Unix(c.ExpiresAt, 0),
      Issuer:    username,
      IssuedAt: time.Now().Unix(),
    },
  }
  if lenBytes == "256" {
    token = jwt.NewWithClaims(jwt.SigningMethodES256, claims)
  } else if lenBytes == "384"{
    token = jwt.NewWithClaims(jwt.SigningMethodES384, claims)
  //} else if lenBytes == "512"{
  //  token = jwt.NewWithClaims(jwt.SigningMethodES512, claims)
  } else {
    return "", ErrInvalidECDSABytes
  }
  tokenString, err := token.SignedString(o.GetECDSAPrivKey())
  fatal(err)
  return tokenString, err
}

// Create token with HMAC-SHA algorithm
func (o *Gojweto) createHMACSHAToken(lenBytes, username string) (string, error) {
  var token *jwt.Token
  // Create the Claims
  claims := Claims{
    o.GetNameServer(),
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * o.GetNumHoursDuration()).Unix(), //time.Unix(c.ExpiresAt, 0),
      Issuer:    username,
      IssuedAt: time.Now().Unix(),
    },
  }
  if lenBytes == "256"{
    token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  } else if lenBytes == "384"{
    token = jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
  } else if lenBytes == "512"{
    token = jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
  } else {
    return "", ErrInvalidHMACHSABytes
  }
  tokenString, err := token.SignedString(o.GetSecretByte())
  fatal(err)
  return tokenString, err
}

// Validate Token RSA/ECDSA algorithm
func (o *Gojweto) validateECD_RSAToken(tokenString string) (bool, string, error) {
  method := o.GetEncryptMethod()
  if tokenString == "" {
    return false, "", ErrInvalidEmptyToken
  }

  token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
    if method == "RSA"{
      return o.GetRSAPubKey(), nil
    } else if method == "ECDSA" {
      return o.GetECDSAPubKey(), nil
    } else {
      return nil, errors.New("token invalid")
    }
  })

  if token == nil {
    return false, "", errors.New("not work")
  }

  if token.Valid {
    //"You look nice today"
    claims, _ := token.Claims.(jwt.MapClaims)
    //var user string = claims["username"].(string)
    iss := claims["iss"].(string)
    return true, iss, nil
  } else if ve, ok := err.(*jwt.ValidationError); ok {
    if ve.Errors&jwt.ValidationErrorMalformed != 0 {
      return false, "", errors.New("That's not even a token")
    } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
      // Token is either expired or not active yet
      return false, "", errors.New("Timing is everything")
    } else {
      //"Couldn't handle this token:"
      return false, "", err
    }
  } else {
    //"Couldn't handle this token:"
    return false, "", err
  }
}

// Validate Token HMAC-SHA algorithm
func (o *Gojweto) validateHMACSHAToken(tokenString string) (bool, string, error) {
  if tokenString == "" {
    return false, "", ErrInvalidEmptyToken
  }
  
  token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return o.GetSecretByte(), nil
	})
	
	if token == nil {
    return false, "", errors.New("not work")
  }
  
  if token.Valid {
    //"You look nice today"
    claims, _ := token.Claims.(*Claims)
    //var user string = claims["username"].(string)
    iss := claims.Issuer
    return true, iss, err
  } else if ve, ok := err.(*jwt.ValidationError); ok {
    if ve.Errors&jwt.ValidationErrorMalformed != 0 {
      return false, "", errors.New("That's not even a token")
    } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
      // Token is either expired or not active yet
      return false, "", errors.New("Timing is everything")
    } else {
      //"Couldn't handle this token:"
      return false, "", err
    }
  } else {
    //"Couldn't handle this token:"
    return false, "", err
  }
}
