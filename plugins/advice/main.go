package advice

import (
	"io/ioutil"
	"log"
	"net/http"

	"strings"

	"github.com/Jeffail/gabs"
	irc "github.com/thoj/go-ircevent"
)

//Initialize  exported
func Initialize(conn *irc.Connection) {
	log.Print("Initialized Advice")

	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		handle(event, conn)
	})
}

func test(message string) bool {
	return strings.ToLower(message) == ".advice"
}

func getAdvice() string {
	resp, _ := http.Get("http://api.adviceslip.com/advice")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	parsedJSON, _ := gabs.ParseJSON(body)

	advicestr, _ := parsedJSON.Path("slip.advice").Data().(string)
	return advicestr
}

func handle(event *irc.Event, conn *irc.Connection) {
	if test(event.Message()) {
		channel := event.Arguments[0]
		adviceStr := getAdvice()
		conn.Privmsg(channel, strings.TrimLeft(event.User, "~")+": "+adviceStr)

	}
}
