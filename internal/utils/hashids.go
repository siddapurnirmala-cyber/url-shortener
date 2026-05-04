package utils

import (
	"log"

	hashids "github.com/speps/go-hashids/v2"
)

var hd *hashids.HashID

func init() {
	data := hashids.NewData()

	// 🔐 Secret salt (change this in production!)
	data.Salt = "my-super-secret-salt"

	// Minimum length of short URL
	data.MinLength = 6

	var err error
	hd, err = hashids.NewWithData(data)
	if err != nil {
		log.Fatal("Hashids init error:", err)
	}
}

// Encode ID → Hashid
func Encode(id int64) (string, error) {
	hash, err := hd.EncodeInt64([]int64{id})
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Decode (optional, useful later)
func Decode(code string) (int64, error) {
	nums, err := hd.DecodeInt64WithError(code)
	if err != nil || len(nums) == 0 {
		return 0, err
	}
	return nums[0], nil
}
