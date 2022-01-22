package worker

import (
	"encoding/json"
	"fmt"
	"github.com/fallenstedt/twitter-stream"
	"github.com/fallenstedt/twitter-stream/rules"
	"github.com/fallenstedt/twitter-stream/token_generator"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
)

type StreamSvc struct {
	AccessToken string
	API         *twitterstream.TwitterApi
}

func (s *StreamSvc) GenerateStreamToken() (*token_generator.RequestBearerTokenResponse, error) {
	twitterConfig := config.TwitterConf()
	tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret(twitterConfig.APIkey, twitterConfig.APIKeySecret).RequestBearerToken()
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func (s *StreamSvc) CreateStream() *twitterstream.TwitterApi {
	return twitterstream.NewTwitterStream(s.AccessToken)
}

func (s *StreamSvc) BuildRulesRequest(rules map[string]string) rules.CreateRulesRequest {
	ruleBuilder := twitterstream.NewRuleBuilder()
	for tag, rule := range rules {
		ruleBuilder.AddRule(rule, tag)
	}
	return ruleBuilder.Build()
}

func (s *StreamSvc) DryRunCreateRule(rules *rules.CreateRulesRequest) (*rules.TwitterRuleResponse, error) {
	res, err := s.API.Rules.Create(*rules, true)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StreamSvc) CreateRule(rules *rules.CreateRulesRequest) (*rules.TwitterRuleResponse, error) {
	res, err := s.API.Rules.Create(*rules, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StreamSvc) SetUnmarshalHook(objectToUnMarshalData interface{})  {
	s.API.Stream.SetUnmarshalHook(func(bytes []byte) (interface{}, error) {

		if err := json.Unmarshal(bytes, &objectToUnMarshalData); err != nil {
			fmt.Printf("failed to unmarshal bytes: %v", err)
		}

		return objectToUnMarshalData, nil
		//return string(bytes), nil
	})
}

func (s *StreamSvc) GetAllRules() (*rules.TwitterRuleResponse, error) {
	res, err := s.API.Rules.Get()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StreamSvc) DryRunDeleteAllRules(ruleIDs []int) (*rules.TwitterRuleResponse, error) {
	res, err := s.API.Rules.Delete(rules.NewDeleteRulesRequest(ruleIDs...), true)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StreamSvc) DeleteAllRules(ruleIDs []int) (*rules.TwitterRuleResponse, error) {
	res, err := s.API.Rules.Delete(rules.NewDeleteRulesRequest(ruleIDs...), false)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *StreamSvc) StartStream() error {
	streamExpansions := twitterstream.NewStreamQueryParamsBuilder().
		AddExpansion("author_id,entities.mentions.username").
		AddTweetField("created_at,author_id,attachments,conversation_id,entities,geo,id,lang,public_metrics,referenced_tweets,reply_settings,source,text,withheld").
		AddUserField("verified,username,url,public_metrics,profile_image_url,name,location,id,entities,description,created_at").
		AddMediaField("preview_image_url,type,url,width,height,alt_text").
		AddPlaceField("country,country_code,full_name,name,place_type").
		Build()

	err := s.API.Stream.StartStream(streamExpansions)
	if err != nil {
		return err
	}
	return nil
}