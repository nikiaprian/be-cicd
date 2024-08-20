package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"codein/handler"
	"codein/project"
)

// MockProject untuk mock fungsionalitas project
type MockProject struct {
	mock.Mock
}

func (m *MockProject) SomeMethod() {
	m.Called()
}

func TestCheckToken(t *testing.T) {
	mockProject := new(MockProject)
	h := handler.NewHandler(mockProject)

	// Setup test context
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodGet, "/check-token", nil)

	// Simulasikan perilaku metode yang dipanggil di handler
	// Contoh: mockProject.On("SomeMethod").Return(nil)

	// Call CheckToken
	h.CheckToken(c)

	// Assertions
	assert.Equal(t, http.StatusOK, c.Writer.Status())
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
