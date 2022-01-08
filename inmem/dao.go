package inmem

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Dao struct{
	Db *redis.Client
}

func (d *Dao) Get(key string) (Dto, error) {
	var dto Dto
	dto.Key = key

	val, err := d.Db.Get(context.Background(), key).Result()
	if err == redis.Nil {
		dto.Exists = false
		return dto, nil
	}
	dto.Value = val
	return dto, err
}

func (d *Dao) Set(dto Dto) error {
	err := d.Db.Set(context.Background(), dto.Key, dto.Value, 0).Err()
	return err
}
