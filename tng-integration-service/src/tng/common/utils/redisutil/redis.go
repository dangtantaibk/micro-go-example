package redisutil

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// Cache provides an access to Redis store.
type Cache interface {
	Set(key string, value interface{}, expireTime int64) error
	HSet(key, field string, value interface{}) error
	Get(key string) (string, error)
	HGet(key, field string) ([]byte, error)
	MGet(key []string) ([]interface{}, error)
	Del(key string) (int64, error)
	Expire(key string, expire int64) (bool, error)
	ExpireWithDuration(key string, expire time.Duration)  error
	Ping() error
	Exist(key string) (bool, error)
	Incr(key string) (int64, error)
	TLL(key string) (time.Duration, error)
}

type cache struct {
	client *redis.Client
}

// NewCache returns a new instance of Store.
func NewCache(client *redis.Client) (Cache, error) {
	return &cache{client}, nil
}

// Del deletes the element with the specified key.
func (s *cache) Del(key string) (int64, error) {
	return s.client.Del(key).Result()
}

// Set sets the new element with specified to Store.
func (s *cache) Set(key string, value interface{}, expired int64) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.client.Set(key, data, time.Duration(expired)*time.Second).Result()
	return err
}

func (s *cache) HSet(key, field string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = s.client.HSet(key, field, data).Err()
	return err
}

// Get get element with specified to Cache.
func (s *cache) Get(key string) (string, error) {
	data, err := s.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s *cache) HGet(key, field string) ([]byte, error) {
	val, err := s.client.HGet(key, field).Bytes()
	return val, err
}

func (s *cache) MGet(key []string) ([]interface{}, error) {
	data, err := s.client.MGet(key...).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Expire sets expiration time for the element with specified key.
func (s *cache) Expire(key string, expire int64) (bool, error) {
	value, err := s.client.Expire(key, time.Duration(expire)*time.Second).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return value, nil
}

// Expire with duration sets expiration time for the element with specified key.
func (s *cache) ExpireWithDuration(key string, expire time.Duration) error {
	return s.client.Expire(key, expire).Err()
}

// Ping checks the Redis connection.
func (s *cache) Ping() error {
	if _, err := s.client.Ping().Result(); err != nil {
		return err
	}
	return nil
}

// Exist check key exist.
func (s *cache) Exist(key string) (bool, error) {
	value, err := s.client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return value == 1, nil
}

// Incr value of key
func (s *cache) Incr(key string) (int64, error) {
	result, err := s.client.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// TLL duration of key
func (s *cache) TLL(key string) (time.Duration, error) {
	result, err:= s.client.TTL(key).Result()
	if err != nil {
		return 0, nil
	}
	return result, nil
}