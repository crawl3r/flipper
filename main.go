package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
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

	writer := bufio.NewWriter(out)
	loadedFileNames := make(chan string, 1)
	var wg sync.WaitGroup

	ch := readStdin()
	go func() {
		//translate stdin channel to domains channel
		for fn := range ch {
			loadedFileNames <- fn
		}
		close(loadedFileNames)
	}()

	// flush to writer periodically
	t := time.NewTicker(time.Millisecond * 500)
	defer t.Stop()
	go func() {
		for {
			select {
			case <-t.C:
				writer.Flush()
			}
		}
	}()

	// create the rules here
	d2u := &FlipperRule {
		Name: "Dash 2 Underscore",
		Actions: map[string]string {
			"-": "_",
		},
	}

	u2d := &FlipperRule {
		Name: "Underscore 2 Dash",
		Actions: map[string]string {
			"_": "-",
		},
	}

	// populate and store them for looped usage
	rules := []*FlipperRule{d2u, u2d}
	fmt.Println("[*] Rules Loaded:", len(rules))

	for lfn := range loadedFileNames {
		wg.Add(1)
		go func(loadedName string) {
			defer wg.Done()
			if !quietMode {
				fmt.Println("Mutating:", loadedName)
			}

			// loop the rules here
			

		}(lfn)
	}
}

// tool specific functionality
func followRule(rule *FlipperRule) {

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
			url := strings.ToLower(sc.Text())
			if url != "" {
				// strip the http:// or https:// here other the IP look up fails
				// Note: we don't care for multiple entries of the same URL
				final := strings.Replace(url, "http://", "", -1)
				final = strings.Replace(final, "https://", "", -1)
				lines <- final
			}
		}
	}()
	return lines
}