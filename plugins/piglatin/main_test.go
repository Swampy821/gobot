package piglatin

import "testing"

func TestConvert(t *testing.T) {

	if convert("hello") != "ellohay" {
		t.Error("hello => ellohay returned " + convert("hello"))
	}

	if convert("hello world") != "ellohay orldway" {
		t.Error("hello world => ellohay orldway returned " + convert("hello world"))
	}

}
