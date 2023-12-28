package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lee266/op-ai-server/router" // パスを実際のプロジェクトの構造に合わせて変更
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	r := router.Router()

	// "/"へのGETリクエストのテスト
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "dashboard")

	// 存在しないパスへのリクエストのテスト
	req, _ = http.NewRequest("GET", "/nonexistent", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "Not Found")
}