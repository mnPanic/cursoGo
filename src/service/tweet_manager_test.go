package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

//UTILITY FUNCTIONS

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

//REGISTERING TEST

func TestCanRegisterUser(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	//Operation
	manager.Register(user)
	//Validation
	if !manager.IsRegistered(user) {
		t.Error("User did not get registered")
	}
}

func TestCantRegisterUserWithInvalidName(t *testing.T) {
	//Initalization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("", "pass")
	//Operation
	err := manager.Register(user)
	//Validation
	validateExpectedError(t, err, "Invalid name")
}

func TestCantRegisterUserWithInvalidPassword(t *testing.T) {
	//Initalization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("name", "")
	//Operation
	err := manager.Register(user)
	//Validation
	validateExpectedError(t, err, "Invalid password")
}

func TestCantRegisterSameUserMoreThanOnce(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	//Operation
	manager.Register(user)
	err := manager.Register(user)
	//Validation
	validateExpectedError(t, err, "The user is already registered")
}

//LOGIN TESTS

func TestCantLoginIfAlreadyLoggedIn(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	//Operation
	err := manager.Login(user)

	//Validation
	validateExpectedError(t, err, "Already logged in")
}
func TestCantLogInWithUnregisteredUser(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")

	//Operation
	err := manager.Login(user)

	//Validation
	validateExpectedError(t, err, "The user is not registered")

}

func TestCanGetLoggedInUser(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	//Operation
	loggedInUser, _ := manager.GetLoggedInUser()
	//Validate
	if user.Name != loggedInUser.Name {
		t.Error("The loggedInUser and the user that logged in do not match")
	}
}

func TestCantGetLoggedInUserIfNoOneLoggedIn(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	//Operation
	_, err := manager.GetLoggedInUser()
	//Validate
	validateExpectedError(t, err, "Not logged in")
}

func TestCantLogInWithIncorrectPassword(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	incorrectUser := domain.NewUser("root", "incorrectPassword")
	manager.Register(user)
	//Operation
	err := manager.Login(incorrectUser)
	//Validation
	validateExpectedError(t, err, "The user is not registered")
}

func TestCanCreateTextTweet(t *testing.T) {}
func TestCantCreateTextTweetWithoutText(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	var text string

	tweet, _ := domain.NewTextTweet(user, text)

	//Operation
	err := manager.PublishTweet(tweet)

	//Validation
	validateExpectedError(t, err, "Text is required")
}
func TestCantCreateTextTweetWithMoreThan140Characters(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."

	//Operation
	_, err := domain.NewTextTweet(user, text)

	//Validation
	validateExpectedError(t, err, "Can't have more than 140 characters")
}
func TestCanCompareTwoTextTweets(t *testing.T) {}

func TestCanCreateImageTweet(t *testing.T)                 {}
func TestCanCompareTwoImageTweets(t *testing.T)            {}
func TestCantCreateImageTweetWithoutImageURL(t *testing.T) {}

func TestCanCreateQuoteTweet(t *testing.T)      {}
func TestCanCompareTwoQuoteTweets(t *testing.T) {}

//PUBLISHING TWEET TESTS
func TestPublishedTweetIsSaved(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "This is my first tweet"
	tweet, _ := domain.NewTextTweet(user, text)
	//Operation
	err := manager.PublishTweet(*tweet)

	if err != nil {
		t.Errorf(err.Error())
	}

	//Validation
	publishedTweet, _ := manager.GetTweet()
	isValidTweet(t, publishedTweet, user, text)
}

func TestMustBeLoggedInToPublishTweet(t *testing.T) {
	//Initalization
	var manager service.TweetManager
	manager.InitializeManager()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	manager.Register(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTextTweet(user, text)
	//Operation
	err := manager.PublishTweet(tweet)
	validateExpectedError(t, err, "You must be logged in to tweet")

}

func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet, _ := domain.NewTextTweet(user, text)
	secondTweet, _ := domain.NewTextTweet(user, secondText)

	//Operation
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)

	//Validation
	publishedTweets, _ := manager.GetTimelineFromUser(user)

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

//TIMELINE TESTS
func TestCanRetrieveTimeline(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("Manuel", "pw")
	manager.Register(user)

	secondUser := domain.NewUser("Gonzalo", "pw")
	manager.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"

	tweet, _ := domain.NewTextTweet(user, text)
	secondTweet, _ := domain.NewTextTweet(user, secondText)
	thirdTweet, _ := domain.NewTextTweet(secondUser, thirdText)

	manager.Login(secondUser)
	manager.PublishTweet(thirdTweet)
	manager.Logout()

	manager.Login(user)
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)

	//Operation
	publishedTweets, _ := manager.GetTimeline()

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
	var manager service.TweetManager
	manager.InitializeManager()

	var tweet *domain.Tweet

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTextTweet(user, text)

	manager.PublishTweet(tweet)
	manager.Logout()

	//Operation
	_, err := manager.GetTimeline()

	//Validation
	validateExpectedError(t, err, "No user logged in")
}

func TestCantRetrieveTimelineOfUnregisteredUser(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")

	//Operation
	_, err := manager.GetTimelineFromUser(user)

	//Validation
	validateExpectedError(t, err, "That user is not registered")
}

//TWEETBYIDTESTS

func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTextTweet(user, text)
	//Operations
	manager.PublishTweet(tweet)

	//Validation
	publishedTweet, err := manager.GetTweetByID(0)
	if err != nil {
		t.Errorf("Did not expect error, but got %s", err.Error())
	}
	isValidTweet(t, *publishedTweet, user, text)
}

func TestCantRetrieveTweetByNonExistentID(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	var tweet *domain.Tweet
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTextTweet(user, text)
	//Operations
	err := manager.PublishTweet(tweet)
	_, err = manager.GetTweetByID(5)

	validateExpectedError(t, err, "A tweet with that ID does not exist")
}

//TWEETEXISTS TESTS
func TestCanCheckIfTweetExists(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "hola soy root")
	manager.PublishTweet(tweet)
	//Operation
	exists := manager.TweetExists(*tweet)
	//Validation
	if !exists {
		t.Error("The tweet should exist")
	}
}

//DELETETWEET TESTS

func TestCanDeleteTweet(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "Tweet 1")
	tweet2, _ := domain.NewTextTweet(user, "Tweet 2")
	manager.PublishTweet(tweet)
	manager.PublishTweet(tweet2)
	//Operation
	manager.DeleteTweetByID(tweet.ID)
	//Validation
	exists := manager.TweetExists(*tweet)
	if exists {
		t.Error("Tweet shouldn't exist")
	}
}

func TestCantDeleteNonExistentTweet(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	//Operation
	user := domain.NewUser("useless", "user")
	manager.Register(user)
	err := manager.DeleteTweetByID(2)
	//Validation
	validateExpectedError(t, err, "Coudln't delete tweet, A tweet with that ID does not exist")
}

func TestCantDeleteATweetThatYouDidntPublish(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user1 := domain.NewUser("root", "root")
	user2 := domain.NewUser("manu", "hunter2")
	manager.Register(user1)
	manager.Register(user2)
	tweet, _ := domain.NewTextTweet(user1, "hola")

	manager.Login(user1)
	manager.PublishTweet(tweet)
	manager.Logout()

	manager.Login(user2)
	//Operation
	err := manager.DeleteTweetByID(tweet.ID)
	//Validation
	validateExpectedError(t, err, "You can't delete a tweet that you didn't publish")
}

func TestCantDeleteATweetIfNotLoggedIn(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user1 := domain.NewUser("root", "root")
	manager.Register(user1)
	tweet, _ := domain.NewTextTweet(user1, "hola")

	manager.Login(user1)
	manager.PublishTweet(tweet)
	manager.Logout()
	//Operation
	err := manager.DeleteTweetByID(tweet.ID)
	//Validation
	validateExpectedError(t, err, "Coudln't delete tweet, Not logged in")
}

//SPECIFIC TWEET TESTS
/*
func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
*/
