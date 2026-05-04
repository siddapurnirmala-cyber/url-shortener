package repository

import (
	"context"
	"time"
	"url-shortener/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URLRepository struct {
	Collection *mongo.Collection
}

// 🔥 Get next incremental ID
func (r *URLRepository) GetNextID() (int64, error) {

	var result struct {
		Seq int64 `bson:"seq"`
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	err := r.Collection.Database().Collection("counters").
		FindOneAndUpdate(
			context.Background(),
			bson.M{"_id": "url_id"},
			bson.M{"$inc": bson.M{"seq": 1}},
			opts,
		).Decode(&result)

	if err != nil {
		return 0, err
	}

	return result.Seq + 1, nil
}

// Insert with generated code
func (r *URLRepository) InsertWithCode(id int64, originalURL, code string) error {
	_, err := r.Collection.InsertOne(context.Background(), bson.M{
		"id":           id,
		"original_url": originalURL,
		"short_code":   code,
		"created_at":   time.Now(),
	})
	return err
}

// Fetch URL
func (r *URLRepository) GetByCode(code string) (string, error) {

	var result model.URL

	err := r.Collection.FindOne(context.Background(),
		bson.M{"short_code": code},
	).Decode(&result)

	return result.OriginalURL, err
}
