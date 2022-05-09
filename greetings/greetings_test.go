package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloFuncWithName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloFuncWithEmptyStr(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

func TestHelloFuncWithWhitespace(t *testing.T) {
    msg, err := Hello(" ")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

func TestHellosFuncWithNamesArray(t *testing.T) {
	names := []string { "Gladys", "Joe " }
	messages, err := Hellos(names)
	if len(messages) != len(names) || err != nil {
		t.Fatalf(`Hellos([]) = %q, %v, error`, messages, err)
	}

}

func TestHellosFuncWithEmptyArray(t *testing.T) {
	messages, err := Hellos([]string {})
	if err != nil {
		t.Fatalf(`Hellos([]) = %q, %v, error`, messages, err)
	}

}