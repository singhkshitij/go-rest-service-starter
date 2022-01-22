package stream

import (
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"github.com/singhkshitij/golang-rest-service-starter/src/schema"
	"strconv"
)

func CleanUpStreamRules() {
	streamSvc := InitStream()
	rules, err := streamSvc.GetAllRules()
	if err != nil {
		logger.Error("Failed to get all existing rules", logger.KV("Error", err))
		panic(err)
	}
	var ruleIds []int
	for _, rule := range rules.Data {
		numericRuleId, err := strconv.Atoi(rule.Id)
		if err != nil {
			logger.Error("Failed to convert rule id to int", logger.KV("Error", err))
		}
		ruleIds = append(ruleIds, numericRuleId)
	}
	if len(ruleIds) > 0 {
		err = streamSvc.DeleteAllRules(ruleIds)
		if err != nil {
			logger.Error("Failed to delete existing rules", logger.KV("Error", err))
		}
	}
	logger.Info("Cleared existing rules with ids ", logger.KV("rule ids", ruleIds))
}

func ConsumeStreamData(s *StreamSvc) {
	for tweet := range s.API.Stream.GetMessages() {

		// Handle disconnections from twitter
		// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/integrate/handling-disconnections
		if tweet.Err != nil {
			logger.Error("got error from twitter while reading tweet: ", logger.KV("error", tweet.Err))

			// Notice we "StopStream" and then "continue" the loop instead of breaking.
			// StopStream will close the long running GET request to Twitter's v2 Streaming endpoint by
			// closing the `GetMessages` channel. Once it's closed, it's safe to perform a new network request
			// with `StartStream`
			s.API.Stream.StopStream()
			continue
		}
		result := tweet.Data.(schema.StreamTweet)

		// Here I am printing out the text.
		// You can send this off to a queue for processing.
		// Or do your processing here in the loop
		logger.Info("Consumed Tweet ", logger.KV("Data", result))
		go addTweetToCache(result)
	}

	logger.Info("Stopped Stream")
}

func addTweetToCache(tweet schema.StreamTweet) {
	for _, ruleMatched := range tweet.MatchingRules {
		_, err := cache.AddNewTweetToJob(ruleMatched.Tag, tweet)
		if err != nil {
			logger.Error("Failed to add tweet to redis cache", logger.KV("Error", err))
		}
	}
}

func StartStreamJob(svc *StreamSvc) {
	rulesForRequest := map[string]string{
		"airdropRule": "lang:en -is:retweet -is:reply -is:quote is:verified airdrop",
		"jobRule":     "lang:en -is:retweet -is:reply -is:quote is:verified job",
	}
	rulesReq := svc.BuildRulesRequest(rulesForRequest)

	svc.SetUnmarshalHook(schema.StreamTweet{})
	rule, err := svc.CreateRule(&rulesReq)
	if err != nil {
		logger.Error("Failed to generate rule for stream", logger.KV("Error", err))
		panic(err)
	}
	err = svc.StartStream()
	if err != nil {
		logger.Error("Failed to start stream", logger.KV("Error", err))
		panic(err)
	}
	logger.Info("Rule created ", logger.KV("Data", rule))
}

func InitStream() *StreamSvc {
	streamSvc := StreamSvc{}
	token, err := streamSvc.GenerateStreamToken()
	if err != nil {
		logger.Error("Failed to generate token for stream", logger.KV("Error", err))
	}
	streamSvc.AccessToken = token.AccessToken
	streamSvc.API = streamSvc.CreateStream()
	return &streamSvc
}
