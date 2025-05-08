package internal

import (
	"errors"
	"math/rand"
	"unicode/utf8"
)

type UrPoll struct {
	source     string
	MaxPulls   uint8
	pulled     uint8
	consumable bool
}

func (poll *UrPoll) CanPull() bool {
	return poll.pulled < poll.MaxPulls
}

func (poll *UrPoll) Pull() (byte, error) {
	if !poll.CanPull() {
		return 0, errors.New("Unable to pull fom UrPoll")
	}
	defer func() {
		poll.pulled += 1
	}()
	strValue := poll.source
	srcLength := utf8.RuneCountInString(strValue)
	targetIndex := rand.Intn(int(srcLength))
	targetByte := strValue[targetIndex]
	if poll.consumable {
		poll.source = poll.source[:targetIndex] + poll.source[targetIndex+1:]
	}
	return targetByte, nil
}

func newUrPoll(source string, maxPulls uint8, consumable bool) (*UrPoll, error) {
	if source == "" {
		return nil, errors.New("source string cannot be empty")
	}
	return &UrPoll{source: source, MaxPulls: maxPulls, consumable: consumable}, nil
}

func NewUrPoll(source string, maxPulls uint8) (*UrPoll, error) {
	return newUrPoll(source, maxPulls, false)
}

func NewConsumableUrPoll(source string, maxPulls uint8) (*UrPoll, error) {
	return newUrPoll(source, maxPulls, true)

}
