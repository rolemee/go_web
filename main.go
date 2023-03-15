package main

import (
	"outsourcing/api"
	// "fmt"
	// "github.com/meilisearch/meilisearch-go"
	// "github.com/meilisearch/meilisearch-go"
	"github.com/gin-gonic/gin"
)


func main() {
  // gin.SetMode(gin.ReleaseMode)
  r := gin.Default()
  r.GET("/api/search/all", api.All)
  r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}