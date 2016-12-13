package catfacts

import (
	"log"

	irc "github.com/thoj/go-ircevent"
)

// Initialize is exported
func Initialize(conn *irc.Connection) {
	log.Print("Intiialized catfacts")
	listen(conn)
}

func listen(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		handle(event, conn)
	})
}
