package catfacts

import "testing"

func TestTest(t *testing.T) {
	success := test(".catfact")
	if !success {
		t.Error("Failed test")
	}

	shouldfail := test("blah blah catfact blah blah")
	if shouldfail {
		t.Error("Should have failed this test")
	}
}

func TestGetFact(t *testing.T) {
	fact := getFact()
	if fact == "" {
		t.Error("Fact not returned!")
	}
}
