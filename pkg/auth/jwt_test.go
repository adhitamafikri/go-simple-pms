// go:build unit

package auth

import (
	"fmt"
	"testing"
)

func TestWhenCreateNewJWT_ThenSuccessProduceJWT(t *testing.T) {
	t.Run("Should create proper JWT with these given parameters", func(t *testing.T) {
		claims := &JWTClaims{
			Iss: "adhitamafikri.dev@gmail.com",
			Aud: []string{"Admin"},
			Extras: &TokenExtras{
				UserId: 64,
				Roles: &[]TokenRole{
					{ID: 1, Name: "Admin"},
					{ID: 2, Name: "Guest"},
				},
				TeamIds:         []int{1, 2},
				BusinessUnitIds: []int{1, 2, 3},
			},
		}

		jwtAuth := NewJWTAuth()
		key := []byte("SomeRandomJWTKey")
		val, err := jwtAuth.CreateNewJWT(key, claims)

		if err != nil {
			t.Fatal("Failed creating JWT", err)
		}

		fmt.Println("[jwtAuth.CreateNewJWT] success:\n", val)
	})
}

func TestWhenParseAndVerifyJWT_ThenReturnSuccess(t *testing.T) {
	t.Run("Should create proper JWT with these given parameters", func(t *testing.T) {
		jwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiQWRtaW4iXSwiZXhwIjoxNzc2NTkzMDM1LCJleHRyYXMiOnsicm9sZXMiOlt7ImlkIjoxLCJuYW1lIjoiQWRtaW4ifSx7ImlkIjoyLCJuYW1lIjoiR3Vlc3QifV0sInVzZXJfaWQiOjY0LCJ0ZWFtX2lkcyI6WzEsMl0sImJ1c2luZXNzX3VuaXRfaWRzIjpbMSwyLDNdfSwiaWF0IjoxNzc2NTkyMTM1LCJpc3MiOiJhZGhpdGFtYWZpa3JpLmRldkBnbWFpbC5jb20ifQ.UZ2wQYSGjUPKWIgORFB_OQWfaM6x9HGFOJDb2A6iS3Q"

		jwtAuth := NewJWTAuth()
		key := []byte("SomeRandomJWTKey")
		val, err := jwtAuth.ParseAndVerifyJWT(jwtString, key)

		if err != nil {
			t.Fatal("Failed verifying JWT", err)
		}

		fmt.Println("[jwtAuth.ParseAndVerifyJWT] success", val)
	})
}
