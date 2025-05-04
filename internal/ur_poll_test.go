package internal

import (
	"testing"
)

func TestPulling(t *testing.T) {
	sourceSet := "qb"
	poll := NewUrPoll(sourceSet, 3)
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
	poll := NewUrPoll(sourceSet, 2)
	poll.Pull()
	poll.Pull()
	if poll.CanPull() {
		t.Errorf("Expected poll.CanPull() to be false after two pulls but instead got %v", poll.CanPull())
	}
}
