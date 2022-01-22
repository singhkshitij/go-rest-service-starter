package cache

import (
	"context"
	"fmt"
	"github.com/singhkshitij/golang-rest-service-starter/src/schema"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
)

var client *redis.Client

func Setup() error {

	if config.RedisConf().Enabled {
		addr := config.RedisConf().Host + ":" + config.RedisConf().Port
		client = redis.NewClient(&redis.Options{
			Addr: addr,
			OnConnect: func(ctx context.Context, cn *redis.Conn) error {
				logger.Info("Connection to redis established")
				return nil
			},
		})

		pong, err := client.Ping(context.Background()).Result()
		logger.Info("Redis ", logger.KV("Ping response ", pong), logger.KV("error ", err))
		if err != nil {
			return err
		}
		return nil
	} else {
		logger.Info("Skipping redis connection as its disabled with ", logger.KV("REDIS_ENABLED", config.RedisConf().Enabled))
		return nil
	}
}

func AddNewTweetToJob(redisReplyJobKey string, tweet schema.StreamTweet) (bool, error) {
	memberVal := fmt.Sprintf("%s:%s", tweet.Data.AuthorId, tweet.Data.ConversationId)
	cmdResult := client.ZAddNX(context.Background(), redisReplyJobKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: memberVal,
	})
	result, err := cmdResult.Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}
