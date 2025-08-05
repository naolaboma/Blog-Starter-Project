package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"Blog-API/internal/domain"
	"Blog-API/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogHandler struct {
	blogUseCase domain.BlogUseCase
	validate    *validator.Validate
}

func NewBlogHandler(blogUseCase domain.BlogUseCase) *BlogHandler {
	return &BlogHandler{
		blogUseCase: blogUseCase,
		validate:    validator.New(),
	}
}

func (h *BlogHandler) CreateBlog(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "User not authenticated"})
		return
	}

	var req domain.CreateBlogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid request data: " + err.Error()})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Validation failed: " + err.Error()})
		return
	}

	blog := &domain.Blog{
		Title:        req.Title,
		Content:      req.Content,
		Tags:         req.Tags,
		AuthorID:     userID,
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
		Likes:        []string{},
		Dislikes:     []string{},
		Comments:     []domain.Comment{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := h.blogUseCase.CreateBlog(blog, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    blog,
	})
}

func (h *BlogHandler) UpdateBlog(c *gin.Context) {
	userID, userExists := middleware.GetUserIDFromContext(c)
	userRole, roleExists := middleware.GetUserRoleFromContext(c)
	if !userExists || !roleExists {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "User not authenticated"})
		return
	}

	blogID := c.Param("id")

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid blog ID"})
		return
	}

	var req domain.UpdateBlogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid request data: " + err.Error()})
		return
	}
	blogUpdate := &domain.Blog{
		UpdatedAt: time.Now(),
	}

	if req.Title != nil {
		blogUpdate.Title = *req.Title
	}
	if req.Content != nil {
		blogUpdate.Content = *req.Content
	}
	if req.Tags != nil {
		blogUpdate.Tags = *req.Tags
	}

	// userRole := domain.RoleUser

	blogUpdate, err = h.blogUseCase.UpdateBlog(id, blogUpdate, userID, userRole)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		} else if strings.Contains(err.Error(), "forbidden") {
			status = http.StatusForbidden
		}
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
		"blog":    blogUpdate,
	})
}

func (h *BlogHandler) DeleteBlog(c *gin.Context) {
	userID, userExists := middleware.GetUserIDFromContext(c)
	userRole, roleExists := middleware.GetUserRoleFromContext(c)
	if !userExists || !roleExists {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "User not authenticated"})
		return
	}

	blogID := c.Param("id")

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid blog ID"})
		return
	}

	//userRole := domain.RoleUser

	err = h.blogUseCase.DeleteBlog(id, userID, userRole)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
		} else if strings.Contains(err.Error(), "forbidden") {
			status = http.StatusForbidden
		}
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})
}

func (h *BlogHandler) SearchBlogsByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Title parameter is required"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	blogs, total, err := h.blogUseCase.SearchBlogsByTitle(title, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, domain.PaginationResponse{
		Data:       blogs,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *BlogHandler) SearchBlogsByAuthor(c *gin.Context) {
	author := c.Query("author")
	if author == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Author parameter is required"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	blogs, total, err := h.blogUseCase.SearchBlogsByAuthor(author, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, domain.PaginationResponse{
		Data:       blogs,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	blogID := c.Param("id")

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid blog ID"})
		return
	}

	blog, err := h.blogUseCase.GetBlog(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "blog not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blog": blog,
	})
}

func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetAllBlogs endpoint",
	})
}

func (h *BlogHandler) GetPopularBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetPopularBlogs endpoint",
	})
}

func (h *BlogHandler) FilterBlogsByTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "FilterBlogsByTags endpoint",
	})
}

func (h *BlogHandler) FilterBlogsByDate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "FilterBlogsByDate endpoint",
	})
}

func (h *BlogHandler) AddComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Addcomment endpoint",
	})
}

func (h *BlogHandler) DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DeleteComment endpoint",
	})
}

func (h *BlogHandler) UpdateComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateComment endpoint",
	})
}

func (h *BlogHandler) LikeBlog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeBlog endpoint",
	})
}

func (h *BlogHandler) DislikeBlog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DislikeBlog endpoint",
	})
}
