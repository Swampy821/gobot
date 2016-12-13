package catfacts

import (
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/Jeffail/gabs"
	irc "github.com/thoj/go-ircevent"
)

func test(message string) bool {
	return strings.ToLower(message) == ".catfact"
}

func getFact() string {
	resp, _ := http.Get("http://catfacts-api.appspot.com/api/facts")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	parsedJSON, _ := gabs.ParseJSON(body)

	fact, _ := parsedJSON.S("facts").Children()
	return fact[0].Data().(string)
}

func handle(event *irc.Event, conn *irc.Connection) {
	if test(event.Message()) {
		channel := event.Arguments[0]
		fact := getFact()
		conn.Privmsg(channel, strings.TrimLeft(event.User, "~")+": "+fact)

	}
}
