package usecase

import (
	"Blog-API/internal/domain"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUseCase struct {
	blogRepo domain.BlogRepository
	userRepo domain.UserRepository
}

func NewBlogUseCase(
	blogRepo domain.BlogRepository, userRepo domain.UserRepository) domain.BlogUseCase {
	return &blogUseCase{blogRepo: blogRepo, userRepo: userRepo}
}

func (uc *blogUseCase) CreateBlog(blog *domain.Blog, authorID primitive.ObjectID) error {
	author, err := uc.userRepo.GetByID(authorID)
	if err != nil {
		return errors.New("author not found")
	}
	//server generated fields
	blog.ID = primitive.NewObjectID()
	blog.AuthorID = authorID
	blog.AuthorUsername = author.Username
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	// initialize slices and countr to ensure they are not nil
	blog.Comments = []domain.Comment{}
	blog.Likes = []string{}
	blog.Dislikes = []string{}
	blog.ViewCount = 0
	blog.LikeCount = 0
	blog.CommentCount = 0

	return uc.blogRepo.Create(blog)
}

func (uc *blogUseCase) GetBlog(id primitive.ObjectID) (*domain.Blog, error) {

	blog, err := uc.blogRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("blog not found")
	}
	go uc.blogRepo.IncrementViewCount(id)
	return blog, nil
}

func (uc *blogUseCase) GetAllBlogs(page, limit int, sort string) ([]*domain.Blog, int64, error) {
	return uc.blogRepo.GetAll(page, limit, sort)
}

func (uc *blogUseCase) UpdateBlog(id primitive.ObjectID, blogUpdate *domain.Blog, userID primitive.ObjectID, userRole string) (*domain.Blog, error) {
	originalBlog, err := uc.blogRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("blog not found")
	}
	if originalBlog.AuthorID != userID && userRole != domain.RoleAdmin {
		return nil, errors.New("forbidden: you are not authorized to update this post")
	}

	originalBlog.Title = blogUpdate.Title
	originalBlog.Content = blogUpdate.Content
	originalBlog.Tags = blogUpdate.Tags
	originalBlog.UpdatedAt = time.Now()

	return originalBlog, nil
}

func (uc *blogUseCase) DeleteBlog(id primitive.ObjectID, userID primitive.ObjectID, userRole string) error {
	blog, err := uc.blogRepo.GetByID(id)
	if err != nil {
		return errors.New("blog not found")
	}
	if blog.AuthorID != userID && userRole != domain.RoleAdmin {
		return errors.New("forbidden: you are not authorized to delete this post")
	}
	return uc.blogRepo.Delete(id)
}

func (uc *blogUseCase) AddComment(blogID primitive.ObjectID, comment *domain.Comment) error {
	author, err := uc.userRepo.GetByID(comment.AuthorID)
	if err != nil {
		return errors.New("comment author not found")
	}
	comment.ID = primitive.NewObjectID()
	comment.AuthorUsername = author.Username
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	return uc.blogRepo.AddComment(blogID, comment)
}

func (uc *blogUseCase) LikeBlog(blogID primitive.ObjectID, userID string) error {
	blog, err := uc.blogRepo.GetByID(blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	isLiked := containsString(blog.Likes, userID)
	isDisliked := containsString(blog.Dislikes, userID)

	if isLiked {
		return uc.blogRepo.RemoveLike(blogID, userID)
	}

	if isDisliked {
		if err := uc.blogRepo.RemoveDislike(blogID, userID); err != nil {
			return err
		}
	}

	return uc.blogRepo.AddLike(blogID, userID)
}

func (uc *blogUseCase) DislikeBlog(blogID primitive.ObjectID, userID string) error {
	blog, err := uc.blogRepo.GetByID(blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	isLiked := containsString(blog.Likes, userID)
	isDisliked := containsString(blog.Dislikes, userID)

	if isDisliked {
		return uc.blogRepo.RemoveDislike(blogID, userID)
	}
	if isLiked {
		if err := uc.blogRepo.RemoveLike(blogID, userID); err != nil {
			return err
		}
	}
	return uc.blogRepo.AddDislike(blogID, userID)
}

func (uc *blogUseCase) SearchBlogsByTitle(title string, page, limit int) ([]*domain.Blog, int64, error) {
	return uc.blogRepo.SearchByTitle(title, page, limit)
}

func (uc *blogUseCase) SearchBlogsByAuthor(author string, page, limit int) ([]*domain.Blog, int64, error) {
	return uc.blogRepo.SearchByAuthor(author, page, limit)
}

func (uc *blogUseCase) FilterBlogsByTags(tags []string, page, limit int) ([]*domain.Blog, int64, error) {
	return uc.blogRepo.FilterByTags(tags, page, limit)
}

func (uc *blogUseCase) FilterBlogsByDate(startDate, endDate time.Time, page, limit int) ([]*domain.Blog, int64, error) {
	return uc.blogRepo.FilterByDate(startDate, endDate, page, limit)
}

func (uc *blogUseCase) GetPopularBlogs(limit int) ([]*domain.Blog, error) {
	return uc.blogRepo.GetPopular(limit)
}

func (uc *blogUseCase) DeleteComment(blogID, commentID primitive.ObjectID, userID primitive.ObjectID) error {
	blog, err := uc.blogRepo.GetByID(blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	var commentAuthorID primitive.ObjectID
	found := false
	for _, c := range blog.Comments {
		if c.ID == commentID {
			commentAuthorID = c.AuthorID
			found = true
			break
		}
	}

	if !found {
		return errors.New("comment not found")
	}

	isCommentAuthor := commentAuthorID == userID
	isBlogAuthor := blog.AuthorID == userID
	isAdmin := user.Role == domain.RoleAdmin

	if !isCommentAuthor && !isBlogAuthor && !isAdmin {
		return errors.New("forbideen: you are not authorized to delete this comment")
	}

	return uc.blogRepo.DeleteComment(blogID, commentID)
}

func (uc *blogUseCase) UpdateComment(blogID, commentID primitive.ObjectID, content string, userID primitive.ObjectID) error {
	// Only the original comment author can update their comment.
	blog, err := uc.blogRepo.GetByID(blogID)
	if err != nil {
		return errors.New("blog not found")
	}
	var commentAuthorID primitive.ObjectID
	found := false
	for _, c := range blog.Comments {
		if c.ID == commentID {
			commentAuthorID = c.AuthorID
			found = true
			break
		}
	}
	if !found {
		return errors.New("comment not found")
	}
	if commentAuthorID != userID {
		return errors.New("forbidden: you are not the author of this comment")
	}
	return uc.blogRepo.UpdateComment(blogID, commentID, content)
}

// helper function
func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
