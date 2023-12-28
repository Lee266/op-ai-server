package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEnsureValidToken(t *testing.T) {
	router := gin.Default()

	router.GET("/protected", EnsureValidToken(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Protected resource"})
	})

	// ミドルウェアをテストするためのリクエストを作成
	req, err := http.NewRequest("GET", "/protected", nil)
	assert.NoError(t, err)

	// レスポンスレコーダーをセットアップ
	w := httptest.NewRecorder()

	// リクエストを実行
	router.ServeHTTP(w, req)

	// ステータスコードを検証
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
