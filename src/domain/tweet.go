package domain

import "time"

//Tweet is a tweet
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

//NewTweet creates a tweet
func NewTweet(usr, txt string) *Tweet {
	now := time.Now()
	tw := Tweet{User: usr, Text: txt, Date: &now}
	return &tw
}
