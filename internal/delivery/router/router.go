package router

import (
	"Blog-API/internal/delivery/controllers"
	"Blog-API/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *controllers.UserHandler, blogHandler *controllers.BlogHandler, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// no auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.RefreshToken)
		}

		// protected auth routes
		authProtected := v1.Group("/auth")
		authProtected.Use(authMiddleware.AuthRequired())
		{
			authProtected.POST("/logout", userHandler.Logout)
		}

		// user routes (authenticated)
		users := v1.Group("/users")
		users.Use(authMiddleware.AuthRequired())
		{
			users.GET("/profile", userHandler.GetProfile)
			users.PUT("/profile", userHandler.UpdateProfile)
		}
		// blog routes
		blogs := v1.Group("/blogs")
		{
			// public routes (no auth)
			blogs.GET("/", blogHandler.GetAllBlogs)
			blogs.GET("/:id", blogHandler.GetBlog)
			blogs.GET("/popular", blogHandler.GetPopularBlogs)

			//search and filter routes
			search := blogs.Group("/search")
			{
				search.GET("/title", blogHandler.SearchBlogsByTitle)
				search.GET("/author", blogHandler.SearchBlogsByAuthor)
			}

			filter := blogs.Group("/filter")
			{
				filter.GET("/tags", blogHandler.FilterBlogsByTags)
				filter.GET("/date", blogHandler.FilterBlogsByDate)
			}

			// protected routes (auth required)
			blogs.Use(authMiddleware.AuthRequired())
			blogs.POST("/", blogHandler.CreateBlog)
			blogs.PUT("/:id", blogHandler.UpdateBlog)
			blogs.DELETE("/:id", blogHandler.DeleteBlog)

			//comments

			blogs.POST("/:id/comments", blogHandler.AddComment)
			blogs.PUT("/:id/comments/:commentId", blogHandler.UpdateComment)
			blogs.DELETE("/:id/comments/:commentId", blogHandler.DeleteComment)

			//Reactions
			blogs.POST("/:id/like", blogHandler.LikeBlog)
			blogs.POST("/:id/dislike", blogHandler.DislikeBlog)
		}
	}

	return router
}
