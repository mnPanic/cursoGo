package service_test

import (
	"testing"

	"github.com/cursoGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet string = "This is my frist tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
