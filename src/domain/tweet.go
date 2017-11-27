package domain

import "time"

//Tweet is a tweet
type Tweet struct {
	User User
	Text string
	Date *time.Time
	ID   int
}

var currentID = -1

//GetNextID returns the id of the next tweet
func getNextID() int {
	currentID++
	return (currentID)
}

//NewTweet creates a tweet
func NewTweet(usr User, txt string) *Tweet {
	now := time.Now()
	tw := Tweet{User: usr, Text: txt, Date: &now, ID: getNextID()}
	return &tw
}

//StringTweet returns a tweet as a formatted string
func StringTweet(tw Tweet) string {
	date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
	st := tw.User.Name + ": " + tw.Text + ", " + date
	return st
}
