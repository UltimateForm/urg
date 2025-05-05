package cmd

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/UltimateForm/urg/internal"
	"github.com/spf13/cobra"
)

type RgSource string
type RgControl []RgSource

const uppers string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowers string = "abcdefghijklmnopqrstuvwxyz"
const numbers string = "0123456789"
const specialCharacters string = "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~\\ "

var allSources = [4]string{uppers, lowers, numbers, specialCharacters}

var min uint8 = 14
var max uint8 = 16
var includeNum, includeSpecial, includeLowers, includeUppers uint8

func stringRun(cmd *cobra.Command, args []string) {
	// log.Printf("Running with data %+v", map[string]uint8{"uppers": includeUppers, "lowers": includeLowers, "numbers": includeNum, "specials": includeSpecial})
	genLength := min + uint8(rand.Int31n((int32(max)+1)-int32(min)))
	var dataset []*internal.UrPoll
	charsetMap := map[string]uint8{
		uppers:            includeUppers,
		lowers:            includeLowers,
		numbers:           includeNum,
		specialCharacters: includeSpecial,
	}
	for charSet, charInclusionCount := range charsetMap {
		if charInclusionCount == 0 {
			continue
		}
		poll, err := internal.NewUrPoll(charSet, charInclusionCount)
		if err != nil {
			log.Fatal(err.Error())
		}
		dataset = append(dataset, poll)
	}

	generatedStr, err := internal.GenerateRandomString(dataset, genLength)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(generatedStr)
}

var stringCmd = &cobra.Command{
	Use:     "str",
	Short:   "generate random string",
	Run:     stringRun,
	Example: "urg str --min=30 --max=50 -lnsu\nurg str -un\nurg str",
}

func init() {
	maxUint8 := strconv.Itoa(math.MaxUint8)
	stringCmd.Flags().Uint8VarP(&includeNum, "numbers", "n", 0, "include numbers")
	stringCmd.Flag("numbers").NoOptDefVal = maxUint8
	stringCmd.Flags().Uint8VarP(&includeSpecial, "special", "s", 0, "include special characters")
	stringCmd.Flag("special").NoOptDefVal = maxUint8
	stringCmd.Flags().Uint8VarP(&includeUppers, "uppers", "u", 255, "include uppercase letters")
	stringCmd.Flag("uppers").NoOptDefVal = maxUint8
	stringCmd.Flags().Uint8VarP(&includeLowers, "lowers", "l", 0, "include lowercase letters")
	stringCmd.Flag("lowers").NoOptDefVal = maxUint8
	stringCmd.Flags().Uint8Var(&max, "max", max, "max length of generated string")
	stringCmd.Flags().Uint8Var(&min, "min", min, "min length of generated string")

	rootCmd.AddCommand(stringCmd)
}
