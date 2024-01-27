package jwthandler

import (
	"fmt"
	"github.com/banderveloper/go-forms/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtHandler struct {
	Key        string
	AccessTTL  int
	RefreshTTL int
	Audience   string
	Issuer     string
}

// New jwthandler constructor
func New(cfg *config.Config) *JwtHandler {

	return &JwtHandler{
		Key:        cfg.Jwt.Key,
		AccessTTL:  cfg.Jwt.AccessTTL,
		RefreshTTL: cfg.Jwt.RefreshTTL,
		Audience:   cfg.Jwt.Audience,
		Issuer:     cfg.Jwt.Issuer,
	}
}

// GetAccessToken creates new JWT-token with user id and token type access claim
func (jwtHandler *JwtHandler) GetAccessToken(userId int) (string, error) {
	const op = "jwthandler.GetAccessToken"

	// create unsigned token with needed claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  jwtHandler.Issuer,
		"sub":  userId,
		"aud":  jwtHandler.Audience,
		"exp":  time.Now().Add(time.Second * time.Duration(jwtHandler.AccessTTL)).Unix(),
		"type": "access",
		"iat":  time.Now().Unix(),
		"jti":  time.Now().UnixNano(),
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

// GetRefreshToken creates new refresh token with user id and token type refresh claim
func (jwtHandler *JwtHandler) GetRefreshToken(userId int) (string, error) {
	const op = "jwthandler.GetRefreshToken"

	// create unsigned token with needed claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  jwtHandler.Issuer,
		"sub":  userId,
		"aud":  jwtHandler.Audience,
		"exp":  time.Now().Add(time.Second * time.Duration(jwtHandler.RefreshTTL)).Unix(),
		"type": "refresh",
		"iat":  time.Now().Unix(),
		"jti":  time.Now().UnixNano(),
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

// IsTokenValid checks token validity
func (jwtHandler *JwtHandler) IsTokenValid(tokenStr string) bool {
	const op = "jwthandler.IsTokenValid"

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.Key), nil
	})

	return err == nil && token.Valid
}
