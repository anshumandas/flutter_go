package redisdb

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error occured with redis functions: %v", err)
	}
}
