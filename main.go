package main

import (
	"smh-api/db"
	"smh-api/router"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	g := gin.Default()
	// Middleware
	// g.Use(gin.Logger())
	// g.Use(middlewares.CORSMiddleware())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	g.Use(gin.Recovery())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://127.0.0.1:8080"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	// }))

	// Routes
	router.Init(g)
	s := &http.Server{
		Addr:           ":9090",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//
	// e.Binder = &models.ExposureArticle{}
	// Start server

	s.ListenAndServe()
}
