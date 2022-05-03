package redisdb

import (
	"log"

	"github.com/flutter_go/settings"
	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func InitRedis() {

	Pool = &redis.Pool{
		MaxIdle:     settings.RedisSettings.MaxIdle,
		MaxActive:   settings.RedisSettings.MaxActive,
		IdleTimeout: settings.RedisSettings.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", settings.RedisSettings.Host)
			if err != nil {
				log.Fatalf("Error occurred while connecting to the redis database, %v", err)
			}

			return conn, err
		},

		// Health check is in docker-compose file instead of here
	}
}
