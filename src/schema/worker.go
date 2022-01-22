package schema

import "time"

type StreamTweet struct {
	Data struct {
		AuthorId       string    `json:"author_id"`
		ConversationId string    `json:"conversation_id"`
		CreatedAt      time.Time `json:"created_at"`
		Entities       struct {
			Annotations []struct {
				Probability    float64 `json:"probability"`
				Type           string  `json:"type"`
				NormalizedText string  `json:"normalized_text"`
			} `json:"annotations"`
			Hashtags []struct {
				Tag   string `json:"tag"`
			} `json:"hashtags"`
			Urls []struct {
				Url         string `json:"url"`
				ExpandedUrl string `json:"expanded_url"`
				DisplayUrl  string `json:"display_url"`
				Status      int    `json:"status,omitempty"`
				Title       string `json:"title,omitempty"`
				Description string `json:"description,omitempty"`
				UnwoundUrl  string `json:"unwound_url,omitempty"`
			} `json:"urls"`
		} `json:"entities"`
		Id            string `json:"id"`
		Lang          string `json:"lang"`
		PublicMetrics struct {
			RetweetCount int `json:"retweet_count"`
			ReplyCount   int `json:"reply_count"`
			LikeCount    int `json:"like_count"`
			QuoteCount   int `json:"quote_count"`
		} `json:"public_metrics"`
		ReplySettings string `json:"reply_settings"`
		Source        string `json:"source"`
		Text          string `json:"text"`
	} `json:"data"`
	Includes struct {
		Media []struct {
			Height        int    `json:"height"`
			MediaKey      string `json:"media_key"`
			PublicMetrics struct {
			} `json:"public_metrics"`
			Type  string `json:"type"`
			Url   string `json:"url"`
			Width int    `json:"width"`
		} `json:"media"`
		Users []struct {
			CreatedAt   time.Time `json:"created_at"`
			Description string    `json:"description"`
			Entities    struct {
				Url struct {
					Urls []struct {
						Url         string `json:"url"`
						ExpandedUrl string `json:"expanded_url"`
						DisplayUrl  string `json:"display_url"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Cashtags []struct {
						Tag   string `json:"tag"`
					} `json:"cashtags"`
				} `json:"description"`
			} `json:"entities"`
			Id              string `json:"id"`
			Location        string `json:"location"`
			Name            string `json:"name"`
			PinnedTweetId   string `json:"pinned_tweet_id"`
			ProfileImageUrl string `json:"profile_image_url"`
			Protected       bool   `json:"protected"`
			PublicMetrics   struct {
				FollowersCount int `json:"followers_count"`
				FollowingCount int `json:"following_count"`
				TweetCount     int `json:"tweet_count"`
				ListedCount    int `json:"listed_count"`
			} `json:"public_metrics"`
			Url      string `json:"url"`
			Username string `json:"username"`
			Verified bool   `json:"verified"`
		} `json:"users"`
	} `json:"includes"`
	MatchingRules []struct {
		Id  string `json:"id"`
		Tag string `json:"tag"`
	} `json:"matching_rules"`
}
