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

func TestCanRetrieveTimeline(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := domain.NewUser("Manuel")
	service.Register(user)
	secondUser := domain.NewUser("Gonzalo")
	service.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(secondUser, thirdText)

	//Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	//Validation
	publishedTweets := service.GetTimelineFromUser(user)

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	for _, tweet := range publishedTweets {
		if tweet.User.Name != user.Name {
			t.Errorf("Expected user is %s but was %s", user.Name, tweet.User.Name)
		}
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	service.Register(user)

	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	//Operations
	err := service.PublishTweet(tweet)

	if err != nil {
		t.Errorf(err.Error())
	}

	//Validation
	publishedTweet, err := service.GetTweetByID(0)
	isValidTweet(t, *publishedTweet, user, text)

	_, err = service.GetTweetByID(5)

	if err == nil {
		t.Errorf("Expected error")
		return
	}

	if err.Error() != "A tweet with that ID does not exist" {
		t.Errorf("The expected error is 'A tweet with that ID does not exist', but was %s ",
			err.Error())
	}
}
