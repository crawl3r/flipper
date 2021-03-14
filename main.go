package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	//"sync"
	"time"
)

var quietMode bool
var out io.Writer = os.Stdout

func main() {
	var outputFileFlag string
	flag.StringVar(&outputFileFlag, "o", "", "Output a list of the identified IP addresses with their URL and the provider (if identified)")
	quietModeFlag := flag.Bool("q", false, "Only output the data we care about")
	flag.Parse()

	quietMode = *quietModeFlag

	if !quietMode {
		banner()
		fmt.Println("")
	}

	writer := bufio.NewWriter(out)
	targetDomains := make(chan string, 1)
	// var wg sync.WaitGroup

	ch := readStdin()
	go func() {
		//translate stdin channel to domains channel
		for u := range ch {
			targetDomains <- u
		}
		close(targetDomains)
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
}

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