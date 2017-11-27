package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

var tweets []domain.Tweet

//GetTweets returns all tweets
func GetTweets() []domain.Tweet {
	return tweets
}

//GetTweet returns the last published Tweet
func GetTweet() domain.Tweet {
	return tweets[len(tweets)-1]
}

//InitializeService initializes the service
func InitializeService() {
	tweets = []domain.Tweet{}
}

//PublishTweet Publishes a tweet
func PublishTweet(tweetToPublish *domain.Tweet) error {

	if tweetToPublish.User.Name == "" {
		return fmt.Errorf("User is required")
	}

	if tweetToPublish.Text == "" {
		return fmt.Errorf("Text is required")
	}

	if len(tweetToPublish.Text) > 140 {
		return fmt.Errorf("Can't have more than 140 characters")
	}

	tweets = append(tweets, *tweetToPublish)
	return nil
}
