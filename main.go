package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	// "math"
	"os"
	"strings"
)

var quietMode bool
var out io.Writer = os.Stdout
var currentCombinations []string

// FlipperRule is a struct that we populate to handle a set of rules to perform against the file name. Allows easy 'modular' rules to be added without affecting codebase
type FlipperRule struct {
	Name 	string
	Actions map[string]string
	OneShot bool
}

func main() {
	quietModeFlag := flag.Bool("q", false, "Only output the data we care about")
	flag.Parse()

	quietMode = *quietModeFlag

	if !quietMode {
		banner()
		fmt.Println("")
	}

	loadedFileNames := make(chan string, 1)

	ch := readStdin()
	go func() {
		//translate stdin channel to domains channel
		for fn := range ch {
			loadedFileNames <- fn
		}
		close(loadedFileNames)
	}()

	// create the rules here
	d2u := &FlipperRule {
		Name: "Dash 2 Underscore",
		Actions: map[string]string {
			"-": "_",
		},
		OneShot: true,
	}

	u2d := &FlipperRule {
		Name: "Underscore 2 Dash",
		Actions: map[string]string {
			"_": "-",
		},
		OneShot: true,
	}

	l337 := &FlipperRule {
		Name: "1337 5p34k",
		Actions: map[string]string {
			"a": "4",
			"e": "3",
			"o": "0",
			"l": "1",
			"s": "5",
			"t": "7",
		},
		OneShot: false,
	}

	// populate and store them for looped usage
	rules := []*FlipperRule{d2u, u2d, l337}
	fmt.Println("[*] Rules Loaded:", len(rules))

	for lfn := range loadedFileNames {
		if !quietMode {
			fmt.Println("Mutating:", lfn)
		}

		// first, we output the raw so we always have the original value
		fmt.Println(lfn)

		// loop the rules here
		for _, r := range rules {
			currentCombinations = []string{}
			followRule(lfn, r)
		}
	}
}

// tool specific functionality
func followRule(loadedFileName string, rule *FlipperRule) {
	// are we a oneshot rule, if so, we do all of it at once
	if rule.OneShot {
		newVal := loadedFileName
		for k, v := range rule.Actions {
			newVal = strings.Replace(newVal, k, v, -1)
		}

		if newVal != loadedFileName {
			fmt.Println(newVal)
		}
	} else {
		// if not, we want to loop the rules map and apply each one, one at a time
		
		//actionLength := float64(len(rule.Actions))
		//maxAttempts := math.Pow(actionLength, actionLength)
	
		// start looping the character - 1st pass, depth = 1, replaces all chars if target found
		foundChars := []string{}
		for _, c := range loadedFileName {
			if _, ok := rule.Actions[string(c)]; ok {
				foundChars = append(foundChars, string(c))
			} 
		}

		// recursively loop through the 'found' chars for every possible combination.
		// use every combination and then generate the new words with these replacement combinations
		getCombinationsRecursive("", foundChars)
		fmt.Println("-------------------------------")
		fmt.Println(currentCombinations)
		fmt.Println("-------------------------------")

		// using all our combinations, 1337 the words relative to current
		for _, combo := range currentCombinations {
			currComboVal := loadedFileName
			
			for _, c := range combo {
				// need better logic here to only replace the next one found?
				currComboVal = strings.Replace(currComboVal, string(c), rule.Actions[string(c)], 1)
			}
			fmt.Println(currComboVal)
		}
	}
}

func getCombinationsRecursive(current string, chars []string) {
	for idx, c := range chars {

		if !existsInArray(currentCombinations, current + c) {
			currentCombinations = append(currentCombinations, current + c)
		}

		if (idx + 1) < len(chars) {
			getCombinationsRecursive(current + c, chars[idx + 1:])
		}
	}
}

// util, generic tool stuff
func banner() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Flipper -> Crawl3r")
	fmt.Println("Reads stdin and mutates each value, line by line to get a range of mutations.")
	fmt.Println("")
	fmt.Println("Run again with -q for cleaner output")
	fmt.Println("---------------------------------------------------")
}

func readStdin() <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			lines <- sc.Text()
		}
	}()
	return lines
}

func existsInArray(arr []string, s string) bool {
	for _, curr := range arr {
		if curr == s {
			return true
		}
	}
	return false
}

func stringContainsChar(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

func generateStringOfLength(c rune, l int) string {
	final := ""
	for i := 0; i < l; i++ {
		final += string(c)
	}
	return final
}