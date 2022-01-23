package v1

import (
	"encoding/json"
	"errors"
	"github.com/singhkshitij/golang-rest-service-starter/schema"
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"strconv"
)

const tweetsPerPage = 20

func GetTweetsForCategory(category string, page string) ([]schema.TweetData, int, error) {
	var startIndex, endIndex int
	var tweets []schema.TweetData

	categoryTweetLength := int(cache.GetListLength(category))

	if page != "" {
		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			return nil, 0, err
		}
		startIndex = tweetsPerPage * pageNumber
		endIndex = startIndex + tweetsPerPage - 1
	} else {
		startIndex = 0
		endIndex = tweetsPerPage - 1
	}

	if startIndex >= categoryTweetLength {
		return nil, 0, errors.New("no more tweets left to be fetched")
	} else if endIndex >= categoryTweetLength {
		endIndex = categoryTweetLength
	}

	results := cache.GetValuesForListKey(category, startIndex, endIndex)
	for _, result := range results {
		var tweet schema.TweetData
		err := json.Unmarshal([]byte(result), &tweet)
		if err != nil {
			logger.Error("Failed to unmarshall redis value to tweet", logger.KV("error", err))
		}
		tweets = append(tweets, tweet)
	}

	return tweets, categoryTweetLength, nil
}
