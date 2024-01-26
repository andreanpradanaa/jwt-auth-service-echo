package helper

import (
	"fmt"
	"jwt-auth-service/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(payload interface{}) (string, error) {

	// create new token with header
	token := jwt.New(jwt.SigningMethodHS256)

	// get config
	cfg, _ := config.LoadConfig("./app.env")
	ttl := cfg.AccessTokenExpiresIn
	secret := cfg.AccessTokenSecret

	// create claims
	claims := token.Claims.(jwt.MapClaims)

	now := time.Now().UTC()
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	// signed token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("gagal generate token %w", err)
	}

	return tokenString, nil

}

func ValidateToken(token string) (interface{}, error) {

	// get config
	cfg, _ := config.LoadConfig("./app.env")
	secret := cfg.AccessTokenSecret

	// parse token
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil

}
