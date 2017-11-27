package main

import (
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
				c.Print("Added successfully")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Who are you? ")

			user := domain.NewUser(c.ReadLine())

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			err := service.PublishTweet(tweet)

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
				c.Println(domain.StringTweet(t))
			}

			return
		},
	})

	shell.Run()

}
