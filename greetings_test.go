package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Abhay"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Abhay")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Abhay") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, wamt "", error`, msg, err)
	}
}
