package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

/*
change-able claims:

	type StandardClaims struct {
	    Audience  string `json:"aud,omitempty"`
	    ExpiresAt int64  `json:"exp,omitempty"`
	    Id        string `json:"jti,omitempty"`
	    IssuedAt  int64  `json:"iat,omitempty"`
	    Issuer    string `json:"iss,omitempty"`
	    NotBefore int64  `json:"nbf,omitempty"`
	    Subject   string `json:"sub,omitempty"`
	}
*/
var (
	err       error
	masterKey string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	masterKey = os.Getenv("JWT_SECRET_KEY")
}

func GenerateJWT(name string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    name,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token_str, err := claims.SignedString([]byte(masterKey))
	if err != nil {
		return "", err
	}
	return token_str, nil
}

func VerifyJWT(authString string) (bool, error) {

	//removes the "Bearer " prefix
	tokenString := authString[7:]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return masterKey, nil
	})

	if err != nil {
		return false, err
	}

	//could directly return here, but i just want to be explicit.
	if !token.Valid {
		return false, nil
	} else {
		return true, nil
	}
}
