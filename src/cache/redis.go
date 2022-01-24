package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/singhkshitij/golang-rest-service-starter/schema"
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

func AddNewTweetToJob(redisReplyJobKey string, tweet schema.TweetData) (bool, error) {
	cmdResult := client.LPush(context.Background(), redisReplyJobKey, tweet)
	result, err := cmdResult.Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func GetValuesForListKey(category string, startIndex int, endIndex int) []string {
	return client.LRange(context.Background(), category, int64(startIndex), int64(endIndex)).Val()
}

func GetListLength(category string) int64 {
	return client.LLen(context.Background(), category).Val()
}

func RemoveListItemFromRight(category string) string {
	return client.RPop(context.Background(), category).Val()
}

func GetItemsFromListAtIndex(category string, pos int64) string {
	return client.LIndex(context.Background(), category, pos).Val()
}
