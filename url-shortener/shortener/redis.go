package shortener

import (
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	c *redis.Client
}

func NewRedis(addr string) (*Redis, error) {
	c := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := c.Ping().Err(); err != nil {
		return nil, err
	}
	return &Redis{c}, nil
}

func (r *Redis) Put(url string) string {
	id := generateID(url)
	if err := r.c.Set(id, url, 0); err != nil {
		log.Println(err)
		return ""
	}
	return id
}

func (r *Redis) Get(id string) string {
	url, err := r.c.Get(id).Result()
	switch {
	case err == redis.Nil:
		return ""
	case err != nil:
		log.Println(err)
		return ""
	}
	return url
}
