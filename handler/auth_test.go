package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"codein/project"
	"codein/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockUsecase adalah mock untuk usecase.Usecase
type MockUsecase struct{}

func (m *MockUsecase) GetUserByToken(c *gin.Context, token string) (interface{}, error) {
	if token == "valid_token" {
		return struct{}{}, nil
	}
	return nil, errors.New("invalid token")
}

// MockProject menggantikan Project dengan mock Usecase
func newMockProject() *project.Project {
	return &project.Project{
		Usecase: &usecase.Usecase{
			// Gunakan mock Usecase di sini jika terdapat di usecase
			GetUserByToken: (&MockUsecase{}).GetUserByToken,
		},
	}
}

func newMockHandler() *Handler {
	return &Handler{
		Project: newMockProject(),
	}
}

func TestCheckToken(t *testing.T) {
	// Setup gin dan handler
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler := newMockHandler()
	r.GET("/auth/check", handler.CheckToken)

	// Pengujian tanpa token (harus mengembalikan 401 Unauthorized)
	req := httptest.NewRequest(http.MethodGet, "/auth/check", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Pengujian dengan token yang tidak valid (harus mengembalikan 401 Unauthorized)
	req = httptest.NewRequest(http.MethodGet, "/auth/check", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Pengujian dengan token yang valid (harus mengembalikan 200 OK)
	req = httptest.NewRequest(http.MethodGet, "/auth/check", nil)
	req.Header.Set("Authorization", "Bearer valid_token")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}



// package handler

// import (
//     "errors"
//     "net/http"
//     "net/http/httptest"
//     "testing"

//     "github.com/gin-gonic/gin"
//     "github.com/stretchr/testify/assert"
// )

// // Mock implementation of GetUserByToken for testing
// func (handler *Handler) GetUserByToken(c *gin.Context) (interface{}, error) {
//     token := c.GetHeader("Authorization")
//     if token == "valid-token" {
//         return &User{}, nil // Assuming User is a struct representing a user
//     }
//     return nil, errors.New("invalid token")
// }

// // Test the CheckToken function
// func TestCheckToken(t *testing.T) {
//     gin.SetMode(gin.TestMode)

//     tests := []struct {
//         token        string
//         expectedCode int
//         expectedBody string
//     }{
//         {"valid-token", 200, `{"message":"","code":200,"success":true}`},
//         {"invalid-token", 401, `{"message":"Unauthorized","code":401,"success":false}`},
//     }

//     for _, tt := range tests {
//         // Create a new gin engine and handler for each test
//         r := gin.New()
//         handler := &Handler{}
//         r.GET("/check", handler.CheckToken)

//         // Create a new HTTP request
//         req, _ := http.NewRequest(http.MethodGet, "/check", nil)
//         req.Header.Set("Authorization", tt.token)

//         // Record the response
//         w := httptest.NewRecorder()
//         r.ServeHTTP(w, req)

//         // Assert the response
//         assert.Equal(t, tt.expectedCode, w.Code)
//         assert.JSONEq(t, tt.expectedBody, w.Body.String())
//     }
// }
