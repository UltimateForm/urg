package internal

import (
	"math/rand"
	"strings"
)

func GenerateRandomString(dataset []*UrPoll, genLength uint8) (string, error) {
	var builder strings.Builder
	builder.Grow(int(genLength))
	for range genLength {
		var targetPoll *UrPoll
		if len(dataset) > 1 {
			elligiblePolls, newCount := FilterList(dataset, func(poll *UrPoll) bool {
				return poll.CanPull()
			})
			targetPoll = elligiblePolls[rand.Intn(newCount)]
		} else {
			targetPoll = dataset[0]
		}

		pulledByte, err := targetPoll.Pull()
		if err != nil {
			return "", err
		}
		builder.WriteByte(pulledByte)
	}
	return builder.String(), nil
}
