package stream

import (
	twitterstream "github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
	"github.com/fallenstedt/twitter-stream/stream"
	"github.com/singhkshitij/golang-rest-service-starter/schema"
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"strconv"
)

func CleanUpStreamRules() *twitterstream.TwitterApi {
	streamSvc := InitStream()
	rulesFetched, err := GetAllRules(streamSvc)
	if err != nil {
		logger.Error("Failed to get all existing rules", logger.KV("Error", err))
		panic(err)
	}
	var ruleIds []int
	for _, rule := range rulesFetched.Data {
		numericRuleId, err := strconv.Atoi(rule.Id)
		if err != nil {
			logger.Error("Failed to convert rule id to int", logger.KV("Error", err))
		}
		ruleIds = append(ruleIds, numericRuleId)
	}
	if len(ruleIds) > 0 {
		err = DeleteAllRules(ruleIds, streamSvc)
		if err != nil {
			logger.Error("Failed to delete existing rules", logger.KV("Error", err))
		}
	}
	logger.Info("Cleared existing rules with ids ", logger.KV("rule ids", ruleIds))
	return streamSvc
}

func ConsumeStreamData(s stream.IStream) {
	for tweet := range s.GetMessages() {

		// Handle disconnections from twitter
		// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/integrate/handling-disconnections
		if tweet.Err != nil {
			logger.Error("got error from twitter while reading tweet: ", logger.KV("error", tweet.Err))

			// Notice we "StopStream" and then "continue" the loop instead of breaking.
			// StopStream will close the long running GET request to Twitter's v2 Streaming endpoint by
			// closing the `GetMessages` channel. Once it's closed, it's safe to perform a new network request
			// with `StartStream`
			s.StopStream()
			continue
		}

		addTweetToCache(tweet.Data)
	}

	logger.Info("Stopped Stream")
}

func addTweetToCache(tweetData interface{}) {
	tweet, ok := tweetData.(schema.TweetData)
	if ok {
		for _, ruleMatched := range tweet.MatchingRules {
			_, err := cache.AddNewTweetToJob(ruleMatched.Tag, tweet)
			if err != nil {
				logger.Error("Failed to add tweet to redis cache", logger.KV("Error", err))
			}
		}
	} else {
		logger.Error("Failed to parse tweet to stream tweet")
	}
}

func CreateStreamRules(svc *twitterstream.TwitterApi) *rules.TwitterRuleResponse {
	rulesForRequest := map[string]string{
		"airdropRule": "lang:en -is:retweet -is:reply -is:quote is:verified airdrop",
		"jobRule":     "lang:en -is:retweet -is:reply -is:quote is:verified job",
	}
	rulesReq := BuildRulesRequest(rulesForRequest)

	rule, err := CreateRule(&rulesReq, svc)
	if err != nil {
		logger.Error("Failed to generate rule for stream", logger.KV("Error", err))
		panic(err)
	}
	logger.Info("Rule created ", logger.KV("Data", rule))
	return rule
}

func StartStreamJob() stream.IStream {

	streamSvc := InitStream()

	streamExpansions := twitterstream.NewStreamQueryParamsBuilder().
		AddExpansion("author_id,entities.mentions.username").
		AddTweetField("created_at,author_id,attachments,conversation_id,entities,geo,id,lang,public_metrics,referenced_tweets,reply_settings,source,text,withheld").
		AddUserField("verified,username,url,public_metrics,profile_image_url,name,location,id,entities,description,created_at").
		AddMediaField("preview_image_url,type,url,width,height,alt_text").
		AddPlaceField("country,country_code,full_name,name,place_type").
		Build()

	SetUnmarshalHook(streamSvc)

	err := streamSvc.Stream.StartStream(streamExpansions)
	if err != nil {
		logger.Error("Failed to start stream", logger.KV("Error", err))
		panic(err)
	}
	logger.Info("Started stream....")
	return streamSvc.Stream
}

func InitStream() *twitterstream.TwitterApi {
	token, err := GenerateStreamToken()
	if err != nil {
		logger.Error("Failed to generate token for stream", logger.KV("Error", err))
	}
	streamSvc := CreateStream(token.AccessToken)
	return streamSvc
}
