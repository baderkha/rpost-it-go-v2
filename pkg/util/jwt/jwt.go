package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT CLAIMS
type Claims struct {
	jwt.StandardClaims
}

// HMAC 265 TOKEN
type HS256 struct {
	Issuer string
}

func MakeJWTHS265(secret string, issuer string) *HS256 {
	return &HS256{
		Issuer: issuer,
	}
}

// Generates a JWT token
func (j *HS256) GenerateWebToken(subject string, ttlMinutes int64, secretString string) (token string, validity *time.Time, err error) {
	expirationTime := time.Now().Add(time.Duration(ttlMinutes) * time.Minute)
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    j.Issuer,
			Subject:   subject,
		},
	}

	tokenStringified, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secretString))

	if err != nil {
		fmt.Println(err.Error())
		return "", nil, err
	}
	return tokenStringified, &expirationTime, nil

}

// Validates the Token
func (j *HS256) ValidateWebToken(token string, secretString string) (isValid bool, jwtStdClaims *Claims) {
	fmt.Println(token)
	tkn, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})
	if err != nil {
		return false, nil
	}

	if claims, ok := tkn.Claims.(*Claims); ok && tkn.Valid {
		return true, claims
	}

	return false, nil

}
