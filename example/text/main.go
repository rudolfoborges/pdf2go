package main

import (
	"fmt"
	"log"

	"github.com/rudolfoborges/pdf2go"
)

func main() {
	reader, err := pdf2go.New("example/doc.pdf", pdf2go.Config{})

	if err != nil {
		log.Printf("Err: %v", err)
		return
	}

	text, err := reader.Text()

	if err != nil {
		log.Printf("Err: %v", err)
		return
	}

	fmt.Printf("TEXT: %s", text)

	pages, err := reader.Pages()

	if err != nil {
		log.Printf("Err: %v", err)
		return
	}

	for _, page := range pages {
		fmt.Println(page.Text())
		fmt.Println(page.Number)
	}
}
