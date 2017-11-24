package domain

import "time"

//Tweet is a tweet
type Tweet struct {
	user string
	text string
	date *time.Time
}

//NewTweet creates a tweet
func NewTweet(usr, txt string) *Tweet {
	now := time.Now()
	tw := Tweet{user: usr, text: txt, date: &now}
	return &tw
}

//StringTweet returns a tweet as a formatted string
func StringTweet(tw Tweet) string {
	st := tw.user + ": " + tw.text + ", " + tw.date.String()
	return st
}
