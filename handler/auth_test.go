package handler_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "codein/handler" // Pastikan ini mengarah ke paket handler Anda
)

// Definisikan struktur User jika belum ada
type User struct {
    ID   int
    Name string
}

// Test the CheckToken function
func TestCheckToken(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        token        string
        expectedCode int
        expectedBody string
    }{
        {"valid-token", 200, `{"message":"","code":200,"success":true}`},
        {"invalid-token", 401, `{"message":"Unauthorized","code":401,"success":false}`},
    }

    for _, tt := range tests {
        // Create a new gin engine and handler for each test
        r := gin.New()
        handler := &handler.Handler{} // Menggunakan struct Handler yang sudah ada
        r.GET("/check", handler.CheckToken)

        // Create a new HTTP request
        req, _ := http.NewRequest(http.MethodGet, "/check", nil)
        req.Header.Set("Authorization", tt.token)

        // Record the response
        w := httptest.NewRecorder()
        r.ServeHTTP(w, req)

        // Assert the response
        assert.Equal(t, tt.expectedCode, w.Code)
        assert.JSONEq(t, tt.expectedBody, w.Body.String())
    }
}
