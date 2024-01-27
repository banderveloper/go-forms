package jwthandler

import (
	"github.com/banderveloper/go-forms/internal/config"
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
