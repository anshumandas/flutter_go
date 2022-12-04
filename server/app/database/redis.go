package db

import (
	"log"

	redisdb "github.com/flutter_go/framework/base/databases/redis"
	"github.com/flutter_go/settings"
	"github.com/gomodule/redigo/redis"
)

func InitRedis() {

	redisdb.Pool = &redis.Pool{
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
