package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title          string             `bson:"title" json:"title" validate:"required,min=1,max=200"`
	Content        string             `bson:"content" json:"content" validate:"required,min=1"`
	AuthorID       primitive.ObjectID `bson:"author_id" json:"author_id"`
	AuthorUsername string             `bson:"author_username" json:"author_username"`
	Tags           []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	ViewCount      int                `bson:"view_count" json:"view_count"`
	LikeCount      int                `bson:"like_count" json:"like_count"`
	CommentCount   int                `bson:"comment_count" json:"comment_count"`
	Likes          []string           `bson:"likes,omitempty" json:"likes,omitempty"`
	Dislikes       []string           `bson:"dislikes,omitempty" json:"dislikes,omitempty"`
	Comments       []Comment          `bson:"comments,omitempty" json:"comments,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type Comment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AuthorID       primitive.ObjectID `bson:"author_id" json:"author_id"`
	AuthorUsername string             `bson:"author_username" json:"author_username"`
	Content        string             `bson:"content" json:"content" validate:"required,min=1"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type Reaction struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BlogID       primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	ReactionType string             `bson:"reaction_type" json:"reaction_type"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

const (
	ReactionLike    = "like"
	ReactionDislike = "dislike"
)

type Tag struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required,min=1,max=50"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type BlogRepository interface {
	Create(blog *Blog) error
	GetByID(id primitive.ObjectID) (*Blog, error)
	GetAll(page, limit int, sort string) ([]*Blog, int64, error)
	Update(blog *Blog) error
	Delete(id primitive.ObjectID) error
	SearchByTitle(title string, page, limit int) ([]*Blog, int64, error)
	SearchByAuthor(author string, page, limit int) ([]*Blog, int64, error)
	FilterByTags(tags []string, page, limit int) ([]*Blog, int64, error)
	FilterByDate(startDate, endDate time.Time, page, limit int) ([]*Blog, int64, error)
	GetPopular(limit int) ([]*Blog, error)
	IncrementViewCount(id primitive.ObjectID) error
	AddComment(blogID primitive.ObjectID, comment *Comment) error
	DeleteComment(blogID, commentID primitive.ObjectID) error
	UpdateComment(blogID, commentID primitive.ObjectID, content string) error
	AddLike(blogID primitive.ObjectID, userID string) error
	RemoveLike(blogID primitive.ObjectID, userID string) error
	AddDislike(blogID primitive.ObjectID, userID string) error
	RemoveDislike(blogID primitive.ObjectID, userID string) error
	GetTagIDByName(name string) (primitive.ObjectID, error)
}

type BlogUseCase interface {
	CreateBlog(blog *Blog, authorID primitive.ObjectID) error
	GetBlog(id primitive.ObjectID) (*Blog, error)
	GetAllBlogs(page, limit int, sort string) ([]*Blog, int64, error)
	UpdateBlog(id primitive.ObjectID, blog *Blog, userID primitive.ObjectID, userRole string) (*Blog, error)
	DeleteBlog(id primitive.ObjectID, userID primitive.ObjectID, userRole string) error
	SearchBlogsByTitle(title string, page, limit int) ([]*Blog, int64, error)
	SearchBlogsByAuthor(author string, page, limit int) ([]*Blog, int64, error)
	FilterBlogsByTags(tags []string, page, limit int) ([]*Blog, int64, error)
	FilterBlogsByDate(startDate, endDate time.Time, page, limit int) ([]*Blog, int64, error)
	GetPopularBlogs(limit int) ([]*Blog, error)
	AddComment(blogID primitive.ObjectID, comment *Comment) error
	DeleteComment(blogID, commentID primitive.ObjectID, userID primitive.ObjectID) error
	UpdateComment(blogID, commentID primitive.ObjectID, content string, userID primitive.ObjectID) error
	LikeBlog(blogID primitive.ObjectID, userID string) error
	DislikeBlog(blogID primitive.ObjectID, userID string) error
}

// We will think about this later

// type ReactionRepository interface {
// 	Create(reaction *Reaction) error
// 	GetByBlogAndUser(blogID, userID primitive.ObjectID) (*Reaction, error)
// 	Update(reaction *Reaction) error
// 	Delete(id primitive.ObjectID) error
// }

// type TagRepository interface {
// 	Create(tag *Tag) error
// 	GetByName(name string) (*Tag, error)
// 	List() ([]*Tag, error)
// }

type CreateBlogRequest struct {
	Title   string   `json:"title" validate:"required,min=5,max=255"`
	Content string   `json:"content" validate:"required,min=20"`
	Tags    []string `json:"tags" validate:"omitempty,dive,alphanum,min=2,max=20"`
}

type UpdateBlogRequest struct {
	Title   *string   `json:"title" validate:"omitempty,min=5,max=255"`
	Content *string   `json:"content" validate:"omitempty,min=20"`
	Tags    *[]string `json:"tags" validate:"omitempty,dive,alphanum,min=2,max=20"`
}

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=2000"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=2000"`
}

type ReactToBlogRequest struct {
	ReactionType string `json:"reaction_type" validate:"required,oneof=like dislike"`
}

// We will think about this in the next tasks
// type ListBlogParams struct {
// 	Page       int
// 	Limit      int
// 	SortBy     string   // e.g., "newest", "popularity"
// 	Tags       []string // Filter by tags
// 	Author     string   // Filter by author username
// 	SearchTerm string   // For text search on title/content
// }

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
}
