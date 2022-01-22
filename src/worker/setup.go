package worker

import (
	"fmt"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
)

func Setup(){
	SetupStream()
}

func SetupStream() {
	streamSvc := initStream()
	startStreamJob(streamSvc)
	defer SetupStream()
	consumeStreamData(streamSvc)
}

func consumeStreamData(s *StreamSvc) {
	for tweet := range s.API.Stream.GetMessages() {

		// Handle disconnections from twitter
		// https://developer.twitter.com/en/docs/twitter-api/tweets/volume-streams/integrate/handling-disconnections
		if tweet.Err != nil {
			fmt.Printf("got error from twitter: %v", tweet.Err)

			// Notice we "StopStream" and then "continue" the loop instead of breaking.
			// StopStream will close the long running GET request to Twitter's v2 Streaming endpoint by
			// closing the `GetMessages` channel. Once it's closed, it's safe to perform a new network request
			// with `StartStream`
			s.API.Stream.StopStream()
			continue
		}
		result := tweet.Data

		// Here I am printing out the text.
		// You can send this off to a queue for processing.
		// Or do your processing here in the loop
		logger.Info("Consumed Tweet ", logger.KV("Data", result))
	}

	logger.Info("Stopped Stream")
}

func startStreamJob(svc *StreamSvc) {
	logger.Debug("Service initialised ", logger.KV("data", svc))
	rulesForRequest := map[string]string{
		"filters unnecessary tweets":"lang:en -is:retweet -is:reply -is:quote is:verified airdrop",
	}
	rulesReq := svc.BuildRulesRequest(rulesForRequest)

	svc.SetUnmarshalHook(StreamTweet{})
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

func initStream() *StreamSvc {
	streamSvc := StreamSvc{}
	token, err := streamSvc.GenerateStreamToken()
	if err != nil {
		logger.Error("Failed to generate token for stream", logger.KV("Error", err))
	}
	streamSvc.AccessToken = token.AccessToken
	streamSvc.API = streamSvc.CreateStream()
	return &streamSvc
}
