package worker

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/singhkshitij/golang-rest-service-starter/schema"
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"github.com/singhkshitij/golang-rest-service-starter/src/worker/stream"
	"time"
)

func Setup() {
	initialiseListPurgeJob()
	svc := stream.CleanUpStreamRules()
	stream.CreateStreamRules(svc)
	SetupStream()
}

func initialiseListPurgeJob() {
	listDataPurgeJob := gocron.NewScheduler(time.UTC)
	logger.Info("Initializing list data purge job...")
	_, err := listDataPurgeJob.Every(30).Minute().Do(setupOldDataClearJob)
	if err != nil {
		logger.Error("Failed to run data purge job job :", logger.KV("error: ", err))
	}
	listDataPurgeJob.StartAsync()
}

func setupOldDataClearJob() {
	//make is scehduled
	for key := range stream.RulesForRequest {
		removeOldDataForKey(key)
	}
}

func removeOldDataForKey(key string) {
	categoryTweetsLength := int(cache.GetListLength(key))
	sevenDaysBack := time.Now().Add(-168 * time.Hour)
	removedItemsCount := 0
	for i := categoryTweetsLength - 1; i >= 0; i-- {
		var tweet schema.TweetData
		result := cache.GetItemsFromListAtIndex(key, int64(i))
		err := json.Unmarshal([]byte(result), &tweet)
		if err != nil {
			logger.Error("Failed to unmarshall redis value to tweet", logger.KV("error", err))
		}
		if tweet.CachedAt.Unix() < sevenDaysBack.Unix() {
			cache.RemoveListItemFromRight(key)
			removedItemsCount++
		} else {
			break
		}
	}

	logger.Info("Running data purge job for list : "+key, logger.KV(" removed total items : ", removedItemsCount))
}

func SetupStream() {
	createdStream := stream.StartStreamJob()
	defer SetupStream()
	stream.ConsumeStreamData(createdStream)
}
