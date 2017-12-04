package domain_test

import (
	"testing"

	"github.com/cursoGo/src/utility"

	"github.com/cursoGo/src/domain"
)

//TextTweet Tests

func TestCanSetTweetText(t *testing.T) {
	//Initialization
	user := domain.NewUser("root", "root")
	text := "first"
	secondText := "second"
	tweet, _ := domain.NewTextTweet(user, text)
	//Operation
	tweet.SetText(secondText)
	//Validation
	if tweet.GetText() != secondText {
		t.Error("Tweet text did not get edited properly")
	}
}
func TestCantCreateTextTweetWithoutText(t *testing.T) {

	//Initialization
	user := domain.NewUser("root", "root")
	var text string
	//Operation
	_, err := domain.NewTextTweet(user, text)

	//Validation
	utility.ValidateExpectedError(t, err, "Tweet can't have no text")
}
func TestCantCreateTextTweetWithMoreThan140Characters(t *testing.T) {
	//Initialization
	user := domain.NewUser("root", "root")
	text := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."

	//Operation
	_, err := domain.NewTextTweet(user, text)

	//Validation
	utility.ValidateExpectedError(t, err, "Can't have more than 140 characters")
}
func TestCanCompareTwoTextTweets(t *testing.T) {
	//Initialization

	user := domain.NewUser("root", "root")

	firstText := "first text"
	secondText := "second text"
	firstTweet, _ := domain.NewTextTweet(user, firstText)
	secondTweet, _ := domain.NewTextTweet(user, secondText)

	//Operation
	firstResult := firstTweet.Equals(firstTweet)
	secondResult := secondTweet.Equals(firstTweet)

	//Validation
	if !firstResult {
		t.Error("First result should be true")
	}
	if secondResult {
		t.Error("Second result should be false")
	}
}

func TestCanCompareTwoImageTweets(t *testing.T) {
	//Initialization
	user := domain.NewUser("root", "root")

	text := "unimportant text"
	firstURL := "https://google.com.ar"
	secondURL := "https://facebook.com.ar"
	firstTweet, _ := domain.NewImageTweet(user, text, firstURL)
	secondTweet, _ := domain.NewImageTweet(user, text, secondURL)

	//Operation
	firstResult := firstTweet.Equals(*firstTweet)
	secondResult := secondTweet.Equals(*firstTweet)

	//Validation
	if !firstResult {
		t.Error("First result should be true")
	}
	if secondResult {
		t.Error("Second result should be false")
	}
}
func TestCantCreateImageTweetWithoutImageURL(t *testing.T) {
	//Initialization

	user := domain.NewUser("root", "root")

	text := "unimportant text"
	var url string
	//Operation
	_, err := domain.NewImageTweet(user, text, url)
	//Validation
	utility.ValidateExpectedError(t, err, "Cant create an image tweet without an URL")
}

func TestCanCompareTwoQuoteTweets(t *testing.T) {
	//Initialization
	user := domain.NewUser("root", "root")

	text := "unimportant text"
	secondText := "boring text"
	firstTweetToBeQuoted, _ := domain.NewTextTweet(user, text)
	secondTweetToBeQuoted, _ := domain.NewTextTweet(user, secondText)

	firstQuoteTweet, _ := domain.NewQuoteTweet(user, text, firstTweetToBeQuoted)
	secondQuoteTweet, _ := domain.NewQuoteTweet(user, text, secondTweetToBeQuoted)

	//Operation
	firstResult := firstQuoteTweet.Equals(*firstQuoteTweet)
	secondResult := secondQuoteTweet.Equals(*firstQuoteTweet)

	//Validation
	if !firstResult {
		t.Error("First result should be true")
	}
	if secondResult {
		t.Error("Second result should be false")
	}
}
