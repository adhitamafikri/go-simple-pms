package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenExtras struct {
	Roles []struct {
		id   int
		name string
	} `json:"roles"`
	UserId          int   `json:"userId"`
	TeamIds         []int `json:"teamIds"`
	BusinessUnitIds []int `json:"businessUnitIds"`
}

type JWTClaims struct {
	Iss    string       `json:"iss"`
	Aud    []string     `json:"aud"`
	Extras *TokenExtras `json:"extras"`
}

func Hello() {
	fmt.Println("Hello From JWT lib")
}

func CreateNewJWT(jwtKey string, claims JWTClaims) string {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    claims.Iss,
		"aud":    claims.Aud,
		"iat":    now.Unix(),
		"exp":    now.Add(15 * time.Minute).Unix(),
		"extras": claims.Extras,
	})

	s, err := t.SignedString(jwtKey)
	if err != nil {
		panic("")
	}

	return s
}

// func VerifyJWT(jwtString string) bool {
// 	jwt.Claim
// }
