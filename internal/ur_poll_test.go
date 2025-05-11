package internal

import (
	"testing"
)

func TestPulling(t *testing.T) {
	sourceSet := "qb"
	poll, _ := NewUrPoll(sourceSet, 3)
	runeByte, err := poll.Pull()
	if err != nil {
		t.Errorf("Expected successful pull but instead encountered error %v", err)
	}
	pulledByte := byte(runeByte)
	if pulledByte != 'q' && pulledByte != 'b' {
		t.Errorf("Expected pulled rune to be 'q' or 'b' but instead got %c", pulledByte)
	}
}

func TestCanPullIsProperlyFalse(t *testing.T) {
	sourceSet := "qwertyuiop"
	poll, _ := NewUrPoll(sourceSet, 2)
	poll.Pull()
	poll.Pull()
	if poll.CanPull() {
		t.Errorf("Expected poll.CanPull() to be false after two pulls but instead got %v", poll.CanPull())
	}
	if _, err := poll.Pull(); err == nil {
		t.Errorf("Expected poll.Pull to return error after max pulls but instead got %v", err)
	}
}

func TestReturnsErrorIfAttemptedCreationWithEmptyString(t *testing.T) {
	sourceSet := ""
	poll, err := NewUrPoll(sourceSet, 2)
	if poll != nil {
		t.Errorf("Expected poll to be nil but instead got %+v", poll)
	}
	if err == nil {
		t.Error("Expected error to be defined but instead got nil")
	}
}

func TestKeepOriginalSourceWhenNotConsumable(t *testing.T) {
	sourceset := "qwertyuiop"
	poll, _ := NewUrPoll(sourceset, 10)
	for range 10 {
		poll.Pull()
	}
	if poll.CanPull() {
		t.Error("Expected poll.CanPull() to return false but instead returns true")
	}
	if poll.source != sourceset {
		t.Errorf("Expected poll.source to equal '%v', but instead it is '%v'", sourceset, poll.source)
	}
}

func TestConsumablePollIsDepleted(t *testing.T) {
	sourceset := "qwertyuiop"
	poll, _ := NewConsumableUrPoll(sourceset, 10)
	for range 10 {
		poll.Pull()
	}
	if poll.CanPull() {
		t.Error("Expected poll.CanPull() to return false but instead returns true")
	}
	if poll.source != "" {
		t.Errorf("Expected poll.source to equal empty string, but instead it is '%v'", poll.source)
	}
}
