package cache

import (
	"context"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type RedisClient struct {
	redis *redis.Client
}

var redisClient *RedisClient
var ctx context.Context

type RedisInterface interface {
	SetRedisVal(key string, data string, duration int) error
	GetRedisVal(key string) (string, error)
	DeleteRedisVal(key string) error
	TestConnection() error
}

func (o *RedisClient) TestConnection() error {
	//if err := redisClient.redis.Ping(ctx); err != nil {
	//	log.Fatal(err)
	//}

	return nil
}

func NewRedisClient(config util.Config) RedisInterface {
	if redisClient == nil {
		redisDb, _ := strconv.Atoi(config.RedisDB)
		client := redis.NewClient(&redis.Options{
			Addr:     config.RedisUrl,
			Password: config.RedisPassword,
			DB:       redisDb,
		})

		redisClient = &RedisClient{
			redis: client,
		}
	}

	return redisClient
}

func (o *RedisClient) SetRedisVal(key string, data string, duration int) error {
	// ttl := time.Duration(3000000) * time.Hour
	ttl := time.Duration(duration) * time.Hour
	logrus.Println("set ttl what second : ", ttl)

	setRedis := o.redis.Set(context.Background(), key, data, ttl)
	if err := setRedis.Err(); err != nil {
		logrus.Info("error set redis : ", setRedis.Err().Error())
		return err
	}

	return nil
}

func (o *RedisClient) GetRedisVal(key string) (string, error) {
	op2 := o.redis.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		return "", err
	}
	result, _ := op2.Result()

	return result, nil
}

func (o *RedisClient) DeleteRedisVal(key string) error {
	op2 := o.redis.Del(context.Background(), key)
	if err := op2.Err(); err != nil {
		return err
	}

	return nil
}
