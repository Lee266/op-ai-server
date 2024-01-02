package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Lee266/op-ai-server/router"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	r := router.Router()

	// "/"へのGETリクエストのテスト
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseBody map[string][]string
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.ElementsMatch(t, []string{"/api/public", "/api/private", "/api/private-scoped"}, responseBody["api_urls"])

	// "/api/public/"へのGETリクエストのテスト
	req, _ = http.NewRequest("GET", "/api/public/", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Hello from a public endpoint")

	// 存在しないパスへのリクエストのテスト
	req, _ = http.NewRequest("GET", "/nonexistent", nil)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "Not Found")
}
