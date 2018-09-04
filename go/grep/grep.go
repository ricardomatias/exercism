package grep

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

// FLAGS
// - `-n` Print the line numbers of each matching line.
// - `-l` Print only the names of files that contain at least one matching line.
// - `-i` Match line using a case-insensitive comparison.
// - `-v` Invert the program -- collect all lines that fail to match the pattern.
// - `-x` Only match entire lines, instead of lines that contain a match.

// Match is a grep-like match
type Match struct {
	lineNumber int
	line       string
}

func processFile(file string) []string {
	return regexp.MustCompile("\r|\n|\f").Split(file, -1)
}

func findMatch(text []string, pattern string) []Match {
	matches := []Match{}
	reg := regexp.MustCompile(pattern)

	for lineNumber, line := range text {
		if match := reg.FindString("(?i)" + line); match != "" {
			matches = append(matches, Match{lineNumber, line})
		}
	}

	return matches
}

// Search a file with a pattern
func Search(pattern string, flags, files []string) (searchResult []string) {
	matches := make(map[string][]Match)

	fmt.Println(files)

	for _, filename := range files {
		file, err := ioutil.ReadFile(filename) // For read access.

		if err != nil {
			log.Fatal(err)
		}

		text := processFile(string(file))

		matches[filename] = findMatch(text, pattern)

		if len(flags) == 0 {
			for _, textMatches := range matches {
				for _, match := range textMatches {
					searchResult = append(searchResult, match.line)
				}
			}
		}

		if flabs
	}

	return searchResult
}
