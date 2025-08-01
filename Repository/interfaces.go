package Repository

import (
	"Mini_PRD/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(ctx context.Context, user *Domain.User) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.User, error)
	GetByEmail(ctx context.Context, email string) (*Domain.User, error)
	GetByUsername(ctx context.Context, username string) (*Domain.User, error)
	Update(ctx context.Context, user *Domain.User) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	UpdateLastLogin(ctx context.Context, id primitive.ObjectID) error
	UpdatePassword(ctx context.Context, id primitive.ObjectID, hashedPassword string) error
	UpdateEmailVerification(ctx context.Context, id primitive.ObjectID, isVerified bool) error
	UpdatePasswordResetToken(ctx context.Context, id primitive.ObjectID, token string, expiresAt *primitive.DateTime) error
	ClearPasswordResetToken(ctx context.Context, id primitive.ObjectID) error
}

// BlogPostRepository defines the interface for blog post data operations
type BlogPostRepository interface {
	Create(ctx context.Context, post *Domain.BlogPost) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.BlogPost, error)
	GetBySlug(ctx context.Context, slug string) (*Domain.BlogPost, error)
	GetByAuthor(ctx context.Context, authorID primitive.ObjectID, page, limit int) ([]*Domain.BlogPost, int64, error)
	GetAll(ctx context.Context, page, limit int, status string) ([]*Domain.BlogPost, int64, error)
	Update(ctx context.Context, post *Domain.BlogPost) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Search(ctx context.Context, query string, tags []string, category, author, status string, sortBy string, page, limit int) ([]*Domain.BlogPost, int64, error)
	GetPopular(ctx context.Context, limit int) ([]*Domain.BlogPost, error)
	GetByTags(ctx context.Context, tags []string, page, limit int) ([]*Domain.BlogPost, int64, error)
	GetByCategory(ctx context.Context, category string, page, limit int) ([]*Domain.BlogPost, int64, error)
	IncrementViewCount(ctx context.Context, id primitive.ObjectID) error
	IncrementLikeCount(ctx context.Context, id primitive.ObjectID) error
	DecrementLikeCount(ctx context.Context, id primitive.ObjectID) error
	IncrementDislikeCount(ctx context.Context, id primitive.ObjectID) error
	DecrementDislikeCount(ctx context.Context, id primitive.ObjectID) error
	IncrementCommentCount(ctx context.Context, id primitive.ObjectID) error
	DecrementCommentCount(ctx context.Context, id primitive.ObjectID) error
}

// CommentRepository defines the interface for comment data operations
type CommentRepository interface {
	Create(ctx context.Context, comment *Domain.Comment) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.Comment, error)
	GetByPostID(ctx context.Context, postID primitive.ObjectID, page, limit int) ([]*Domain.Comment, int64, error)
	GetByAuthor(ctx context.Context, authorID primitive.ObjectID, page, limit int) ([]*Domain.Comment, int64, error)
	Update(ctx context.Context, comment *Domain.Comment) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	SoftDelete(ctx context.Context, id primitive.ObjectID) error
	IncrementLikeCount(ctx context.Context, id primitive.ObjectID) error
	DecrementLikeCount(ctx context.Context, id primitive.ObjectID) error
	IncrementDislikeCount(ctx context.Context, id primitive.ObjectID) error
	DecrementDislikeCount(ctx context.Context, id primitive.ObjectID) error
}

// UserInteractionRepository defines the interface for user interaction data operations
type UserInteractionRepository interface {
	Create(ctx context.Context, interaction *Domain.UserInteraction) error
	GetByUserAndPost(ctx context.Context, userID, postID primitive.ObjectID, interactionType string) (*Domain.UserInteraction, error)
	GetByPost(ctx context.Context, postID primitive.ObjectID, interactionType string) ([]*Domain.UserInteraction, error)
	GetByUser(ctx context.Context, userID primitive.ObjectID, interactionType string, page, limit int) ([]*Domain.UserInteraction, int64, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	DeleteByUserAndPost(ctx context.Context, userID, postID primitive.ObjectID, interactionType string) error
	Exists(ctx context.Context, userID, postID primitive.ObjectID, interactionType string) (bool, error)
}

// AuthTokenRepository defines the interface for authentication token data operations
type AuthTokenRepository interface {
	Create(ctx context.Context, token *Domain.AuthToken) error
	GetByToken(ctx context.Context, token string) (*Domain.AuthToken, error)
	GetByUserAndType(ctx context.Context, userID primitive.ObjectID, tokenType string) ([]*Domain.AuthToken, error)
	RevokeToken(ctx context.Context, token string) error
	RevokeAllUserTokens(ctx context.Context, userID primitive.ObjectID) error
	DeleteExpired(ctx context.Context) error
	Exists(ctx context.Context, token string) (bool, error)
}

// TagRepository defines the interface for tag data operations
type TagRepository interface {
	Create(ctx context.Context, tag *Domain.Tag) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.Tag, error)
	GetByName(ctx context.Context, name string) (*Domain.Tag, error)
	GetAll(ctx context.Context, page, limit int) ([]*Domain.Tag, int64, error)
	GetPopular(ctx context.Context, limit int) ([]*Domain.Tag, error)
	Update(ctx context.Context, tag *Domain.Tag) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	IncrementPostCount(ctx context.Context, id primitive.ObjectID) error
	DecrementPostCount(ctx context.Context, id primitive.ObjectID) error
	ExistsByName(ctx context.Context, name string) (bool, error)
}

// CategoryRepository defines the interface for category data operations
type CategoryRepository interface {
	Create(ctx context.Context, category *Domain.Category) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.Category, error)
	GetBySlug(ctx context.Context, slug string) (*Domain.Category, error)
	GetAll(ctx context.Context, page, limit int) ([]*Domain.Category, int64, error)
	Update(ctx context.Context, category *Domain.Category) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	IncrementPostCount(ctx context.Context, id primitive.ObjectID) error
	DecrementPostCount(ctx context.Context, id primitive.ObjectID) error
	ExistsBySlug(ctx context.Context, slug string) (bool, error)
}

// AISuggestionRepository defines the interface for AI suggestion data operations
type AISuggestionRepository interface {
	Create(ctx context.Context, suggestion *Domain.AISuggestion) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Domain.AISuggestion, error)
	GetByUser(ctx context.Context, userID primitive.ObjectID, page, limit int) ([]*Domain.AISuggestion, int64, error)
	GetByPost(ctx context.Context, postID primitive.ObjectID, page, limit int) ([]*Domain.AISuggestion, int64, error)
	GetByType(ctx context.Context, userID primitive.ObjectID, suggestionType string, page, limit int) ([]*Domain.AISuggestion, int64, error)
	Update(ctx context.Context, suggestion *Domain.AISuggestion) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	MarkAsUsed(ctx context.Context, id primitive.ObjectID) error
}

// UserSessionRepository defines the interface for user session data operations
type UserSessionRepository interface {
	Create(ctx context.Context, session *Domain.UserSession) error
	GetBySessionID(ctx context.Context, sessionID string) (*Domain.UserSession, error)
	GetByUser(ctx context.Context, userID primitive.ObjectID) ([]*Domain.UserSession, error)
	UpdateLastActivity(ctx context.Context, sessionID string) error
	DeactivateSession(ctx context.Context, sessionID string) error
	DeactivateAllUserSessions(ctx context.Context, userID primitive.ObjectID) error
	DeleteExpired(ctx context.Context) error
	Exists(ctx context.Context, sessionID string) (bool, error)
}
