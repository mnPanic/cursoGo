package service

import (
	"github.com/cursoGo/src/domain"
)

var tweet domain.Tweet

//GetTweet returns the tweet
func GetTweet() domain.Tweet {
	return tweet
}

//PublishTweet Publishes a tweet
func PublishTweet(tw *domain.Tweet) {
	tweet = *tw
}
