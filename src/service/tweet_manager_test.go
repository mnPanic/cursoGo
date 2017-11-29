package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func isValidTweet(t *testing.T, publishedTweet domain.Tweet, user domain.User, text string) bool {
	if !publishedTweet.User.Equals(user) && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut was %s: %s",
			user.Name, text, publishedTweet.User.Name, publishedTweet.Text)
		return false
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}
	return true
}

func validateExpectedError(t *testing.T, err error, expectedError string) {
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error is '%s', but was %s", expectedError, err.Error())
		return
	}
}

func TestCanRegisterUser(t *testing.T) {

	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	//Operation
	service.Register(user)
	//Validation
	if !service.IsRegistered(user) {
		t.Error("User did not get registered")
	}
}

func TestCantRegisterUserWithInvalidName(t *testing.T) {
	//Initalization
	service.InitializeService()
	user := domain.NewUser("", "pass")
	//Operation
	err := service.Register(user)
	//Validation
	validateExpectedError(t, err, "Invalid name")
}

func TestCantRegisterUserWithInvalidPassword(t *testing.T) {
	//Initalization
	service.InitializeService()
	user := domain.NewUser("name", "")
	//Operation
	err := service.Register(user)
	//Validation
	validateExpectedError(t, err, "Invalid password")
}

func TestCantRegisterSameUserMoreThanOnce(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	//Operation
	service.Register(user)
	err := service.Register(user)
	//Validation
	validateExpectedError(t, err, "The user is already registered")
}

func TestCantLoginIfAlreadyLoggedIn(t *testing.T) {
	//Initialization
	service.InitializeService()

	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)

	//Operation
	err := service.Login(user)

	//Validation
	validateExpectedError(t, err, "Already logged in")
}
func TestCantLogInWithUnregisteredUser(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")

	//Operation
	err := service.Login(user)

	//Validation
	validateExpectedError(t, err, "The user is not registered")

}

func TestCanGetLoggedInUser(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)

	//Operation
	loggedInUser, _ := service.GetLoggedInUser()
	//Validate
	if user.Name != loggedInUser.Name {
		t.Error("The loggedInUser and the user that logged in do not match")
	}
}

func TestCantGetLoggedInUserIfNoOneLoggedIn(t *testing.T) {
	//Initialization
	service.InitializeService()
	//Operation
	_, err := service.GetLoggedInUser()
	//Validate
	validateExpectedError(t, err, "Not logged in")
}

func TestCantLogInWithIncorrectPassword(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	incorrectUser := domain.NewUser("root", "incorrectPassword")
	service.Register(user)
	//Operation
	err := service.Login(incorrectUser)
	//Validation
	validateExpectedError(t, err, "The user is not registered")
}

func TestPublishedTweetIsSaved(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	text := "This is my first tweet"
	tweet, _ = domain.NewTweet(user, text)
	//Operation
	err := service.PublishTweet(tweet)

	if err != nil {
		t.Errorf(err.Error())
	}

	//Validation
	publishedTweet, _ := service.GetTweet()
	isValidTweet(t, publishedTweet, user, text)
}

func TestMustBeLoggedInToPublishTweet(t *testing.T) {
	//Initalization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	service.Register(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTweet(user, text)
	//Operation
	err := service.PublishTweet(tweet)
	validateExpectedError(t, err, "You must be logged in to tweet")

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	var text string

	tweet, _ = domain.NewTweet(user, text)

	//Operation
	err := service.PublishTweet(tweet)

	//Validation
	validateExpectedError(t, err, "Text is required")
}

func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet, _ = domain.NewTweet(user, text)
	secondTweet, _ = domain.NewTweet(user, secondText)

	//Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	//Validation
	publishedTweets, _ := service.GetTimelineFromUser(user)

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

func TestCanRetrieveTimeline(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := domain.NewUser("Manuel", "pw")
	service.Register(user)

	secondUser := domain.NewUser("Gonzalo", "pw")
	service.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"

	tweet, _ = domain.NewTweet(user, text)
	secondTweet, _ = domain.NewTweet(user, secondText)
	thirdTweet, _ = domain.NewTweet(secondUser, thirdText)

	service.Login(secondUser)
	service.PublishTweet(thirdTweet)
	service.Logout()

	service.Login(user)
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	//Operation
	publishedTweets, _ := service.GetTimeline()

	//Validation
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

func TestCantRetrieveTimelineWithoutLoggingIn(t *testing.T) {

	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet

	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTweet(user, text)

	service.PublishTweet(tweet)
	service.Logout()

	//Operation
	_, err := service.GetTimeline()

	//Validation
	validateExpectedError(t, err, "No user logged in")
}

func TestCantRetrieveTimelineOfUnregisteredUser(t *testing.T) {

	//Initialization
	service.InitializeService()

	user := domain.NewUser("root", "root")

	//Operation
	_, err := service.GetTimelineFromUser(user)

	//Validation
	validateExpectedError(t, err, "That user is not registered")
}

func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTweet(user, text)
	//Operations
	service.PublishTweet(tweet)

	//Validation
	publishedTweet, err := service.GetTweetByID(0)
	if err != nil {
		t.Errorf("Did not expect error, but got %s", err.Error())
	}
	isValidTweet(t, *publishedTweet, user, text)
}

func TestCantRetrieveTweetByNonExistentID(t *testing.T) {
	//Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTweet(user, text)
	//Operations
	err := service.PublishTweet(tweet)
	_, err = service.GetTweetByID(5)

	validateExpectedError(t, err, "A tweet with that ID does not exist")
}

func TestCantCreateTweetWithMoreThan140Characters(t *testing.T) {
	//Initialization
	service.InitializeService()

	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	text := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."

	//Operation
	_, err := domain.NewTweet(user, text)

	//Validation
	validateExpectedError(t, err, "Can't have more than 140 characters")
}

func TestCanCheckIfTweetExists(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	tweet, _ := domain.NewTweet(user, "hola soy root")
	service.PublishTweet(tweet)
	//Operation
	exists := service.TweetExists(*tweet)
	//Validation
	if !exists {
		t.Error("The tweet should exist")
	}
}

func TestCanDeleteTweet(t *testing.T) {
	//Initialization
	service.InitializeService()
	user := domain.NewUser("root", "root")
	service.Register(user)
	service.Login(user)
	tweet, _ := domain.NewTweet(user, "Tweet 1")
	tweet2, _ := domain.NewTweet(user, "Tweet 2")
	service.PublishTweet(tweet)
	service.PublishTweet(tweet2)
	//Operation
	service.DeleteTweetByID(tweet.ID)
	//Validation
	exists := service.TweetExists(*tweet)
	if exists {
		t.Error("Tweet shouldn't exist")
	}
}

func TestCantDeleteNonExistentTweet(t *testing.T) {
	//Initialization
	service.InitializeService()
	//Operation
	user := domain.NewUser("useless", "user")
	service.Register(user)
	err := service.DeleteTweetByID(2)
	//Validation
	validateExpectedError(t, err, "Coudln't delete tweet, A tweet with that ID does not exist")
}

func TestCantDeleteATweetThatYouDidntPublish(t *testing.T) {
	//Initialization
	service.InitializeService()
	user1 := domain.NewUser("root", "root")
	user2 := domain.NewUser("manu", "hunter2")
	service.Register(user1)
	service.Register(user2)
	tweet, _ := domain.NewTweet(user1, "hola")

	service.Login(user1)
	service.PublishTweet(tweet)
	service.Logout()

	service.Login(user2)
	//Operation
	err := service.DeleteTweetByID(tweet.ID)
	//Validation
	validateExpectedError(t, err, "You can't delete a tweet that you didn't publish")
}

func TestCantDeleteATweetIfNotLoggedIn(t *testing.T) {
	//Initialization
	service.InitializeService()
	user1 := domain.NewUser("root", "root")
	service.Register(user1)
	tweet, _ := domain.NewTweet(user1, "hola")

	service.Login(user1)
	service.PublishTweet(tweet)
	service.Logout()
	//Operation
	err := service.DeleteTweetByID(tweet.ID)
	//Validation
	validateExpectedError(t, err, "Coudln't delete tweet, Not logged in")
}
