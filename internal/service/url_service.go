package service

import (
	"context"
	"errors"
	"net/url"
	"time"
	"url-shortener/internal/repository"
	"url-shortener/internal/utils"

	"github.com/redis/go-redis/v9"
)

type URLService struct {
	Repo *repository.URLRepository
	RDB  *redis.Client
	Ctx  context.Context
}

// URL validation
func isValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// Create short URL
func (s *URLService) CreateShortURL(original string) (string, error) {

	if !isValidURL(original) {
		return "", errors.New("invalid URL")
	}

	// 1. Get unique ID
	id, err := s.Repo.GetNextID()
	if err != nil {
		return "", err
	}

	// 2. Convert to Base62
	code := utils.Encode(id)

	// 3. Store in DB
	err = s.Repo.InsertWithCode(id, original, code)
	if err != nil {
		return "", err
	}

	// 4. Cache in Redis
	if s.RDB != nil {
		s.RDB.Set(s.Ctx, code, original, 24*time.Hour)
	}

	return code, nil
}

// Redirect logic
func (s *URLService) GetOriginalURL(code string) (string, error) {

	// Redis first
	if s.RDB != nil {
		val, err := s.RDB.Get(s.Ctx, code).Result()
		if err == nil {
			return val, nil
		}
	}

	// DB fallback
	url, err := s.Repo.GetByCode(code)
	if err != nil {
		return "", err
	}

	// Cache it
	if s.RDB != nil {
		s.RDB.Set(s.Ctx, code, url, 24*time.Hour)
	}

	return url, nil
}
