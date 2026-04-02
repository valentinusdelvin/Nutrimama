	package utils

	import (
		"os"
		"time"
		"github.com/golang-jwt/jwt/v4"
	)

	func getJWTKey() []byte {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "fallback_secret_override_token"
		}
		return []byte(secret)
	}

	type Claims struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
		jwt.RegisteredClaims
	}

	func GenerateToken(ID int, email string, role string) (string, error) {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			ID  : ID,
			Email: email,
			Role: role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(getJWTKey())
	}

	func ValidateToken(tokenStr string) (*Claims, error) {
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return getJWTKey(), nil
		})
		if err != nil || !token.Valid {
			return nil, err
		}
		return claims, nil
	}
