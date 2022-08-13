package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type MokuJWTClaims struct {
	Roles []string `json:"rle,omitempty"`
	Scope string   `json:"scp,omitempty"`
	Type  string   `json:"typ,omitempty"`
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenManager struct {
	Secret     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func (m *TokenManager) GenerateToken(userId int, roles []string) (token *Token, err error) {
	token = &Token{}

	/*
		Generate access token
	*/
	at := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&MokuJWTClaims{
			roles,
			"user role permission",
			"access",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(m.AccessTTL).Unix(),
				Id:        uuid.New().String(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "Moku",
				NotBefore: time.Now().Unix(),
				Subject:   strconv.Itoa(userId),
			},
		},
	)
	token.AccessToken, err = at.SignedString([]byte(m.Secret))

	if err != nil {
		return &Token{}, err
	}

	/*
		Generate Refresh token
	*/
	rt := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&MokuJWTClaims{
			[]string{},
			"",
			"refresh",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(m.RefreshTTL).Unix(),
				Id:        uuid.NewString(),
				Subject:   strconv.Itoa(userId),
			},
		},
	)
	token.RefreshToken, err = rt.SignedString([]byte(m.Secret))

	if err != nil {
		return &Token{}, err
	}

	return token, nil

}

func (m *TokenManager) Validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpeted signing method: %v", t.Header["alg"])
		}
		return []byte(m.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
