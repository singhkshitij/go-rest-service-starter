package v1

import "github.com/singhkshitij/golang-rest-service-starter/schema"

// User contains user information
type User struct {
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
	Age       uint8  `validate:"gte=0,lte=130" json:"age"`
	Email     string `validate:"required,email" json:"email"`
}

type TweetCategoryResponse struct {
	Results     []schema.TweetData `json:"results"`
	TotalTweets int                `json:"total_tweets"`
}
