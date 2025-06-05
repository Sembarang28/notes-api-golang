package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenKey  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshTokenKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
)

type Token struct {
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func NewAccessToken(UserID string) (string, error) {
	claims := AccessTokenClaims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute).UTC()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(accessTokenKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewRefreshToken(SessionID string) (*Token, error) {
	iat := time.Now().UTC()
	exp := iat.Add(7 * 24 * time.Hour).UTC() // 7 days expiration

	claims := RefreshTokenClaims{
		SessionID: SessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(refreshTokenKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:     tokenString,
		IssuedAt:  iat,
		ExpiresAt: exp,
	}, nil
}

func VerifyRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnauthorized, err)
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("%w: invalid token", ErrUnauthorized)
	}

	// Expiry check is already handled by jwt.ParseWithClaims if RegisteredClaims is used correctly
	return claims, nil
}

func VerifyAccessToken(tokenStr string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnauthorized, err)
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("%w: invalid token", ErrUnauthorized)
	}

	// Expiry check is already handled by jwt.ParseWithClaims if RegisteredClaims is used correctly
	return claims, nil
}
