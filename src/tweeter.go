package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Registers a new user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Pick a name: ")
			user := domain.NewUser(c.ReadLine())
			err := service.Register(user)
			if err != nil {
				c.Printf("Invalid name, %s", err.Error())
				return
			}
			if service.IsRegistered(user) {
				c.Print("Registered successfully")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Logs into twitter",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Insert name: ")
			user := domain.NewUser(c.ReadLine())
			err := service.Login(user)
			if err != nil {
				c.Printf("Invalid login, %s", err.Error())
				return
			}
			c.Print("Login successfull")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Logs out of twitter",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			err := service.Logout()
			if err != nil {
				c.Printf("Couldn't log out, %s", err.Error())
				return
			}
			c.Print("Logged out")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			loggedInUser, err := service.GetLoggedInUser()

			if err != nil {
				c.Printf(err.Error())
				return
			}

			tweet, err := domain.NewTweet(*loggedInUser, text)

			if err != nil {
				c.Printf("Invalid tweet, %s", err.Error())
				return
			}

			err = service.PublishTweet(tweet)

			if err != nil {
				c.Printf("Tweet not published, %s", err.Error())
			} else {
				c.Print("Tweet sent\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "timeline",
		Help: "Shows timeline from logged in user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets, err := service.GetTimeline()
			if err != nil {
				c.Printf("Can't retrieve timeline, %s", err.Error())
				return
			}
			for _, t := range tweets {
				c.Println(t.ToString())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "tweetByID",
		Help: "Finds a tweet by its ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write the ID of the tweet: ")
			id, _ := strconv.Atoi(c.ReadLine())

			tweet, err := service.GetTweetByID(id)
			if err != nil {
				c.Printf("Couldn't retrieve, %s", err.Error())
				return
			}
			c.Printf(tweet.ToString())
			return
		},
	})

	shell.Run()

}
