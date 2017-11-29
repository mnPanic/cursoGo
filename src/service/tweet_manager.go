package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

var userTweets map[domain.User][]domain.Tweet
var loggedInUser domain.User

//InitializeService initializes the service
func InitializeService() {
	userTweets = make(map[domain.User][]domain.Tweet)
	domain.ResetCurrentID()
	Logout()
}

//Register register a user
func Register(userToRegister domain.User) error {
	if userToRegister.Name == "" {
		return fmt.Errorf("Invalid name")
	}
	if userToRegister.Password == "" {
		return fmt.Errorf("Invalid password")
	}

	if IsRegistered(userToRegister) {
		return fmt.Errorf("The user is already registered")
	}

	userTweets[userToRegister] = make([]domain.Tweet, 0)
	return nil
}

//IsRegistered verifies that a user is registered
func IsRegistered(user domain.User) bool {
	_, ok := userTweets[user]
	return ok
}

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

//GetLoggedInUser returns the logged in user
func GetLoggedInUser() (*domain.User, error) {
	if !isLoggedIn() {
		return nil, fmt.Errorf("Not logged in")
	}
	return &loggedInUser, nil
}

//Logout logs the user out
func Logout() error {
	if !isLoggedIn() {
		return fmt.Errorf("Not logged in")
	}
	loggedInUser = domain.User{}
	return nil
}

//isLoggedIn checks if there is a logged in user
func isLoggedIn() bool {
	return loggedInUser.Name != ""
}

//GetTweet returns the last published Tweet
func GetTweet() (domain.Tweet, error) {
	tw, err := GetTweetByID(domain.GetCurrentID())
	return *tw, err
}

//GetTweetByID returns the tweet that has that ID
func GetTweetByID(id int) (*domain.Tweet, error) {
	for _, tweets := range userTweets {
		for _, tweet := range tweets {
			if tweet.ID == id {
				return &tweet, nil
			}
		}
	}
	return nil, fmt.Errorf("A tweet with that ID does not exist")
}

//GetTimelineFromUser returns all tweets from one user
func GetTimelineFromUser(user domain.User) ([]domain.Tweet, error) {
	if !IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline := userTweets[user]
	return timeline, nil
}

//GetTimeline returns the loggedInUser's timeline
func GetTimeline() ([]domain.Tweet, error) {
	if !isLoggedIn() {
		return nil, fmt.Errorf("No user logged in")
	}
	return GetTimelineFromUser(loggedInUser)
}

//PublishTweet Publishes a tweet
func PublishTweet(tweetToPublish *domain.Tweet) error {
	if !loggedInUser.Equals(tweetToPublish.User) {
		return fmt.Errorf("You must be logged in to tweet")
	}

	if tweetToPublish.Text == "" {
		return fmt.Errorf("Text is required")
	}

	userTweets[tweetToPublish.User] = append(userTweets[tweetToPublish.User], *tweetToPublish)
	return nil
}

//DeleteTweetByID deletes a tweet by its ID
func DeleteTweetByID(id int) error {
	tweet, err := GetTweetByID(id)
	if err != nil {
		return fmt.Errorf("Coudln't delete tweet, %s", err.Error())
	}
	user, err := GetLoggedInUser()
	if err != nil {
		return fmt.Errorf("Coudln't delete tweet, %s", err.Error())
	}

	if !tweet.User.Equals(*user) {
		return fmt.Errorf("You can't delete a tweet that you didn't publish")
	}
	return deleteTweet(*tweet)
}

//DeleteTweet deletes a tweet
func deleteTweet(tweet domain.Tweet) error {
	tweets := userTweets[tweet.User]
	tweets = deleteElementFromTweets(tweets, tweet)
	userTweets[tweet.User] = tweets
	return nil
}

func deleteElementFromTweets(slice []domain.Tweet, element domain.Tweet) []domain.Tweet {
	var newTweets []domain.Tweet
	for _, tweet := range slice {
		if !tweet.Equals(element) {
			newTweets = append(newTweets, tweet)
		}
	}
	return newTweets
}

//TweetExists returns if a given tweet exists
func TweetExists(tweet domain.Tweet) bool {
	for _, tweets := range userTweets {
		for _, tw := range tweets {
			if tweet.Equals(tw) {
				return true
			}
		}
	}
	return false
}
