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

//
func GetTweet(tw Tweet) string {
	st := tw.user + ": " + tw.text
	return st
}
