package main

import (
	"fmt"
	"log"

	"github.com/rudolfoborges/pdf2go"
)

func main() {
	reader, err := pdf2go.New("teste.pdf", &pdf2go.Options{})

	if err != nil {
		log.Printf("Err: %v", err)
		panic(err)
	}

	pages, err := reader.Pages()

	if err != nil {
		panic(err)
	}

	fmt.Println("Number of pages:", reader.PagesNumber())

	for _, page := range pages {
		fmt.Println(page.Text())
	}
}
