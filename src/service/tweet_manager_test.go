package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func isValidTweet(t *testing.T, publishedTweet domain.Tweet, user domain.User, text string) bool {
	if publishedTweet.User.Name != user.Name && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user.Name, text, publishedTweet.User.Name, publishedTweet.Text)
		return false
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}
	return true
}

func TestPublishedTweetIsSaved(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	service.Register(user)

	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	//Operation
	err := service.PublishTweet(tweet)

	if err != nil {
		t.Errorf(err.Error())
	}

	//Validation
	publishedTweet := service.GetTweet()
	isValidTweet(t, publishedTweet, user, text)
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := domain.NewUser("Gonzalo")
	service.Register(user)
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	err := service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "Text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := domain.NewUser("Manuel")
	service.Register(user)
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	//Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	//Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}
	isValidTweet(t, secondPublishedTweet, user, secondText)
}

func TestCanRegisterUser(t *testing.T) {

	//Initialization
	service.InitializeService()
	user := domain.NewUser("Gonza")
	//Operation
	service.Register(user)
	//Validation
	if !service.IsRegistered(user) {
		t.Error("User did not get registered")
	}
}

func TestCantPublishTweetWithUnregisteredUser(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	err := service.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "User is not registered" {
		t.Error("Expected error is 'User not registered'")
	}
}
