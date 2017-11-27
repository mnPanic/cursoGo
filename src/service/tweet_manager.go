package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

var tweets []domain.Tweet
var users []domain.User
var loggedInUser domain.User

//Login logs the user in
func Login(user domain.User) error {
	if isLoggedIn() {
		return fmt.Errorf("Already logged in")
	}
	if !IsRegistered(user) {
		return fmt.Errorf("The user is not registered")
	}

	loggedInUser = user
	return nil
}

//Logout logs the user out
func Logout() {
	loggedInUser = domain.User{Name: ""}
}

//isLoggedIn checks if there is a logged in user
func isLoggedIn() bool {
	return loggedInUser.Name != ""
}

//GetTweets returns all tweets
func GetTweets() []domain.Tweet {
	return tweets
}

//GetTweet returns the last published Tweet
func GetTweet() domain.Tweet {
	return tweets[len(tweets)-1]
}

//GetTweetByID returns the tweet that has that ID
func GetTweetByID(id int) (*domain.Tweet, error) {
	for _, tweet := range tweets {
		if tweet.ID == id {
			return &tweet, nil
		}
	}
	return nil, fmt.Errorf("A tweet with that ID does not exist")
}

//GetTimelineFromUser returns all tweets from one user
func getTimelineFromUser(user domain.User) (timeline []domain.Tweet) {
	for _, t := range tweets {
		if t.User.Name == user.Name {
			timeline = append(timeline, t)
		}
	}
	return
}

//GetTimeline returns the loggedInUser's timeline
func GetTimeline() ([]domain.Tweet, error) {
	if !isLoggedIn() {
		return nil, fmt.Errorf("No user logged in")
	}
	return getTimelineFromUser(loggedInUser), nil
}

//InitializeService initializes the service
func InitializeService() {
	tweets = []domain.Tweet{}
	users = []domain.User{}
	Logout()
}

//PublishTweet Publishes a tweet
func PublishTweet(tweetToPublish *domain.Tweet) error {

	if !IsRegistered(tweetToPublish.User) {
		return fmt.Errorf("User is not registered")
	}

	if loggedInUser.Name != tweetToPublish.User.Name {
		return fmt.Errorf("You must be logged in to tweet")
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

	if IsRegistered(userToRegister) {
		return fmt.Errorf("The user is already registered")
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
