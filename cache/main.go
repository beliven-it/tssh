package cache

import (
	"errors"
	"os"
)

type cache struct{}

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Exist(string) bool
	Flush(string)
}

func (c *cache) Flush(key string) {
	os.Remove(key)
}

func (c *cache) Exist(key string) bool {
	_, err := os.Stat(key)
	return !errors.Is(err, os.ErrNotExist)
}

func (c *cache) Set(key string, data []byte) error {
	return os.WriteFile(key, data, 0600)
}

func (c *cache) Get(key string) ([]byte, error) {
	return os.ReadFile(key)
}

func NewCache() Cache {
	return &cache{}
}
