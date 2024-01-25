package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/josh1248/forum-website-backend/internal/controllers"
)

func SetUserRoutes(r *gin.Engine) {
	r.GET("api/users", controllers.GetAllUsers)
	r.GET("api/users/:name", controllers.GetUserByName)
	//test route
	r.GET("api/check-cookie", controllers.CheckCookie)
	r.POST("api/new/user", controllers.CreateUser)
	r.POST("api/auth/login", controllers.VerifyUser)
}

func SetPostRoutes(r *gin.Engine) {
	r.GET("api/posts", controllers.GetAllPosts)
	r.POST("api/new/post", controllers.CreatePost)
	r.DELETE("api/delete/post", controllers.DeletePost)
}

func SetCommentRoutes(r *gin.Engine) {
	/*
		r.GET("api/posts/:id", controllers.GetCommentsOfPost)
		r.POST("api/new/comment", controllers.CreateComment)
		r.DELETE("api/delete/comment", controllers.DeleteComment)
	*/
}
