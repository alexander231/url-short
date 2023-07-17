package storage

import "context"

type Storage interface {
	Get(ctx context.Context, id string) (*URLModel, error)
	Set(ctx context.Context, url string) error
}
