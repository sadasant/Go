package main

import (
	"encoding/json"
	"flag"
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"io/ioutil"
	"time"
)

type Config struct {
	Nick     string
	Password string
	Server   string
	Channel  string
}

var config Config

func readConfig(filename string) {
	defer println("Reading Config at: " + filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
		return
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		println(err)
		return
	}
}

func main() {
	flag.Parse()
	readConfig(flag.Arg(0))
	client := irc.SimpleClient(config.Nick)
	client.AddHandler("connected", connected)
	client.AddHandler("disconnected", disconnected)
	quit := make(chan bool)
	err := client.Connect(config.Server)
	if err != nil {
		fmt.Printf("Connection Error: %s\n", err)
	}
	<-quit
}

func connected(conn *irc.Conn, line *irc.Line) {
	println("Connected!")
	conn.Pass(config.Password)
	conn.Join(config.Channel)
	conn.Privmsg(config.Channel, "¡Saludos!")
	constantlyReadTwitter(conn)
}

func constantlyReadTwitter(conn *irc.Conn) {
	twitter, err := readTwitter()
	defer constantlyReadTwitter(conn)
	defer time.Sleep(10 * time.Second)
	if err != nil {
		println(err)
		return
	}
	for _, tweet := range twitter.Results {
		if tweet.New {
			conn.Privmsg(config.Channel, "@" + tweet.From_user + ": \"" + tweet.Text + "\"")
			conn.Privmsg(config.Channel, tweet.URL)
		}
	}
}

func disconnected(conn *irc.Conn, line *irc.Line) {
	println("Disconnected")
}
