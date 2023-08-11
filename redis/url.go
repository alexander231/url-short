package redis

import (
	"time"
)

type URLService struct {
	db *Storage
}

func NewURLService(db *Storage) *URLService {
	return &URLService{db: db}
}

func (s *URLService) Get(id string) (string, error) {
	return "", nil
}

func (s *URLService) Set(id string, url string, expiration time.Duration) error {
	s.db.client.Set(s.db.ctx, id, url, expiration)
	return nil
}
