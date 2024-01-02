package router

import (
	"net/http"

	"github.com/Lee266/op-ai-server/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// 使えるAPIのURL一覧を定義
		apiUrls := []string{
			"/api/public",
			"/api/private",
			"/api/private-scoped",
		}

		c.JSON(http.StatusOK, gin.H{"api_urls": apiUrls})
	})

	// 保護されていないエンドポイント
	publicGroup := r.Group("/api/public")
	{
		publicGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello from a public endpoint! You don't need to be authenticated to see this."})
		})
	}

	// トークンの検証が必要なエンドポイント
	privateGroup := r.Group("/api/private")
	privateGroup.Use(middlewares.EnsureValidToken())
	{
		privateGroup.GET("/", func(c *gin.Context) {
			// CORS Headers
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Headers", "Authorization")

			c.JSON(http.StatusOK, gin.H{"message": "Hello from a private endpoint! You need to be authenticated to see this."})
		})
	}

	// トークンの検証とスコープのチェックが必要なエンドポイント
	privateScopedGroup := r.Group("/api/private-scoped")
	privateScopedGroup.Use(middlewares.EnsureValidToken())
	{
		privateScopedGroup.GET("/", func(c *gin.Context) {
			// CORS Headers
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Headers", "Authorization")

			// トークンが必要なスコープを持っているか確認
			tokenClaims, exists := c.Get("claims")
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token claims not found"})
				return
			}

			claims, ok := tokenClaims.(*middlewares.CustomClaims)
			if !ok || !claims.HasScope("read:messages") {
				c.JSON(http.StatusForbidden, gin.H{"message": "Insufficient scope."})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Hello from a private-scoped endpoint! You need to be authenticated and have the 'read:messages' scope to see this."})
		})
	}

	// 存在しないルートの処理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})

	return r
}
