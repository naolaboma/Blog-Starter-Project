package domain

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// interface for JWT operations
type JWTService interface {
	GenerateAccessToken(userID primitive.ObjectID, email, role string) (string, error)
	GenerateRefreshToken(userID primitive.ObjectID, email, role string) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	RefreshAccessToken(refreshToken string) (string, error)
}

// claims in a JWT token
type JWTClaims struct {
	UserID primitive.ObjectID `json:"user_id"`
	Email  string             `json:"email"`
	Role   string             `json:"role"`
	Exp    int64              `json:"exp"`
	Iat    int64              `json:"iat"`
}

// returns the expiration time
func (c *JWTClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Exp, 0)), nil
}

// returns the not before time
func (c *JWTClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Iat, 0)), nil
}

// returns the issued at time
func (c *JWTClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Iat, 0)), nil
}

// returns the issuer
func (c *JWTClaims) GetIssuer() (string, error) {
	return "blog-api", nil
}

// returns the subject
func (c *JWTClaims) GetSubject() (string, error) {
	return c.UserID.Hex(), nil
}

// returns the audience
func (c *JWTClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

// interface for authentication middleware
type AuthMiddleware interface {
	AuthRequired() func(http.Handler) http.Handler
	AdminRequired() func(http.Handler) http.Handler
	OptionalAuth() func(http.Handler) http.Handler
	ExtractUserFromContext(ctx context.Context) (*User, bool)
}

// interface for email operations
type EmailService interface {
	SendPasswordResetEmail(email, token string) error
	SendWelcomeEmail(email, username string) error
	SendVerificationEmail(email, username, token string) error
}

// defines the interface for password operations
type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
	ValidatePassword(password string) error
	GenerateSecureToken(length int) string
}

// type for context keys
type ContextKey string

const (
	UserContextKey ContextKey = "user"
)

// password reset token
type PasswordResetToken struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
	Used      bool               `bson:"used" json:"used"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// interface for password reset token operations
type PasswordResetTokenRepository interface {
	Create(token *PasswordResetToken) error
	GetByToken(token string) (*PasswordResetToken, error)
	GetByUserID(userID primitive.ObjectID) (*PasswordResetToken, error)
	MarkAsUsed(token string) error
	DeleteExpired() error
}

// request for password reset
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// request for password reset with token
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
