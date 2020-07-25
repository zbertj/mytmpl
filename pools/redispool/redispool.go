package redispool

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// LocalRedis0 定义*redis.Client的变量
var LocalRedis0 *redis.Client

// LocalRedis1 定义*redis.Client的变量
var LocalRedis1 *redis.Client

// InitRedis 初始化函数
func InitRedis() error {
	if LocalRedis0 == nil {

		LocalRedis0 = redis.NewClient(&redis.Options{
			Addr:         "root",
			Password:     "000000",
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			PoolSize:     10,
			PoolTimeout:  3 * time.Second,
			DB:           3,
		})
		_, err := LocalRedis0.Ping().Result()
		if err != nil {
			fmt.Println("LocalRedis0 Redis连接失败: ", err)
			return err
		}
	}
	if LocalRedis1 == nil {

		LocalRedis1 = redis.NewClient(&redis.Options{
			Addr:         "root",
			Password:     "000000",
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			PoolSize:     10,
			PoolTimeout:  3 * time.Second,
			DB:           3,
		})
		_, err := LocalRedis1.Ping().Result()
		if err != nil {
			fmt.Println("LocalRedis1 Redis连接失败: ", err)
			return err
		}
	}

	return nil
}

/*
val, err := LocalRedis0.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
*/
