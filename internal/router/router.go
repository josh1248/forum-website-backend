package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/josh1248/forum-website-backend/internal/routes"
)

func StartRoutes() {
	fmt.Println("Test route now at http://localhost:8080")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "withCredentials"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	routes.SetUserRoutes(router)
	routes.SetPostRoutes(router)
	routes.SetCommentRoutes(router)

	router.Run("localhost:8080")
}
