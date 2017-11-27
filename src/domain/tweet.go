package domain

import "time"

//Tweet is a tweet
type Tweet struct {
	User User
	Text string
	Date *time.Time
}

//NewTweet creates a tweet
func NewTweet(usr User, txt string) *Tweet {
	now := time.Now()
	tw := Tweet{User: usr, Text: txt, Date: &now}
	return &tw
}

//StringTweet returns a tweet as a formatted string
func StringTweet(tw Tweet) string {
	st := tw.User.Name + ": " + tw.Text + ", " + tw.Date.String()
	return st
}
