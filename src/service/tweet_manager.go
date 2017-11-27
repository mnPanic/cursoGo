package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

var tweets []domain.Tweet
var users []domain.User

//GetTweets returns all tweets
func GetTweets() []domain.Tweet {
	return tweets
}

//GetTweet returns the last published Tweet
func GetTweet() domain.Tweet {
	return tweets[len(tweets)-1]
}

//GetTimelineFromUser returns all tweets from one user
func GetTimelineFromUser(user domain.User) (timeline []domain.Tweet) {
	for _, t := range tweets {
		if t.User.Name == user.Name {
			timeline = append(timeline, t)
		}
	}
	return
}

//InitializeService initializes the service
func InitializeService() {
	tweets = []domain.Tweet{}
}

//PublishTweet Publishes a tweet
func PublishTweet(tweetToPublish *domain.Tweet) error {

	if !IsRegistered(tweetToPublish.User) {
		return fmt.Errorf("User is not registered")
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

//Register register a user
func Register(userToRegister domain.User) error {
	if userToRegister.Name == "" {
		return fmt.Errorf("Name is required")
	}

	users = append(users, userToRegister)
	return nil
}

//IsRegistered verify that a user is registered
func IsRegistered(user domain.User) bool {
	for _, u := range users {
		if u.Name == user.Name {
			return true
		}
	}
	return false
}
