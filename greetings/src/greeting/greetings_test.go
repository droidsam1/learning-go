package greeting

import (
	"regexp"
	"testing"
)

func TestShouldReturnGreetingWithANameWhenReceivedAName(t *testing.T) {
	name := "Golang"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Golang")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`calling Hello("Golang)" returns %q, with error = %v. Returned msg is expected to contains the input name %#q`, msg, err, name)
	}
}

func TestShouldFailWhenNoNameIsReceived(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`calling Hello("") returns %q, with error %v. Expected to return a empty msg and an not nil error`, msg, err)
	}
}
