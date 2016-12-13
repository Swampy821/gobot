package main

import (
	irc "github.com/thoj/go-ircevent"
)

func connect(config Config) *irc.Connection {
	conn := irc.IRC(config.Name, config.Name)
	conn.Connect(config.Server)
	return conn
}

func main() {

	config := readConfig()
	conn := connect(config)
	conn.AddCallback("001", func(event *irc.Event) {
		for _, channel := range config.Channels {
			conn.Join(channel)
		}
	})
	go loadPlugins(conn)

	conn.Loop()
}
