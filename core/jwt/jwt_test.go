package jwt

import (
	"os"
	"testing"
)

func LoadEnv() {
	os.Setenv("JWT_SECRET", "kitchen")
}

func TestJWT(t *testing.T) {
	LoadEnv()
	t.Run("Testing JWT token generation", func(t *testing.T) {
		_, err := generateToken("test")
		if err != nil {
			t.Fatalf("Failed to generate token: %s", err.Error())
		}
	})

	t.Run("Testing JWT token validation", func(t *testing.T) {
		tokenStr, err := generateToken("test")
		if err != nil {
			t.Fatalf("Failed to generate token: %s", err.Error())
		}
		succeeed, err := validate(tokenStr)
		if err != nil || !succeeed {
			t.Fatalf("Failed to validate token: %v", err.Error())
		}
	})
}
