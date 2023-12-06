package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	var url, wordlist, character string
	var help bool

	flag.StringVar(&url, "u", "", "URL to scan")
	flag.StringVar(&wordlist, "w", "", "Wordlist to use")
	flag.StringVar(&character, "ch", "", "Character to filter")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if url == "" && wordlist == "" {
		fmt.Println("Please enter a URL and a wordlist")
		return
	}

	file, err := os.Open(wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	results := make(chan string)

	for scanner.Scan() {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()

			resp, err := http.Get(url + "/" + word)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				if character == "" {
					results <- url + "/" + word
				} else {
					if strings.Contains(word, character) {
						results <- url + "/" + word
					}
				}
			}
		}(scanner.Text())
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
