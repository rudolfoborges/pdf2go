package main

import (
	"fmt"
	"log"

	"github.com/rudolfoborges/pdf2go"
)

func main() {
	pdf, err := pdf2go.New("example/doc.pdf", &pdf2go.Options{})

	if err != nil {
		log.Printf("Err: %v", err)
		panic(err)
	}

	fmt.Printf(
		"PDF Info Author (%s) Encrypted (%v) PagesNumber (%d) \n",
		pdf.Author(), pdf.Encrypted(), pdf.PagesNumber(),
	)
}
