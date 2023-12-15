package main

import (
	"kel15/config"
	"kel15/handler"
	"kel15/project"
	"kel15/repository"
	"kel15/storage"
	"kel15/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := config.NewPostgreSQLDB()
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(&repository)
	storage := storage.NewS3()

	newProject := project.NewProject(usecase, *storage)

	handler := handler.NewHandler(newProject)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://fe.codein.studio"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Auth //
	router.GET("/auth/check-token", handler.CheckToken)
	router.POST("/auth/login", handler.UserLogin)
	router.POST("/auth/register", handler.UserRegister)
	router.GET("/auth/login/:provider", handler.UserLoginByProvider)
	router.GET("/testing-middleware-user", handler.CheckUserRole, handler.TestingMiddlewareUser)
	router.GET("/testing-middleware-admin", handler.CheckAdminRole, handler.TestingMiddlewareAdmin)
	router.GET("/auth/callback/:provider", handler.UserLoginByProviderCallback)

	// User //
	router.GET("/users", handler.UserList)
	router.GET("/user/profile", handler.CheckUserRole, handler.GetUserProfile)
	router.PATCH("/user/update-profile", handler.CheckUserRole, handler.UserProfileUpdate)

	// Blog - Like //
	router.POST("/like/blog/:id", handler.CheckUserRole, handler.CreateLikeByBlogId)
	router.DELETE("/like/blog/:id", handler.CheckUserRole, handler.DeleteLikeByBlogId)

	// Blog //
	router.GET("/blogs", handler.CheckUserLoginOptional, handler.GetAllBlog)
	router.GET("/blogs/:id", handler.CheckUserLoginOptional, handler.GetBlogByID)
	router.POST("/blogs/new", handler.CheckUserRole, handler.CreateBlog)
	// router.PUT("/blogs/:id", handler.CheckUserRole, handler.UpdateBlog)
	router.DELETE("/blogs/:id", handler.CheckUserRole, handler.DeleteBlog)

	// Forum - Like //
	router.POST("/like/forum/:id", handler.CheckUserRole, handler.CreateLikeByForumId)
	router.DELETE("/like/forum/:id", handler.CheckUserRole, handler.DeleteLikeByForumId)

	// Forum //
	router.GET("/forums", handler.CheckUserLoginOptional, handler.GetAllForum)
	router.GET("/forums/:id", handler.CheckUserLoginOptional, handler.GetForumById)
	router.POST("/forums/new", handler.CheckUserRole, handler.CreateForum)
	// router.PUT("/forums/:id", handler.CheckUserRole, handler.UpdateForum)
	router.DELETE("/forums/:id", handler.CheckUserRole, handler.DeleteForum)

	// Forum Comment - Like //
	router.POST("/like/forum/comment/:id", handler.CheckUserRole, handler.CreateLikeByForumCommentId)
	router.DELETE("/like/forum/comment/:id", handler.CheckUserRole, handler.DeleteLikeByForumCommentId)

	// Forum Comment - Selected Answer //
	router.PATCH("/forum/comment/:id/selected-answer", handler.CheckUserRole, handler.SelectedCommentAnswer)

	// Forum Comment //
	router.GET("/commentsforum/:id", handler.CheckUserLoginOptional, handler.GetAllCommentByForumID)
	router.POST("/commentsforum/:id", handler.CheckUserRole, handler.CreateCommentForum)
	router.DELETE("/commentsforum/:id", handler.CheckUserRole, handler.DeleteCommentForum)

	// Blog Comment - Like //
	router.POST("/like/blog/comment/:id", handler.CheckUserRole, handler.CreateLikeByBlogCommentId)
	router.DELETE("/like/blog/comment/:id", handler.CheckUserRole, handler.DeleteLikeByBlogCommentId)

	// Blog Comment //
	router.GET("/comments/:id", handler.CheckUserLoginOptional, handler.GetAllCommentByBlogID)
	router.POST("/comments/:id", handler.CheckUserRole, handler.CreateCommentBlog)
	router.DELETE("/comments/:id", handler.CheckUserRole, handler.DeleteCommentByID)

	// Setup Server //
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
