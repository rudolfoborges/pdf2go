package main

import (
	"fmt"
	"log"

	"github.com/rudolfoborges/pdf2go"
)

func main() {
	reader, err := pdf2go.New("example/doc.pdf", &pdf2go.Options{})

	if err != nil {
		log.Printf("Err: %v", err)
		panic(err)
	}

	text, err := reader.Text()
	fmt.Printf("TEXT: %s", text)

	pages, err := reader.Pages()

	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		fmt.Println(page.Text())
		fmt.Println(page.Number)
	}
}
