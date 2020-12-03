package input

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(path string) string {
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
