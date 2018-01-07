package gojweto

import (
  "fmt";
  "time";
  "errors";
  "io/ioutil";
  "github.com/dgrijalva/jwt-go";
)

// JWT schema of the data it will store
type Claims struct {
	jwt.StandardClaims
}

func initRS256Vars(){
  verifyBytes, err := ioutil.ReadFile(pwd+pubKey)
  fatal(err)

  verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  fatal(err)

  signBytes, err := ioutil.ReadFile(pwd+privKey)
  fatal(err)

  signingKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  fatal(err)
}

// Create token with RSA256 algorithm
func CreateRS256Token(username string) (string, error) {
  // Create the Claims
  claims := Claims{
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * GetNumHoursDuration()).Unix(), //time.Unix(c.ExpiresAt, 0),
      Issuer:    username,
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
  tokenString, err := token.SignedString(signingKey)
  fatal(err)
  return tokenString, err
}

// Create token with HSA256 algorithm
func CreateHS256Token(username string) (string, error) {
  // Create the Claims
  claims := Claims{
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * GetNumHoursDuration()).Unix(), //time.Unix(c.ExpiresAt, 0),
      Issuer:    username,
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  tokenString, err := token.SignedString(GetSecretByte())
  fatal(err)
  return tokenString, err
}

// Validate Token RS256 algorithm
func ValidateRS256Token(tokenString string) (bool, string, error) {
  if tokenString == "" {
    return false, "", errors.New("token is empty")
  }

  token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {
    return verifyKey, nil
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

// Validate Token HS256 algorithm
func ValidateHS256Token(tokenString string) (bool, string, error) {
  if tokenString == "" {
    return false, "", errors.New("token is empty")
  }
  
  token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return GetSecretByte(), nil
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
