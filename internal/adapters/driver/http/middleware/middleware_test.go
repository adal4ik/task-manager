package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const testSecret = "your_secret_key"

func generateTestToken(userID string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestAuthenticateJWT_Success(t *testing.T) {
	userID := "test-user-id"
	token, err := generateTestToken(userID, testSecret)
	assert.NoError(t, err)

	handlerCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		val := r.Context().Value(UserIDKey)
		assert.Equal(t, userID, val)
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rr := httptest.NewRecorder()

	AuthenticateJWT(next).ServeHTTP(rr, req)
	assert.True(t, handlerCalled)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthenticateJWT_MissingHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call next handler")
	})
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	AuthenticateJWT(next).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing Authorization header")
}

func TestAuthenticateJWT_InvalidFormat(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call next handler")
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	rr := httptest.NewRecorder()

	AuthenticateJWT(next).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid Authorization format")
}

func TestAuthenticateJWT_InvalidToken(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call next handler")
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	rr := httptest.NewRecorder()

	AuthenticateJWT(next).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid or expired token")
}

func TestAuthenticateJWT_MissingUserID(t *testing.T) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(testSecret))

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call next handler")
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)
	rr := httptest.NewRecorder()

	AuthenticateJWT(next).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "user_id missing in token")
}
