package store

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Storage interface {
	Set()
	Get()
	Status() (string, error)
}

type storage struct {
	connector *redis.Client
	ctx       context.Context
}

func (s *storage) Status() (string, error) {
	status := s.connector.Ping(s.ctx)
	val, err := status.Result()

	return val, err
}

func (s *storage) Set() {
}

func (s *storage) Get() {
}

func NewStorage() Storage {
	s := storage{}
	s.connector = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	s.ctx = context.Background()
	return &s
}
