package main

import (
	"context"
	"url-shortener/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/middleware"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	collection := config.ConnectMongo()
	rdb := config.ConnectRedis()

	ctx := context.Background()

	repo := &repository.URLRepository{Collection: collection}

	service := &service.URLService{
		Repo: repo,
		RDB:  rdb,
		Ctx:  ctx,
	}

	handler := &handler.URLHandler{Service: service}

	// Token bucket limiter
	tokenLimiter := &middleware.TokenBucketLimiter{
		RDB:        rdb,
		Ctx:        ctx,
		Capacity:   10,
		RefillRate: 1,
	}

	r := gin.Default()

	// Apply limiter globally
	r.Use(tokenLimiter.Limit())

	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)

	r.Run(":9002")
}
