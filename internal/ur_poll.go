package internal

import (
	"errors"
	"unicode/utf8"

	"math/rand"
)

type UrPoll struct {
	source   string
	MaxPulls uint8
	pulled   uint8
}

func (poll *UrPoll) CanPull() bool {
	return poll.pulled < poll.MaxPulls
}

func (poll *UrPoll) Pull() (byte, error) {
	defer func() {
		poll.pulled += 1
	}()
	if !poll.CanPull() {
		return 0, errors.New("Unable to pull fom UrPoll")
	}
	strValue := poll.source
	srcLength := utf8.RuneCountInString(strValue)
	return strValue[rand.Intn(int(srcLength-1))], nil
}

func NewUrPoll(source string, maxPulls uint8) *UrPoll {
	return &UrPoll{source: source, MaxPulls: maxPulls}
}
