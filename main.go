package main

import (
	"net/http"
	"rest_api_portfolio/app"
	authcontroller "rest_api_portfolio/controller/auth-controller"
	blogcontroller "rest_api_portfolio/controller/blog-controller"
	projectcontroller "rest_api_portfolio/controller/project-controller"
	"rest_api_portfolio/exception"
	"rest_api_portfolio/middleware"
	authrepository "rest_api_portfolio/repository/auth-repository"
	blogrepository "rest_api_portfolio/repository/blog-repository"
	projectrepository "rest_api_portfolio/repository/project-repository"
	authservice "rest_api_portfolio/service/auth-service"
	blogservice "rest_api_portfolio/service/blog-service"
	projectservice "rest_api_portfolio/service/project-service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "/api/v1/*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db := app.DatabaseConnect()
	validate := validator.New()
	router := httprouter.New()

	// blog inisiation
	blogRepository := blogrepository.NewBlogRepositoryIplm()
	blogService := blogservice.NewBlogServiceIplm(validate, db, blogRepository)
	blogController := blogcontroller.NewBlogControllerIplm(blogService)
	// blog router
	router.POST("/api/v1/blogs", blogController.Create)
	router.PUT("/api/v1/blogs/:blog_id", blogController.Update)
	router.DELETE("/api/v1/blogs/:blog_id", blogController.Delete)
	router.GET("/api/v1/blogs/:blog_id", blogController.FindById)
	router.GET("/api/v1/blogs", blogController.FindAll)

	// auth inisiation
	authRepository := authrepository.NewAuthRepositoryIplm()
	authService := authservice.NewAuthServiceIplm(validate, db, authRepository)
	authController := authcontroller.NewAuthControllerIplm(authService)
	// auth router
	router.POST("/api/v1/auth/sign-in", authController.Create)
	router.DELETE("/api/v1/auth/:auth_id", authController.Delete)
	router.GET("/api/v1/auth/:auth_id", authController.FindById)
	router.POST("/api/v1/auth/log-in", authController.CreateToken)
	router.POST("/api/v1/auth/refresh-token", authController.RefreshToken)

	// project inisiation
	projectRepository := projectrepository.NewProjectRepositoryIplm()
	projectService := projectservice.NewProjectServiceIplm(validate, db, projectRepository)
	projectController := projectcontroller.NewProjectControllerIplm(projectService)
	// project router
	router.POST("/api/v1/projects", projectController.Create)
	router.PUT("/api/v1/projects/:project_id", projectController.Update)
	router.DELETE("/api/v1/projects/:project_id", projectController.Delete)
	router.GET("/api/v1/projects/:project_id", projectController.FindById)
	router.GET("/api/v1/projects", projectController.FindAll)

	// middleware and cors
	authMiddleware := middleware.NewAuthMiddleware(router)

	finalHandle := corsHandler(authMiddleware)

	// error handler
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: finalHandle,
	}

	server.ListenAndServe()
}
