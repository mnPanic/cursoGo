package service

var tweet string

//GetTweet returns the tweet
func GetTweet() string {
	return tweet
}

//PublishTweet Publishes a tweet
func PublishTweet(tw string) {
	tweet = tw
}
