func TestGetAllBlog(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Inisialisasi Mock
    mockProject := &MockProject{
        Usecase: &MockUsecase{},
        Storage: &MockStorage{}, // Inisialisasi dengan MockStorage
    }

    // Inisialisasi handler
    blogHandler := NewHandler(mockProject) // Pastikan mockProject sesuai tipe yang diterima

    // Panggil fungsi
    blogHandler.GetAllBlog(c)

    // Assert respons
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBlog(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    // Inisialisasi Mock
    mockProject := &MockProject{
        Usecase: &MockUsecase{},
        Storage: &MockStorage{}, // Inisialisasi dengan MockStorage
    }

    // Inisialisasi handler
    blogHandler := NewHandler(mockProject) // Pastikan mockProject sesuai tipe yang diterima

    // Simulasi input
    c.Request, _ = http.NewRequest(http.MethodPost, "/blogs", nil)

    // Panggil fungsi
    blogHandler.CreateBlog(c)

    // Assert respons
    assert.Equal(t, http.StatusCreated, w.Code)
}
