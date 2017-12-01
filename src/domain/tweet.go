package domain

import (
	"fmt"
	"time"
)

var currentID = -1

//getNextID returns the id of the next tweet
func getNextID() int {
	currentID++
	return (currentID)
}

//GetCurrentID returns the id of the last tweet
func GetCurrentID() int {
	return currentID
}

//ResetCurrentID serves as an initialization, resetting the current ID
func ResetCurrentID() {
	currentID = -1
}

//Tweeter is an interface that defines a tweet
type Tweeter interface {
	String()
	Equals(Tweeter)
	IsADuplicateOf(Tweeter)
	GetUser()
	GetID()
	GetDate()
}

//Tweet is a tweet
type Tweet struct {
	user User
	date *time.Time
	id   int
}

//NewTweet creates a tweet
func NewTweet(usr User) Tweet {
	now := time.Now()
	tweet := Tweet{user: usr, date: &now, id: getNextID()}
	return tweet
}

//GetUser returns the user that posted the tweet
func (t Tweet) GetUser() User {
	return t.user
}

//GetDate returns the date at which the tweet was posted
func (t Tweet) GetDate() *time.Time {
	return t.date
}

//GetID returns the ID of the tweet
func (t Tweet) GetID() int {
	return t.id
}

//String returns a formatted string of the tweet
func (t Tweet) String() string {
	//date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
	formattedString := fmt.Sprintf("[%d] @%s", t.id, t.user)
	return formattedString
}

//Equals returns if two tweets are the same
func (t Tweet) Equals(other Tweet) bool {
	return (t.date == other.date &&
		t.id == other.id &&
		t.user.Equals(other.user))
}

//TextTweet is a tweet that has just text
type TextTweet struct {
	Tweet
	text string
}

//NewTextTweet returns a new TextTweet
func NewTextTweet(user User, txt string) (*TextTweet, error) {
	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}
	textTweet := TextTweet{Tweet: NewTweet(user), text: txt}
	return &textTweet, nil
}

//GetText returns the text of the text tweet
func (t TextTweet) GetText() string {
	return t.text
}

func (t TextTweet) String() string {
	formattedString := fmt.Sprintf("%s: %s", t.Tweet.String(), t.text)
	return formattedString
}

//Equals returns if a given TextTweet is the same as another
func (t TextTweet) Equals(other TextTweet) bool {
	return (t.Tweet.Equals(other.Tweet) &&
		t.text == other.text)
}

//ImageTweet is a tweet that contains an image
type ImageTweet struct {
	Tweet
	imageURL string
}

//NewImageTweet returns a new ImageTweet
func NewImageTweet(user User, url string) (*ImageTweet, error) {
	if url == "" {
		return nil, fmt.Errorf("Cant publish an image tweet without an URL")
	}
	imageTweet := ImageTweet{Tweet: NewTweet(user), imageURL: url}
	return &imageTweet, nil
}

//GetURL returns the URL of the imageTweet
func (t ImageTweet) GetURL() string {
	return t.imageURL
}

//String returns a formatted string of the ImageTweet
func (t ImageTweet) String() string {
	formattedString := fmt.Sprintf("%s: %s", t.Tweet.String(), t.imageURL)
	return formattedString
}

//Equals returns if a given TextTweet is the same as another
func (t ImageTweet) Equals(other ImageTweet) bool {
	return (t.Tweet.Equals(other.Tweet) &&
		t.imageURL == other.imageURL)
}

//QuoteTweet is a tweet that quotes another
type QuoteTweet struct {
	TextTweet
	quoted Tweet
}

func NewImageTweet() {}

func NewQuoteTweet() {

}
