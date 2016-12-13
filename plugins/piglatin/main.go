package piglatin

import (
	"strings"

	irc "github.com/thoj/go-ircevent"
)

//Initialize publix exported
func Initialize(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {

	})
}

func convert(message string) string {

	wordArray := strings.Split(message, " ")
	var convertedArray []string
	for _, word := range wordArray {
		convertedArray = append(convertedArray, convertWord(word))
	}
	complete := strings.Join(convertedArray, " ")
	return complete

}

func convertWord(word string) string {
	sArray := strings.Split(word, "")

	firstChar := sArray[0]

	newArray := append(sArray[:0], sArray[1:]...)
	newArray = append(newArray, firstChar)

	fixedString := strings.Join(newArray, "")

	return fixedString + "ay"
}
