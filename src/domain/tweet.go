package domain

import (
	"fmt"
	"strconv"
	"time"
)

//Tweet is a tweet
type Tweet struct {
	User User
	Text string
	Date *time.Time
	ID   int
}

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

//NewTweet creates a tweet
func NewTweet(usr User, txt string) (*Tweet, error) {
	now := time.Now()
	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}
	tw := Tweet{User: usr, Text: txt, Date: &now, ID: getNextID()}
	return &tw, nil
}

//String returns a formatted string of the tweet, implements Stringer
func (tw Tweet) String() string {
	//date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
	id := strconv.Itoa(tw.ID)
	formattedString := fmt.Sprintf("[%s] @%s: %s", id, tw.User.Name, tw.Text)
	return formattedString
}

//Equals returns if two tweets are the same
func (tw Tweet) Equals(other Tweet) bool {
	return (tw.Date == other.Date &&
		tw.ID == other.ID &&
		tw.Text == other.Text &&
		tw.User.Equals(other.User))
}
