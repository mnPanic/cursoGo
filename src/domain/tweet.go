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
	String() string
	Equals(Tweeter) bool
	GetUser() User
	GetID() int
	GetDate() *time.Time
	GetText() string
}

//TextTweet is a tweet that has just text
type TextTweet struct {
	user User
	date *time.Time
	id   int
	text string
}

//NewTextTweet returns a new TextTweet
func NewTextTweet(usr User, txt string) (*TextTweet, error) {
	if txt == "" {
		return nil, fmt.Errorf("Tweet can't have no text")
	}
	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}
	now := time.Now()
	textTweet := TextTweet{user: usr, date: &now, id: getNextID(), text: txt}
	return &textTweet, nil
}

//GetUser returns the user that posted the tweet
func (t TextTweet) GetUser() User {
	return t.user
}

//GetDate returns the date at which the tweet was posted
func (t TextTweet) GetDate() *time.Time {
	return t.date
}

//GetID returns the ID of the tweet
func (t TextTweet) GetID() int {
	return t.id
}

//GetText returns the text of the text tweet
func (t TextTweet) GetText() string {
	return t.text
}

func (t TextTweet) String() string {
	//date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
	formattedString := fmt.Sprintf("[%d] @%s: %s", t.id, t.user, t.text)
	return formattedString
}

//Equals returns if a given TextTweet is the same as another
func (t TextTweet) Equals(other Tweeter) bool {
	return (t.date == other.GetDate() &&
		t.id == other.GetID() &&
		t.user.Equals(other.GetUser()) &&
		t.text == other.GetText())
}

//ImageTweet is a tweet that contains an image
type ImageTweet struct {
	TextTweet
	imageURL string
}

//NewImageTweet returns a new ImageTweet
func NewImageTweet(user User, text string, url string) (*ImageTweet, error) {
	if url == "" {
		return nil, fmt.Errorf("Cant create an image tweet without an URL")
	}

	textTweet, err := NewTextTweet(user, text)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create ImageTweet, %s", err.Error())
	}

	imageTweet := ImageTweet{TextTweet: *textTweet, imageURL: url}
	return &imageTweet, nil
}

//GetURL returns the URL of the imageTweet
func (t ImageTweet) GetURL() string {
	return t.imageURL
}

//String returns a formatted string of the ImageTweet
func (t ImageTweet) String() string {
	formattedString := fmt.Sprintf("%s\n%s", t.TextTweet, t.imageURL)
	return formattedString
}

//Equals returns if a given TextTweet is the same as another
func (t ImageTweet) Equals(other ImageTweet) bool {
	return (t.TextTweet.Equals(other.TextTweet) &&
		t.imageURL == other.imageURL)
}

//QuoteTweet is a tweet that quotes another
type QuoteTweet struct {
	TextTweet
	quotedTweet Tweeter
}

//NewQuoteTweet returns a new QuoteTweet
func NewQuoteTweet(user User, text string, quoted Tweeter) (*QuoteTweet, error) {
	textTweet, err := NewTextTweet(user, text)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create QuoteTweet, %s", err.Error())
	}
	quoteTweet := QuoteTweet{TextTweet: *textTweet, quotedTweet: quoted}
	return &quoteTweet, nil
}

//GetQuotedTweet returns the quotedtweet of the QuoteTweet
func (t QuoteTweet) GetQuotedTweet() Tweeter {
	return t.quotedTweet
}

//String returns a formatted string of the QuoteTweet
func (t QuoteTweet) String() string {
	formattedString := fmt.Sprintf("%s %q", t.TextTweet, t.quotedTweet)
	return formattedString
}

//Equals returns if a given TextTweet is the same as another
func (t QuoteTweet) Equals(other QuoteTweet) bool {
	return (t.TextTweet.Equals(other.TextTweet) && t.quotedTweet.Equals(other.quotedTweet))
}
