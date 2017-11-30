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
	var manager service.TweetManager
	manager.InitializeManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Registers a new user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Pick a name: ")
			name := c.ReadLine()

			c.Print("Pick a password: ")
			password := c.ReadLine()

			user := domain.NewUser(name, password)
			err := manager.Register(user)
			if err != nil {
				c.Printf("Couldn't register, %s\n", err.Error())
				return
			}
			if manager.IsRegistered(user) {
				c.Print("Registered successfully\n")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Logs into twitter",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Insert name: ")
			name := c.ReadLine()

			c.Print("Insert password: ")
			password := c.ReadLine()

			user := domain.NewUser(name, password)
			err := manager.Login(user)
			if err != nil {
				c.Printf("Invalid login, %s\n", err.Error())
				return
			}
			c.Print("Login successfull\n")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Logs out of twitter",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			err := manager.Logout()
			if err != nil {
				c.Printf("Couldn't log out, %s\n", err.Error())
				return
			}
			c.Print("Logged out\n")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			loggedInUser, err := manager.GetLoggedInUser()

			if err != nil {
				c.Printf(err.Error())
				return
			}

			tweet, err := domain.NewTweet(*loggedInUser, text)

			if err != nil {
				c.Printf("Invalid tweet, %s\n", err.Error())
				return
			}

			err = manager.PublishTweet(tweet)

			if err != nil {
				c.Printf("Tweet not published, %s\n", err.Error())
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

			tweets, err := manager.GetTimeline()
			if err != nil {
				c.Printf("Can't retrieve timeline, %s\n", err.Error())
				return
			}
			for _, t := range tweets {
				c.Println(t)
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

			tweet, err := manager.GetTweetByID(id)
			if err != nil {
				c.Printf("Couldn't retrieve, %s\n", err.Error())
				return
			}
			c.Print(tweet)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "deleteTweet",
		Help: "Deletes a tweet by its ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Which tweet do you want to delete?: ")

			id, _ := strconv.Atoi(c.ReadLine())
			err := manager.DeleteTweetByID(id)
			if err != nil {
				c.Printf("Coudln't delete tweet, %s\n", err.Error())
				return
			}
			c.Print("Tweet deleted successfully\n")
			return
		},
	})

	shell.Run()

}
