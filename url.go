package shorturl

import "time"

type URLService interface {
	Get(string) (string, error)
	Set(string, string, time.Duration) error
}
