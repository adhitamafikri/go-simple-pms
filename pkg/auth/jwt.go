package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TokenExtras struct {
	Roles           *[]TokenRole `json:"roles"`
	UserId          int          `json:"user_id"`
	TeamIds         []int        `json:"team_ids"`
	BusinessUnitIds []int        `json:"business_unit_ids"`
}

type JWTClaims struct {
	Iss    string       `json:"iss"`
	Aud    []string     `json:"aud"`
	Extras *TokenExtras `json:"extras"`
}

type JWTAuth interface {
	Hello()
	CreateNewJWT(jwtKey string, claims JWTClaims) (string, error)
}

type jwtAuth struct {
	signed bool
}

func NewJWTAuth() *jwtAuth {
	return &jwtAuth{signed: true}
}

func (j *jwtAuth) CreateNewJWT(jwtKey []byte, claims *JWTClaims) (string, error) {
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
		return "", err
	}

	return s, nil
}

// func VerifyJWT(jwtString string) bool {
// 	jwt.Claim
// }
