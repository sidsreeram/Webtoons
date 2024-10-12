package api

import (
	"github.com/gin-gonic/gin"
	"github.com/webtoons/pkg/api/delivery"
	"github.com/webtoons/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	UserHandler *delivery.AuthHandler,
	WebtoonHandler *delivery.WebtoonHandler,
) *ServerHTTP {

	engine := gin.New()
	engine.Use(gin.Logger())

	// User routes
	user := engine.Group("/user")
	{
		user.POST("/register", UserHandler.Register)
		user.POST("/login", UserHandler.Login)
		user.Use(middleware.JWTAuthMiddleware()) // JWT-protected routes for logged-in users
	}

	// Webtoon routes (protected by JWT)
	webtoon := engine.Group("/webtoons")
	{
		webtoon.GET("/", WebtoonHandler.GetAll)       // Get all webtoons
		webtoon.POST("/", WebtoonHandler.Create)      // Create new webtoon
		webtoon.GET("/:id", WebtoonHandler.GetByID)   // Get a webtoon by ID
		webtoon.DELETE("/:id", WebtoonHandler.Delete) // Delete a webtoon by ID
	}

	return &ServerHTTP{engine: engine}
}
func (s *ServerHTTP) Start() {
	// s.engine.LoadHTMLGlob("template/*.html")
	s.engine.Run(":3000")
}
