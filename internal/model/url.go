package model

import "time"

type URL struct {
	ID          int64     `bson:"id"`
	OriginalURL string    `bson:"original_url"`
	ShortCode   string    `bson:"short_code"`
	CreatedAt   time.Time `bson:"created_at"`
}
