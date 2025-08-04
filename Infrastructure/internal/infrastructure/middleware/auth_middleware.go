package middleware

import (
	"net/http"
	"strings"

	"Blog-API/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthMiddleware struct {
	jwtService    domain.JWTService
	sessionRepo   domain.SessionRepository
}

func NewAuthMiddleware(jwtService domain.JWTService, sessionRepo domain.SessionRepository) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService:  jwtService,
		sessionRepo: sessionRepo,
	}
}

//  checks if user is authenticated
func (a *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Authorization token required"})
			c.Abort()
			return
		}

		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Invalid or expired token"})
			c.Abort()
			return
		}

		// Additional security: Check if session exists and is active
		session, err := a.sessionRepo.GetByUserID(claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Session not found"})
			c.Abort()
			return
		}
		
		if !session.IsActive {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Session inactive"})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// checks if user is admin
func (a *AuthMiddleware) AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// First check if user is authenticated
		a.AuthRequired()(c)
		if c.IsAborted() {
			return
		}

		// Then check if user is admin
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "User role not found"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, domain.ErrorResponse{Error: "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// middleware checks for token but doesn't require it
func (a *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.Next()
			return
		}

		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		// Set user info in context if token is valid
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// extractToken extracts JWT token from Authorization header
func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	// Check if it is Bearer token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

// extracts user ID from gin context
func GetUserIDFromContext(c *gin.Context) (primitive.ObjectID, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return primitive.NilObjectID, false
	}

	if id, ok := userID.(primitive.ObjectID); ok {
		return id, true
	}

	return primitive.NilObjectID, false
}

// extracts user role from gin context
func GetUserRoleFromContext(c *gin.Context) (string, bool) {
	role, exists := c.Get("user_role")
	if !exists {
		return "", false
	}

	if r, ok := role.(string); ok {
		return r, true
	}

	return "", false
} 