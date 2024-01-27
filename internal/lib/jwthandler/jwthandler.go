package jwthandler

import (
	"fmt"
	"time"

	"github.com/banderveloper/go-forms/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtHandler struct {
	Key             string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

// New jwthandler constructor
func New(cfg *config.Config) *JwtHandler {

	return &JwtHandler{
		Key:             cfg.Jwt.Key,
		AccessTokenTTL:  cfg.Jwt.AccessTokenTTL,
		RefreshTokenTTL: cfg.Jwt.RefreshTokenTTL,
	}
}

// GetAccessToken creates new JWT-token with user id and token type access claim
func (jwtHandler *JwtHandler) GetAccessToken(userId int) (string, error) {
	const op = "jwthandler.GetAccessToken"

	// create unsigned token with needed claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  userId,
		"exp":  time.Now().Add(time.Second * time.Duration(jwtHandler.AccessTokenTTL)),
		"type": "access",
	})

	// transform key to required array of bytes
	signKey := []byte(jwtHandler.Key)

	// sign token
	signedToken, err := token.SignedString(signKey)

	// if error during signing
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return signedToken, nil
}
