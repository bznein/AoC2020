package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bznein/AoC2020/pkg/timing"
)

func ReadInput(path string) string {
	defer timing.TimeTrack(time.Now())
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return strings.TrimSuffix(text, "\n")
}

func InputToIntSlice(input string) []int {
	ints := []int{}
	for _, s := range strings.Split(input, "\n") {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func InputToStringSlice(input string) []string {
	return strings.Split(input, "\n")
}

func AsInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Called AsInt with non-int string")
	}
	return i
}
