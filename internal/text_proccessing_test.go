package internal

import (
	"testing"
	"unicode/utf8"
)

func TestGenerateRandomString(t *testing.T) {
	polls := []*UrPoll{}
	strings := [3]string{"", "asdfghjkl", "zxcvbnm"}
	for _, i := range strings {
		poll, _ := NewUrPoll(i, 10)
		if poll == nil {
			t.Fatalf("Test setup failed, poll creation with %v failed", i)
		}
		polls = append(polls, poll)
	}
	generatedStr, err := GenerateRandomString(polls, 10)
	if err != nil {
		t.Errorf("Expected error to be nil but instead got %v", err)
	}
	if utf8.RuneCountInString(generatedStr) != 10 {
		t.Errorf("Expected generated string to have length 10 but instead it has %v", utf8.RuneCountInString(generatedStr))
	}
}

func TestWorksWithASinglePoll(t *testing.T) {
	poll, err := NewUrPoll("qwertyuiop", 5)
	if err != nil {
		t.Fatalf("Expected error to be nil but instead got %v", err)
	}
	generatedStr, err := GenerateRandomString([]*UrPoll{poll}, 5)
	if err != nil {
		t.Fatalf("Expected generation err to be nil but instead got %v", err)
	}
	if poll.CanPull() {
		t.Errorf("Expected poll to be exhausted during generation. %+v", poll)
	}
	if utf8.RuneCountInString(generatedStr) != 5 {
		t.Errorf("Expected generated string to have length 5 but instead has %v", utf8.RuneCountInString(generatedStr))
	}
}
