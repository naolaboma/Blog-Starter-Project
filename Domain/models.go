package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID                       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username                 string             `json:"username" bson:"username" validate:"required,min=3,max=50"`
	Email                    string             `json:"email" bson:"email" validate:"required,email"`
	Password                 string             `json:"password,omitempty" bson:"password" validate:"required,min=6"`
	FirstName                string             `json:"firstName" bson:"firstName"`
	LastName                 string             `json:"lastName" bson:"lastName"`
	Bio                      string             `json:"bio" bson:"bio"`
	ProfilePicture           string             `json:"profilePicture" bson:"profilePicture"`
	Role                     string             `json:"role" bson:"role" validate:"oneof=user admin"`
	IsActive                 bool               `json:"isActive" bson:"isActive"`
	IsEmailVerified          bool               `json:"isEmailVerified" bson:"isEmailVerified"`
	EmailVerificationToken   string             `json:"-" bson:"emailVerificationToken,omitempty"`
	EmailVerificationExpires *time.Time         `json:"-" bson:"emailVerificationExpires,omitempty"`
	PasswordResetToken       string             `json:"-" bson:"passwordResetToken,omitempty"`
	PasswordResetExpires     *time.Time         `json:"-" bson:"passwordResetExpires,omitempty"`
	LastLogin                *time.Time         `json:"lastLogin" bson:"lastLogin"`
	CreatedAt                time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt                time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// BlogPost represents a blog post
type BlogPost struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title         string             `json:"title" bson:"title" validate:"required,min=1,max=200"`
	Content       string             `json:"content" bson:"content" validate:"required,min=1"`
	Excerpt       string             `json:"excerpt" bson:"excerpt"`
	AuthorID      primitive.ObjectID `json:"authorId" bson:"authorId" validate:"required"`
	AuthorName    string             `json:"authorName" bson:"authorName"`
	Tags          []string           `json:"tags" bson:"tags"`
	Category      string             `json:"category" bson:"category"`
	Status        string             `json:"status" bson:"status" validate:"oneof=draft published archived"`
	ViewCount     int                `json:"viewCount" bson:"viewCount"`
	LikeCount     int                `json:"likeCount" bson:"likeCount"`
	DislikeCount  int                `json:"dislikeCount" bson:"dislikeCount"`
	CommentCount  int                `json:"commentCount" bson:"commentCount"`
	FeaturedImage string             `json:"featuredImage" bson:"featuredImage"`
	Slug          string             `json:"slug" bson:"slug"`
	ReadingTime   int                `json:"readingTime" bson:"readingTime"`
	PublishedAt   *time.Time         `json:"publishedAt" bson:"publishedAt"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Comment represents a comment on a blog post
type Comment struct {
	ID              primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	PostID          primitive.ObjectID  `json:"postId" bson:"postId" validate:"required"`
	AuthorID        primitive.ObjectID  `json:"authorId" bson:"authorId" validate:"required"`
	AuthorName      string              `json:"authorName" bson:"authorName"`
	Content         string              `json:"content" bson:"content" validate:"required,min=1"`
	ParentCommentID *primitive.ObjectID `json:"parentCommentId,omitempty" bson:"parentCommentId,omitempty"`
	LikeCount       int                 `json:"likeCount" bson:"likeCount"`
	DislikeCount    int                 `json:"dislikeCount" bson:"dislikeCount"`
	IsEdited        bool                `json:"isEdited" bson:"isEdited"`
	IsDeleted       bool                `json:"isDeleted" bson:"isDeleted"`
	CreatedAt       time.Time           `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt" bson:"updatedAt"`
}

// UserInteraction represents user interactions with blog posts
type UserInteraction struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID          primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	PostID          primitive.ObjectID `json:"postId" bson:"postId" validate:"required"`
	InteractionType string             `json:"interactionType" bson:"interactionType" validate:"oneof=like dislike view"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}

// AuthToken represents authentication tokens
type AuthToken struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	TokenType string             `json:"tokenType" bson:"tokenType" validate:"oneof=access refresh"`
	Token     string             `json:"token" bson:"token" validate:"required"`
	IsRevoked bool               `json:"isRevoked" bson:"isRevoked"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"expiresAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

// Tag represents a tag for blog posts
type Tag struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required,min=1,max=50"`
	Description string             `json:"description" bson:"description"`
	PostCount   int                `json:"postCount" bson:"postCount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

// Category represents a category for blog posts
type Category struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required,min=1,max=100"`
	Description string             `json:"description" bson:"description"`
	Slug        string             `json:"slug" bson:"slug" validate:"required"`
	PostCount   int                `json:"postCount" bson:"postCount"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

// AISuggestion represents AI-generated suggestions
type AISuggestion struct {
	ID               primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	UserID           primitive.ObjectID  `json:"userId" bson:"userId" validate:"required"`
	PostID           *primitive.ObjectID `json:"postId,omitempty" bson:"postId,omitempty"`
	SuggestionType   string              `json:"suggestionType" bson:"suggestionType" validate:"oneof=content title tags improvement"`
	OriginalContent  string              `json:"originalContent" bson:"originalContent"`
	SuggestedContent string              `json:"suggestedContent" bson:"suggestedContent" validate:"required"`
	Keywords         []string            `json:"keywords" bson:"keywords"`
	Confidence       float64             `json:"confidence" bson:"confidence"`
	IsUsed           bool                `json:"isUsed" bson:"isUsed"`
	CreatedAt        time.Time           `json:"createdAt" bson:"createdAt"`
}

// UserSession represents user sessions
type UserSession struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"userId" bson:"userId" validate:"required"`
	SessionID    string             `json:"sessionId" bson:"sessionId" validate:"required"`
	IPAddress    string             `json:"ipAddress" bson:"ipAddress"`
	UserAgent    string             `json:"userAgent" bson:"userAgent"`
	IsActive     bool               `json:"isActive" bson:"isActive"`
	LastActivity time.Time          `json:"lastActivity" bson:"lastActivity"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	ExpiresAt    time.Time          `json:"expiresAt" bson:"expiresAt"`
}

// Request/Response DTOs

// UserRegistrationRequest represents user registration request
type UserRegistrationRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserLoginRequest represents user login request
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// UserLoginResponse represents user login response
type UserLoginResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// BlogPostCreateRequest represents blog post creation request
type BlogPostCreateRequest struct {
	Title         string   `json:"title" validate:"required,min=1,max=200"`
	Content       string   `json:"content" validate:"required,min=1"`
	Excerpt       string   `json:"excerpt"`
	Tags          []string `json:"tags"`
	Category      string   `json:"category"`
	FeaturedImage string   `json:"featuredImage"`
	Status        string   `json:"status" validate:"oneof=draft published"`
}

// BlogPostUpdateRequest represents blog post update request
type BlogPostUpdateRequest struct {
	Title         string   `json:"title" validate:"omitempty,min=1,max=200"`
	Content       string   `json:"content" validate:"omitempty,min=1"`
	Excerpt       string   `json:"excerpt"`
	Tags          []string `json:"tags"`
	Category      string   `json:"category"`
	FeaturedImage string   `json:"featuredImage"`
	Status        string   `json:"status" validate:"omitempty,oneof=draft published archived"`
}

// CommentCreateRequest represents comment creation request
type CommentCreateRequest struct {
	Content         string              `json:"content" validate:"required,min=1"`
	ParentCommentID *primitive.ObjectID `json:"parentCommentId,omitempty"`
}

// SearchRequest represents search request
type SearchRequest struct {
	Query    string   `json:"query"`
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	Author   string   `json:"author"`
	Status   string   `json:"status"`
	SortBy   string   `json:"sortBy" validate:"omitempty,oneof=recent popular views"`
	Page     int      `json:"page" validate:"min=1"`
	Limit    int      `json:"limit" validate:"min=1,max=100"`
}

// PaginationResponse represents paginated response
type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"totalPages"`
	HasNext    bool        `json:"hasNext"`
	HasPrev    bool        `json:"hasPrev"`
}

// APIResponse represents standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Constants
const (
	RoleUser  = "user"
	RoleAdmin = "admin"

	StatusDraft     = "draft"
	StatusPublished = "published"
	StatusArchived  = "archived"

	InteractionLike    = "like"
	InteractionDislike = "dislike"
	InteractionView    = "view"

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	SuggestionTypeContent     = "content"
	SuggestionTypeTitle       = "title"
	SuggestionTypeTags        = "tags"
	SuggestionTypeImprovement = "improvement"

	SortByRecent  = "recent"
	SortByPopular = "popular"
	SortByViews   = "views"
)
