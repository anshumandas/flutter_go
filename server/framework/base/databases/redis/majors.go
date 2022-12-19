package redisdb

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	exists, err := Exists(key)

	var data []byte

	if exists {
		data, err = redis.Bytes(conn.Do("GET", key))
		CheckError(err)
	} else {
		if err != nil {
			CheckError(err)
		}
		err = fmt.Errorf("key does not exist")
	}

	return data, err
}

func Set(key string, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	CheckError(err)

	return err
}

func Exists(key string) (bool, error) {

	conn := Pool.Get()
	defer conn.Close()

	var err error
	var isOK bool

	isOK, err = redis.Bool(conn.Do("EXISTS", key))
	CheckError(err)

	return isOK, err
}

//This function should not be available to the application API and should only be used for internal framework usage
func Delete(key string) error {
	//TODO for items delete only sets state as inactive and does not delete
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	fmt.Println(key)
	return err
}

func GetKeys(pattern string) ([]string, error) {
	conn := Pool.Get()
	defer conn.Close()

	var (
		i    int
		keys []string
	)

	for {
		arr, err := redis.Values(conn.Do("SCAN", i, "MATCH", pattern))
		if err != nil {
			log.Printf("failed to retrieve keys: %v", pattern)
		}
		i, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if i == 0 {
			break
		}
	}

	return keys, nil
}
