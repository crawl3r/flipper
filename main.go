package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var quietMode bool
var out io.Writer = os.Stdout

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

	// populate and store them for looped usage
	rules := []*FlipperRule{d2u, u2d}
	fmt.Println("[*] Rules Loaded:", len(rules))

	for lfn := range loadedFileNames {
		if !quietMode {
			fmt.Println("Mutating:", lfn)
		}

		// loop the rules here
		for _, r := range rules {
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
		fmt.Println("o:", loadedFileName, "n:", newVal)
	} else {
		// if not, we want to loop the rules map and apply each one, one at a time
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