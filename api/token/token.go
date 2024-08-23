package token

import (
	"auth/config"
	"auth/model"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToke(user *model.UserInfo) (string, error) {
	cfg := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim{
		Id:   user.Id,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	})

	tokemString, err := token.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		return "", fmt.Errorf("tokenni imzolashda xatolik: %v", err)
	}
	return tokemString, nil
}

func GenerateRefreshToken(user *model.UserInfo) (string, error) {
	cfg := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim{
		Id:   user.Id,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	tokemString, err := token.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		return "", fmt.Errorf("tokenni imzolashda xatolik")
	}
	return tokemString, nil
}

func ExtractClaimToken(stringToken string) (*Claim, error) {
	cfg := config.LoadConfig()
	token, err := jwt.ParseWithClaims(stringToken, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SIGNING_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
