package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
	"github.com/cursoGo/src/utility"
)

//UTILITY FUNCTIONS

func isValidTweet(t *testing.T, publishedTweet domain.Tweeter, user domain.User, text string) bool {
	if !publishedTweet.GetUser().Equals(user) && publishedTweet.GetText() != text {
		t.Errorf("Expected other tweet")
		return false
	}

	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}
	return true
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
	utility.ValidateExpectedError(t, err, "Invalid name")
}

func TestCantRegisterUserWithInvalidPassword(t *testing.T) {
	//Initalization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("name", "")
	//Operation
	err := manager.Register(user)
	//Validation
	utility.ValidateExpectedError(t, err, "Invalid password")
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
	utility.ValidateExpectedError(t, err, "The user is already registered")
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
	utility.ValidateExpectedError(t, err, "Already logged in")
}
func TestCantLogInWithUnregisteredUser(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")

	//Operation
	err := manager.Login(user)

	//Validation
	utility.ValidateExpectedError(t, err, "The user is not registered")

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
	utility.ValidateExpectedError(t, err, "Not logged in")
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
	utility.ValidateExpectedError(t, err, "The user is not registered")
}

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
	err := manager.PublishTweet(tweet)

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

	user := domain.NewUser("root", "root")
	manager.Register(user)

	text := "This is my first tweet"
	tweet, _ := domain.NewTextTweet(user, text)
	//Operation
	err := manager.PublishTweet(tweet)
	utility.ValidateExpectedError(t, err, "You must be logged in to tweet")

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
	publishedTweets, _ := manager.GetTweetsFromUser(user)

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
		if tweet.GetUser().Name != user.Name {
			t.Errorf("Expected user is %s but was %s", user.Name, tweet.GetUser().Name)
		}
	}
}

func TestCantRetrieveTimelineWithoutLoggingIn(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"
	tweet, _ := domain.NewTextTweet(user, text)

	manager.PublishTweet(tweet)
	manager.Logout()

	//Operation
	_, err := manager.GetTimeline()

	//Validation
	utility.ValidateExpectedError(t, err, "No user logged in")
}

func TestCantGetTweetsOfUnregisteredUser(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	//Operation
	_, err := manager.GetTweetsFromUser(user)

	//Validation
	utility.ValidateExpectedError(t, err, "That user is not registered")
}

func TestCantRetrieveTimelineOfUnregisteredUser(t *testing.T) {

	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")

	//Operation
	_, err := manager.GetTimelineFromUser(user)

	//Validation
	utility.ValidateExpectedError(t, err, "That user is not registered")
}

func TestCanRetrieveTimelineWithFollowedUsersTweets(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("Manuel", "pw")
	manager.Register(user)

	secondUser := domain.NewUser("Gonzalo", "pw")
	manager.Register(secondUser)

	otherUser := domain.NewUser("name", "pw")
	manager.Register(otherUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"
	fourthText := "This should not be picked up"

	tweet, _ := domain.NewTextTweet(user, text)
	secondTweet, _ := domain.NewTextTweet(user, secondText)
	thirdTweet, _ := domain.NewTextTweet(secondUser, thirdText)
	fourthTweet, _ := domain.NewTextTweet(otherUser, fourthText)

	manager.Login(otherUser)
	manager.PublishTweet(fourthTweet)
	manager.Logout()

	manager.Login(secondUser)
	manager.PublishTweet(thirdTweet)
	manager.Logout()

	manager.Login(user)
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)
	manager.FollowUser(secondUser.Name)

	//Operation
	publishedTweets, _ := manager.GetTimeline()

	//Validation
	if len(publishedTweets) != 3 {
		t.Errorf("Expected size is 3 but was %d", len(publishedTweets))
		return
	}

	for _, tweet := range publishedTweets {
		if !(tweet.GetUser().Equals(user) || tweet.GetUser().Equals(secondUser)) {
			t.Errorf("Got unexpected user")
		}
	}
}

//TWEETBYIDTESTS

func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ := domain.NewTextTweet(user, text)
	//Operations
	manager.PublishTweet(tweet)

	//Validation
	publishedTweet, err := manager.GetTweetByID(0)
	if err != nil {
		t.Errorf("Did not expect error, but got %s", err.Error())
	}
	isValidTweet(t, publishedTweet, user, text)
}

func TestCantRetrieveTweetByNonExistentID(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()

	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ := domain.NewTextTweet(user, text)
	//Operations
	err := manager.PublishTweet(tweet)
	_, err = manager.GetTweetByID(5)

	utility.ValidateExpectedError(t, err, "A tweet with that ID does not exist")
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
	exists := manager.TweetExists(tweet)
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
	manager.DeleteTweetByID(tweet.GetID())
	//Validation
	exists := manager.TweetExists(tweet)
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
	utility.ValidateExpectedError(t, err, "Coudln't delete tweet, A tweet with that ID does not exist")
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
	err := manager.DeleteTweetByID(tweet.GetID())
	//Validation
	utility.ValidateExpectedError(t, err, "You can't delete a tweet that you didn't publish")
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
	err := manager.DeleteTweetByID(tweet.GetID())
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't delete tweet, Not logged in")
}

//Editing text tests
func TestCanEditTweetText(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	newText := "modified sample"
	//Operation
	manager.EditTweetTextByID(tweet.GetID(), newText)
	//Validation
	retrievedTweet, _ := manager.GetTweet()
	if retrievedTweet.GetText() != newText {
		t.Error("Edited tweet text does not match")
	}
}
func TestCantEditNonExistentTweet(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	newText := "modified sample"
	//Operation
	err := manager.EditTweetTextByID(4, newText)
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't edit tweet, A tweet with that ID does not exist")
}
func TestCantEditTweetIfNotLoggedIn(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	newText := "modified sample"
	manager.Logout()
	//Operation
	err := manager.EditTweetTextByID(tweet.GetID(), newText)
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't edit tweet, Not logged in")
}
func TestCantEditTweetThatYouDidNotPublish(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	otherUser := domain.NewUser("manu", "hunter2")
	manager.Register(otherUser)
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	newText := "modified sample"
	manager.Logout()
	manager.Login(otherUser)
	//Operation
	err := manager.EditTweetTextByID(tweet.GetID(), newText)
	//Validation
	utility.ValidateExpectedError(t, err, "You can't edit a tweet that you didn't publish")
}
func TestCantEditTweetWithEmptyText(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	invalidText := ""
	//Operation
	err := manager.EditTweetTextByID(tweet.GetID(), invalidText)
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't edit tweet, Can't have no text")
}

func TestCantEditTextTweetWithLongText(t *testing.T) {
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("root", "root")
	manager.Register(user)
	manager.Login(user)
	text := "sample"
	tweet, _ := domain.NewTextTweet(user, text)
	manager.PublishTweet(tweet)
	invalidText := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."
	//Operation
	err := manager.EditTweetTextByID(tweet.GetID(), invalidText)
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't edit tweet, Can't have more than 140 characters")
}

func TestCanFollowUser(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("manu", "hunter2")
	secondUser := domain.NewUser("gonza", "hunter3")

	manager.Register(user)
	manager.Register(secondUser)
	manager.Login(user)
	//Operation
	err := manager.FollowUser(secondUser.Name)
	if err != nil {
		t.Errorf("Unexpected error, %s", err.Error())
		return
	}
	//Validation
	u, _ := manager.GetLoggedInUser()
	following := u.IsFollowing(secondUser)
	if !following {
		t.Error("User not followed correctly")
	}
}
func TestCantFollowNonexistentUser(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("manu", "hunter2")
	secondUser := domain.NewUser("gonza", "hunter3")

	manager.Register(user)
	manager.Login(user)
	//Operation
	err := manager.FollowUser(secondUser.Name)
	//Validation
	utility.ValidateExpectedError(t, err, "Couldn't follow user, User not registered")
}
func TestCantFollowIfNotLoggedIn(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	secondUser := domain.NewUser("gonza", "hunter3")
	//Operation
	err := manager.FollowUser(secondUser.Name)
	//Validation
	utility.ValidateExpectedError(t, err, "Coudln't follow user, Not logged in")
}

func TestCantFollowYourself(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("manu", "hunter2")

	manager.Register(user)
	manager.Login(user)
	//Operation
	err := manager.FollowUser(user.Name)
	//Validation
	utility.ValidateExpectedError(t, err, "Can't follow yourself")
}

func TestCantFollowSameUserTwice(t *testing.T) {
	//Initialization
	var manager service.TweetManager
	manager.InitializeManager()
	user := domain.NewUser("manu", "hunter2")
	secondUser := domain.NewUser("gonza", "hunter3")

	manager.Register(user)
	manager.Register(secondUser)
	manager.Login(user)
	//Operation
	manager.FollowUser(secondUser.Name)
	err := manager.FollowUser(secondUser.Name)
	//Validation
	utility.ValidateExpectedError(t, err, "Can't follow same user twice")
}
