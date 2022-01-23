package schema

import (
	"encoding/json"
	"time"
)

type TweetData struct {
	Data struct {
		AuthorId       string    `json:"author_id,omitempty"`
		ConversationId string    `json:"conversation_id,omitempty"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
		Entities       struct {
			Hashtags []struct {
				Tag string `json:"tag,omitempty"`
			} `json:"hashtags,omitempty"`
			Mentions []struct {
				Username string `json:"username,omitempty"`
				Id       string `json:"id,omitempty"`
			} `json:"mentions,omitempty"`
			Urls []struct {
				Url         string `json:"url,omitempty"`
				ExpandedUrl string `json:"expanded_url,omitempty"`
				DisplayUrl  string `json:"display_url,omitempty"`
			} `json:"urls,omitempty"`
		} `json:"entities,omitempty"`
		Id            string `json:"id,omitempty"`
		Lang          string `json:"lang,omitempty"`
		PublicMetrics struct {
			RetweetCount int `json:"retweet_count,omitempty"`
			ReplyCount   int `json:"reply_count,omitempty"`
			LikeCount    int `json:"like_count,omitempty"`
			QuoteCount   int `json:"quote_count,omitempty"`
		} `json:"public_metrics,omitempty"`
		ReplySettings string `json:"reply_settings,omitempty"`
		Source        string `json:"source,omitempty"`
		Text          string `json:"text,omitempty"`
	} `json:"data,omitempty"`
	Includes struct {
		Users []struct {
			CreatedAt   time.Time `json:"created_at,omitempty"`
			Description string    `json:"description,omitempty"`
			Entities    struct {
				Url struct {
					Urls []struct {
						Url         string `json:"url,omitempty"`
						ExpandedUrl string `json:"expanded_url,omitempty"`
						DisplayUrl  string `json:"display_url,omitempty"`
					} `json:"urls,omitempty"`
				} `json:"url,omitempty"`
				Description struct {
					Urls []struct {
						Url         string `json:"url,omitempty"`
						ExpandedUrl string `json:"expanded_url,omitempty"`
						DisplayUrl  string `json:"display_url,omitempty"`
					} `json:"urls,omitempty"`
					Mentions []struct {
						Username string `json:"username,omitempty"`
					} `json:"mentions,omitempty"`
					Hashtags []struct {
						Tag string `json:"tag,omitempty"`
					} `json:"hashtags,omitempty"`
				} `json:"description,omitempty"`
			} `json:"entities,omitempty"`
			Id              string `json:"id,omitempty"`
			Location        string `json:"location,omitempty"`
			Name            string `json:"name,omitempty"`
			ProfileImageUrl string `json:"profile_image_url,omitempty"`
			PublicMetrics   struct {
				FollowersCount int `json:"followers_count,omitempty"`
				FollowingCount int `json:"following_count,omitempty"`
				TweetCount     int `json:"tweet_count,omitempty"`
				ListedCount    int `json:"listed_count,omitempty"`
			} `json:"public_metrics,omitempty"`
			Url      string `json:"url,omitempty"`
			Username string `json:"username,omitempty"`
			Verified bool   `json:"verified,omitempty"`
		} `json:"users,omitempty"`
	} `json:"includes,omitempty"`
	MatchingRules []struct {
		Id  string `json:"id,omitempty"`
		Tag string `json:"tag,omitempty"`
	} `json:"matching_rules,omitempty"`
}

func (i TweetData) MarshalBinary() (data []byte, err error) {
	return json.Marshal(i)
}
