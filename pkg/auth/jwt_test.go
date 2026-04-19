//go:build unit

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
