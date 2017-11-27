package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet
	user := "root"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet

	user := "Gonzalo"
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
		t.Error("EXpected error is text is required")
	}
}
