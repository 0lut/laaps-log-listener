package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var redisdb *redis.Client

const prefix = "laaps_"

func makeApiKeyKey(name string) string {
	return fmt.Sprintf("%skey:%s", prefix, name)
}

func InitRedis() (err error) {
	var REDIS_HOST = "redis-17236.c59.eu-west-1-2.ec2.cloud.redislabs.com"
	var REDIS_PORT = "17236"
	redisdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: "8aEz1mj0zl2sLt3mPfFxcnwPLO9nsnrX", // no password set
		DB:       0,                                  // use default DB
	})

	_, err = redisdb.Ping().Result()
	log.Printf("Connected to %s:%s\n", REDIS_HOST, REDIS_PORT)
	if err != nil {
		return errors.New("initRedis: \n\t" + err.Error())
	}
	return
}

func GetApiKeyOwner(apiKey string) (string, error) {
	apiKeyKey := makeApiKeyKey(apiKey)
	resRedis, errRedis := redisdb.Get(apiKeyKey).Result()
	if errRedis != nil {
		return "", errors.New("getScenario:\n\t" + errRedis.Error())
	}

	owner := string([]byte(resRedis))
	return owner, nil
}

func SetApiKey(name string, key string) error {
	apiKeyKey := makeApiKeyKey(key)
	err := redisdb.Set(apiKeyKey, name, 0).Err()
	return err
}

// func main() {
// 	InitRedis()
// 	SetApiKey("Anket", "d51a127a-023a-11e9-bce9-3c15c2c09304")
// }
