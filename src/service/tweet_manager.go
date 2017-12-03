package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

//TweetManager is a tweet manager
type TweetManager struct {
	userTweets   map[domain.User][]domain.Tweeter
	loggedInUser domain.User
}

//InitializeManager initializes the manager
func (m *TweetManager) InitializeManager() {
	m.userTweets = make(map[domain.User][]domain.Tweeter)
	domain.ResetCurrentID()
	m.Logout()
}

//Register register a user
func (m *TweetManager) Register(userToRegister domain.User) error {
	if userToRegister.Name == "" {
		return fmt.Errorf("Invalid name")
	}
	if userToRegister.Password == "" {
		return fmt.Errorf("Invalid password")
	}

	if m.IsRegistered(userToRegister) {
		return fmt.Errorf("The user is already registered")
	}

	m.userTweets[userToRegister] = make([]domain.Tweeter, 0)
	return nil
}

//IsRegistered verifies that a user is registered
func (m *TweetManager) IsRegistered(user domain.User) bool {
	_, ok := m.userTweets[user]
	return ok
}

//Login logs the user in
func (m *TweetManager) Login(user domain.User) error {
	if m.isLoggedIn() {
		return fmt.Errorf("Already logged in")
	}
	if !m.IsRegistered(user) {
		return fmt.Errorf("The user is not registered")
	}

	m.loggedInUser = user
	return nil
}

//GetLoggedInUser returns the logged in user
func (m *TweetManager) GetLoggedInUser() (*domain.User, error) {
	if !m.isLoggedIn() {
		return nil, fmt.Errorf("Not logged in")
	}
	return &m.loggedInUser, nil
}

//Logout logs the user out
func (m *TweetManager) Logout() error {
	if !m.isLoggedIn() {
		return fmt.Errorf("Not logged in")
	}
	m.loggedInUser = domain.User{}
	return nil
}

//isLoggedIn checks if there is a logged in user
func (m *TweetManager) isLoggedIn() bool {
	return m.loggedInUser.Name != ""
}

//GetTweet returns the last published Tweet
func (m *TweetManager) GetTweet() (domain.Tweeter, error) {
	tw, err := m.GetTweetByID(domain.GetCurrentID())
	return tw, err
}

//GetTweetByID returns the tweet that has that ID
func (m *TweetManager) GetTweetByID(id int) (domain.Tweeter, error) {
	for _, tweets := range m.userTweets {
		for _, tweet := range tweets {
			if tweet.GetID() == id {
				return tweet, nil
			}
		}
	}
	return nil, fmt.Errorf("A tweet with that ID does not exist")
}

//GetTimelineFromUser returns all tweets from one user
func (m *TweetManager) GetTimelineFromUser(user domain.User) ([]domain.Tweeter, error) {
	if !m.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline := m.userTweets[user]
	return timeline, nil
}

//GetTimeline returns the loggedInUser's timeline
func (m *TweetManager) GetTimeline() ([]domain.Tweeter, error) {
	if !m.isLoggedIn() {
		return nil, fmt.Errorf("No user logged in")
	}
	return m.GetTimelineFromUser(m.loggedInUser)
}

//PublishTweet Publishes a tweet
func (m *TweetManager) PublishTweet(tweetToPublish domain.Tweeter) error {
	if !m.loggedInUser.Equals(tweetToPublish.GetUser()) {
		return fmt.Errorf("You must be logged in to tweet")
	}
	m.userTweets[tweetToPublish.GetUser()] = append(m.userTweets[tweetToPublish.GetUser()], tweetToPublish)
	return nil
}

//DeleteTweetByID deletes a tweet by its ID
func (m *TweetManager) DeleteTweetByID(id int) error {
	tweet, err := m.GetTweetByID(id)
	if err != nil {
		return fmt.Errorf("Coudln't delete tweet, %s", err.Error())
	}
	user, err := m.GetLoggedInUser()
	if err != nil {
		return fmt.Errorf("Coudln't delete tweet, %s", err.Error())
	}

	if !tweet.GetUser().Equals(*user) {
		return fmt.Errorf("You can't delete a tweet that you didn't publish")
	}
	return m.deleteTweet(tweet)
}

//DeleteTweet deletes a tweet
func (m *TweetManager) deleteTweet(tweet domain.Tweeter) error {
	tweets := m.userTweets[tweet.GetUser()]
	tweets = m.deleteElementFromTweets(tweets, tweet)
	m.userTweets[tweet.GetUser()] = tweets
	return nil
}

func (m *TweetManager) deleteElementFromTweets(slice []domain.Tweeter, element domain.Tweeter) []domain.Tweeter {
	var newTweets []domain.Tweeter
	for _, tweet := range slice {
		if !tweet.Equals(element) {
			newTweets = append(newTweets, tweet)
		}
	}
	return newTweets
}

func (m *TweetManager) tweetAppearsByCriteria(tweet domain.Tweeter, criteria func(domain.Tweeter, domain.Tweeter) bool) bool {
	for _, tweets := range m.userTweets {
		for _, tw := range tweets {
			if criteria(tw, tweet) {
				return true
			}
		}
	}
	return false
}

//Deprecated
//func isADuplicateOfCriteria(t1, t2 domain.Tweeter) bool {
//return t1.IsADuplicateOf(t2)
//}
func isEqualToCriteria(t1, t2 domain.Tweeter) bool {
	return t1.Equals(t2)
}

//TweetIsDuplicated returns if a given tweet is duplicated.
//func (m *TweetManager) TweetIsDuplicated(tweet domain.Tweeter) bool {
//	return m.tweetAppearsByCriteria(tweet, isADuplicateOfCriteria)
//}

//TweetExists returns if a given tweet exists
func (m *TweetManager) TweetExists(tweet domain.Tweeter) bool {
	return m.tweetAppearsByCriteria(tweet, isEqualToCriteria)
}
