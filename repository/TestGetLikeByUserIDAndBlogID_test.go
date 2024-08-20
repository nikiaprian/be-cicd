func TestGetLikeByUserIDAndBlogID(t *testing.T) {
	mockRepo := new(MockRepository)
	ctx := gin.Context{}

	// Siapkan ekspektasi pada mock
	expectedLike := &models.BlogsLikesResponse{ID: 1, BlogID: 1}
	mockRepo.On("GetLikeByUserIDAndBlogID", &ctx, 1, 1).Return(expectedLike, nil)

	// Panggil metode yang diuji
	result, err := mockRepo.GetLikeByUserIDAndBlogID(&ctx, 1, 1)

	// Periksa hasil
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// Verifikasi bahwa ekspektasi terpenuhi
	mockRepo.AssertExpectations(t)
}