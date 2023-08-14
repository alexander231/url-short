package redis

import (
	"github.com/pkg/errors"
	"time"
)

type URLService struct {
	db *Storage
}

func NewURLService(db *Storage) *URLService {
	return &URLService{db: db}
}

func (s *URLService) Get(id string) (string, error) {
	val, err := s.db.client.Get(s.db.ctx, id).Result()
	if err != nil {
		return "", errors.Wrap(err, "Getting URL from db")
	}
	return val, nil
}

func (s *URLService) Set(id string, url string, expiration time.Duration) error {
	if err := s.db.client.Set(s.db.ctx, id, url, expiration).Err(); err != nil {
		return errors.Wrap(err, "Saving URL to db")
	}
	return nil
}
