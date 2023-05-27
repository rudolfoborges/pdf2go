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
		panic(err)
	}

	html, err := reader.Html()
	fmt.Printf("HTML: %s", html)

	pages, err := reader.Pages()

	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		fmt.Println(page.Html())
		fmt.Println(page.Number)
	}
}
