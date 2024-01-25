package main

import (
	"fmt"

	"github.com/josh1248/forum-website-backend/internal/models"
	"github.com/josh1248/forum-website-backend/internal/router"

	// dependency check with go mod tidy.
	"github.com/josh1248/forum-website-backend/repotest"
)

func main() {
	fmt.Println(repotest.Hello())
	models.ConnectToDB()
	router.StartRoutes()
}
