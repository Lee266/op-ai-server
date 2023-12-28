package middlewares

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// トークンから取得したいカスタムデータ。
type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// 特定のスコープを持っているかどうかを確認。
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}
	return false
}

func parseIssuerURL() (*url.URL, error) {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
		return nil, err
	}
	return issuerURL, nil
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("JWTの検証中にエラーが発生しました: %v", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	if err != nil {
		log.Printf("レスポンスの書き込みエラー: %v", err)
	}

	// レスポンスボディに書に込み
	if _, writeErr := w.Write([]byte(`{"message":"JWTの検証に失敗しました。"}`)); writeErr != nil {
		log.Printf("レスポンスボディの書き込みエラー: %v", writeErr)
		log.Printf("レスポンスボディの書き込みエラー: %v", writeErr)
	}
}

func EnsureValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerURL, err := parseIssuerURL()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			log.Fatalf("Failed to set up the jwt validator")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		middleware := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler),
		)

		// Gin handlerをhttp.Handlerに変換
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})

		// リクエスト内のJWTを確認。
		middleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)
	}
}
