package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

//TweetManager is a tweet manager
type TweetManager struct {
	users        []domain.User
	userTweets   map[string][]domain.Tweeter
	loggedInUser domain.User
}

//InitializeManager initializes the manager
func (m *TweetManager) InitializeManager() {
	m.users = make([]domain.User, 0)
	m.userTweets = make(map[string][]domain.Tweeter)
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
	m.users = append(m.users, userToRegister)
	m.userTweets[userToRegister.Name] = make([]domain.Tweeter, 0)
	return nil
}

//IsRegistered verifies that a user is registered
func (m *TweetManager) IsRegistered(user domain.User) bool {
	_, ok := m.userTweets[user.Name]
	return ok
}

func (m *TweetManager) validateLogin(user domain.User) bool {
	for _, u := range m.users {
		if u.Equals(user) {
			return true
		}
	}
	return false
}

//Login logs the user in
func (m *TweetManager) Login(user domain.User) error {
	if m.isLoggedIn() {
		return fmt.Errorf("Already logged in")
	}
	if !m.validateLogin(user) {
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

//GetTweetsFromUser returns all tweets from one user
func (m *TweetManager) GetTweetsFromUser(user domain.User) ([]domain.Tweeter, error) {
	if !m.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline := append(m.userTweets[user.Name])
	return timeline, nil
}

func (m *TweetManager) getTweetsFromFollowing(user domain.User) []domain.Tweeter {
	var tweets []domain.Tweeter
	for _, followedUser := range user.Following {
		followedUserTweets, _ := m.GetTweetsFromUser(followedUser)
		tweets = append(tweets, followedUserTweets...)
	}
	return tweets
}

//GetTimelineFromUser returns all tweets from one user and who they are following
func (m *TweetManager) GetTimelineFromUser(user domain.User) ([]domain.Tweeter, error) {
	if !m.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline := append(m.userTweets[user.Name], m.getTweetsFromFollowing(user)...)
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
	m.userTweets[tweetToPublish.GetUser().Name] = append(m.userTweets[tweetToPublish.GetUser().Name], tweetToPublish)
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
	tweets := m.userTweets[tweet.GetUser().Name]
	tweets = m.deleteElementFromTweets(tweets, tweet)
	m.userTweets[tweet.GetUser().Name] = tweets
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

func isEqualToCriteria(t1, t2 domain.Tweeter) bool {
	return t1.Equals(t2)
}

//TweetExists returns if a given tweet exists
func (m *TweetManager) TweetExists(tweet domain.Tweeter) bool {
	return m.tweetAppearsByCriteria(tweet, isEqualToCriteria)
}

//EditTweetTextByID edits a given tweet by its ID
func (m *TweetManager) EditTweetTextByID(id int, newText string) error {
	tweet, err := m.GetTweetByID(id)
	if err != nil {
		return fmt.Errorf("Coudln't edit tweet, %s", err.Error())
	}
	user, err := m.GetLoggedInUser()
	if err != nil {
		return fmt.Errorf("Coudln't edit tweet, %s", err.Error())
	}

	if !tweet.GetUser().Equals(*user) {
		return fmt.Errorf("You can't edit a tweet that you didn't publish")
	}
	return m.editTweetText(tweet, newText)
}

func (m *TweetManager) editTweetText(t domain.Tweeter, text string) error {
	err := t.SetText(text)
	if err != nil {
		return fmt.Errorf("Coudln't edit tweet, %s", err.Error())
	}
	return nil
}

//FollowUser follows a user
func (m *TweetManager) FollowUser(userName string) error {
	user, err := m.GetLoggedInUser()
	if err != nil {
		return fmt.Errorf("Coudln't follow user, %s", err.Error())
	}
	userToFollow, err := m.getUserByName(userName)

	if err != nil {
		return fmt.Errorf("Couldn't follow user, %s", err.Error())
	}
	if user.Equals(*userToFollow) {
		return fmt.Errorf("Can't follow yourself")
	}
	if user.IsFollowing(*userToFollow) {
		return fmt.Errorf("Can't follow same user twice")
	}
	user.Follow(*userToFollow)
	return nil
}

func (m *TweetManager) getUserByName(name string) (*domain.User, error) {
	for _, user := range m.users {
		if user.Name == name {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not registered")
}

//QuoteTweet returns a new tweet that quotes the given tweet
func QuoteTweet() {

}
