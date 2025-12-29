package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-scaffold.iserranodev.net/internal/domain"
	"time"
)

type JWTService struct {
	secret string
}

type Claims struct {
	jwt.RegisteredClaims
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

func (s *JWTService) GenerateToken(user *domain.User) (string, error) {
	var roles []string
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		UserId: user.ID,
		Roles:  roles,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *JWTService) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Token inválido")
	}
	return token.Claims.(*Claims), nil
}
