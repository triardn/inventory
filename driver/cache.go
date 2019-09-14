package driver

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type CacheOption struct {
	Host               string
	Port               int
	DialConnectTimeout time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	Wait               bool
	MaxConnLifetime    time.Duration
	Password           string
	Namespace          string
}

func NewCache(option CacheOption) (*redis.Pool, error) {
	/* dialConnectTimeoutOption := redis.DialConnectTimeout(option.DialConnectTimeout)
	readTimeoutOption := redis.DialReadTimeout(option.ReadTimeout)
	writeTimeoutOption := redis.DialWriteTimeout(option.WriteTimeout) */

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(fmt.Sprintf("redis://%s@%s:%d", option.Password, option.Host, option.Port))
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", option.Password); err != nil {
				c.Close()
				return nil, err
			}

			if _, err := c.Do("SELECT", option.Namespace); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         option.MaxIdle,
		MaxActive:       option.MaxActive,
		IdleTimeout:     option.IdleTimeout * time.Second,
		Wait:            option.Wait,
		MaxConnLifetime: option.MaxConnLifetime * time.Second,
	}

	_, err := pool.Get().Do("PING")
	if err != nil {
		return nil, err
	}

	return pool, nil
}
