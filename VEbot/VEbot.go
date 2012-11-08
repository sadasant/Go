package main

import (
	"flag"
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"strings"
	"time"
)

var config Config

func main() {
	flag.Parse()
	readConfig(flag.Arg(0)) // config.go
	client := irc.SimpleClient(config.Nick)
	client.SSL = true
	client.AddHandler("connected", connected)
	client.AddHandler("privmsg", message)
	client.AddHandler("disconnected", disconnected)
	quit := make(chan bool)
	err := client.Connect(config.Server)
	if err != nil {
		fmt.Printf("Connection Error: %s\n", err)
	}
	<-quit
}

func message(conn *irc.Conn, line *irc.Line) {
	fmt.Printf("%s\n", line.Args)
	msgs := strings.Split(line.Args[1], ":")
	if msgs[0] == config.Nick {
		// Improve this with a JSON file full of messages and responses
		switch {
		case strings.LastIndex(msgs[1], "hola") >= 0:
			conn.Privmsg(config.Channel, line.Nick+": Hola!")
		case strings.LastIndex(msgs[1], "o/") >= 0:
			conn.Privmsg(config.Channel, line.Nick+": \\o")
		case strings.LastIndex(msgs[1], "hello") >= 0:
			conn.Privmsg(config.Channel, line.Nick+": Hi!")
		}
	}
}

func connected(conn *irc.Conn, line *irc.Line) {
	println("Connected!")
	conn.Pass(config.Password)
	conn.Join(config.Channel)
	conn.Privmsg(config.Channel, "¡Saludos!")
	constantlyReadTwitter(conn)
}

func constantlyReadTwitter(conn *irc.Conn) {
	twitter, err := readTwitter() // twitter.go
	defer constantlyReadTwitter(conn)
	defer time.Sleep(10 * time.Second)
	if err != nil {
		println(err)
		return
	}
	for _, tweet := range twitter.Results {
		if tweet.New {
			conn.Privmsg(config.Channel, "@"+tweet.From_user+": \""+tweet.Text+"\"")
			conn.Privmsg(config.Channel, tweet.URL)
		}
	}
}

func disconnected(conn *irc.Conn, line *irc.Line) {
	println("Disconnected")
}
