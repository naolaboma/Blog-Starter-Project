package jwt

import (
	"fmt"
	"time"

	"Blog-API/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTService struct {
	secretKey     string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService(secretKey string, accessExpiry, refreshExpiry time.Duration) domain.JWTService {
	return &JWTService{
		secretKey:     secretKey,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

// generates a new access token
func (j *JWTService) GenerateAccessToken(userID primitive.ObjectID, email, role string) (string, error) {
	claims := &domain.JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		Exp:    time.Now().Add(j.accessExpiry).Unix(),
		Iat:    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// generates a new refresh token
func (j *JWTService) GenerateRefreshToken(userID primitive.ObjectID, email, role string) (string, error) {
	claims := &domain.JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		Exp:    time.Now().Add(j.refreshExpiry).Unix(),
		Iat:    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// validates a JWT token and returns claims
func (j *JWTService) ValidateToken(tokenString string) (*domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// generates a new access token using a valid refresh token
func (j *JWTService) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := j.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	// Generate new access token
	return j.GenerateAccessToken(claims.UserID, claims.Email, claims.Role)
} 