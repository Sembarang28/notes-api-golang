package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessToken  = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	refreshToken = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
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
	tokenString, err := token.SignedString(accessToken)
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
	tokenString, err := token.SignedString(refreshToken)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:     tokenString,
		IssuedAt:  iat,
		ExpiresAt: exp,
	}, nil
}
