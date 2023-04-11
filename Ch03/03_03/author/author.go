package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("yankee.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	author, err := findAuthor(data)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(string(author))
}

func findAuthor(text []byte) ([]byte, error) {
	author := []byte("Author: ")
	i := bytes.Index(text, author)
	if i == -1 {
		return nil, fmt.Errorf("can't find author")
	}

	i += len(author) // Skip "Author: "
	j := bytes.IndexByte(text[i:], '\n')
	if j == -1 {
		return nil, fmt.Errorf("can't find author")
	}

	return text[i : i+j], nil
}
