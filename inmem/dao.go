package inmem

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Dao struct{
	Db *redis.Client
}

func (d *Dao) Get(key string) (string, error) {
	val, err := d.Db.Get(context.Background(), key).Result()
	return val, err
}

func (d *Dao) Set(key string, value string) error {
	err := d.Db.Set(context.Background(), key, value, 0).Err()
	return err
}
