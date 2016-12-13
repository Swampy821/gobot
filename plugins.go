package main

//Add your import
import (
	"./Plugins/advice"
	"./plugins/catfacts"
	irc "github.com/thoj/go-ircevent"
)

func loadPlugins(conn *irc.Connection) {
	//Add initialize functio here
	catfacts.Initialize(conn)
	advice.Initialize(conn)
}
