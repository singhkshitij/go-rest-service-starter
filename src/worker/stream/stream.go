package stream

import (
	"encoding/json"
	"fmt"
	"github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
	"github.com/fallenstedt/twitter-stream/token_generator"
	"github.com/singhkshitij/golang-rest-service-starter/schema"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
)

type StreamSvc struct {
	AccessToken string
	API         *twitterstream.TwitterApi
}

func GenerateStreamToken() (*token_generator.RequestBearerTokenResponse, error) {
	twitterConfig := config.TwitterConf()
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(twitterConfig.APIkey, twitterConfig.APIKeySecret).RequestBearerToken()
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func CreateStream(accessToken string) *twitterstream.TwitterApi {
	return twitterstream.NewTwitterStream(accessToken)
}

func BuildRulesRequest(rules map[string]string) rules.CreateRulesRequest {
	ruleBuilder := twitterstream.NewRuleBuilder()
	for tag, rule := range rules {
		ruleBuilder.AddRule(rule, tag)
	}
	return ruleBuilder.Build()
}

func CreateRule(rules *rules.CreateRulesRequest, s *twitterstream.TwitterApi) (*rules.TwitterRuleResponse, error) {
	res, err := s.Rules.Create(*rules, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SetUnmarshalHook(s *twitterstream.TwitterApi) {
	s.Stream.SetUnmarshalHook(func(bytes []byte) (interface{}, error) {
		dataType := schema.TweetData{}
		if err := json.Unmarshal(bytes, &dataType); err != nil {
			fmt.Printf("failed to unmarshal bytes: %v", err)
		}
		return dataType, nil

		//return string(bytes), nil
	})
}

func GetAllRules(s *twitterstream.TwitterApi) (*rules.TwitterRuleResponse, error) {
	res, err := s.Rules.Get()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteAllRules(ruleIDs []int, s *twitterstream.TwitterApi) error {
	_, err := s.Rules.Delete(rules.NewDeleteRulesRequest(ruleIDs...), false)
	if err != nil {
		return err
	}
	return nil
}
