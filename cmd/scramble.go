package cmd

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/UltimateForm/urg/internal"

	"strings"

	"github.com/spf13/cobra"
)

func scrambleRun(cmd *cobra.Command, args []string) {
	fullStr := strings.Join(args, " ")
	strLen := uint8(utf8.RuneCountInString(fullStr))
	poll, err := internal.NewConsumableUrPoll(fullStr, strLen)
	if err != nil {
		log.Fatal(err.Error())
	}
	var stringBuilder strings.Builder
	stringBuilder.Grow(int(strLen))
	for range strLen {
		bt, err := poll.Pull()
		if err != nil {
			log.Fatal(err.Error())
		}
		stringBuilder.WriteByte(bt)
	}
	fmt.Println(stringBuilder.String())
}

var scrambleCmd = &cobra.Command{
	Use:     "scr",
	Short:   "scramble a string",
	Example: "urg scr lorem ipsum",
	Args:    cobra.MinimumNArgs(1),
	Run:     scrambleRun,
}

func init() {
	rootCmd.AddCommand(scrambleCmd)
}
