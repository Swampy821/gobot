package advice

import "testing"

func TestGetAdvice(t *testing.T) {
	str := getAdvice()
	if str == "" {
		t.Error("Did not return advice :(")
	}
}
