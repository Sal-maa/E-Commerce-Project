package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(username string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("iya")

func AuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userName"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
