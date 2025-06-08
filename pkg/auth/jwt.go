package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohit838/inventory-managements-golang/pkg/config"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	cfg *config.Config
}

func NewJWTService(cfg *config.Config) *JWTService {
	return &JWTService{cfg: cfg}
}

func (j *JWTService) GenerateToken(userID int64, tokenType TokenType) (string, error) {
	var secret string
	var expiresIn time.Duration

	switch tokenType {
	case TokenTypeAccess:
		secret = j.cfg.Jwt.Secret
		dur, _ := time.ParseDuration(j.cfg.Jwt.ExpiresIn)
		expiresIn = dur
	case TokenTypeRefresh:
		secret = j.cfg.RefreshToken.Secret
		dur, _ := time.ParseDuration(j.cfg.RefreshToken.ExpiresIn)
		expiresIn = dur
	default:
		return "", errors.New("invalid token type")
	}

	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (j *JWTService) ParseToken(tokenStr string, tokenType TokenType) (*CustomClaims, error) {
	var secret string
	switch tokenType {
	case TokenTypeAccess:
		secret = j.cfg.Jwt.Secret
	case TokenTypeRefresh:
		secret = j.cfg.RefreshToken.Secret
	default:
		return nil, errors.New("invalid token type")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
