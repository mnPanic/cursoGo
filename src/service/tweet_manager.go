package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

var tweet domain.Tweet

//GetTweet returns the tweet
func GetTweet() domain.Tweet {
	return tweet
}

//PublishTweet Publishes a tweet
func PublishTweet(tweetToPublish *domain.Tweet) error {

	if tweetToPublish.User == "" {
		return fmt.Errorf("User is required")
	}

	if tweetToPublish.Text == "" {
		return fmt.Errorf("Text is required")
	}

	if len(tweetToPublish.Text) > 140 {
		return fmt.Errorf("Can't have more than 140 characters")
	}

	tweet = *tweetToPublish
	return nil
}
