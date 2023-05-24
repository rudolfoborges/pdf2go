package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func main() {
	pdfPath := "teste.pdf"
	text, err := ConvertPDFToText(pdfPath)

	if err != nil {
		log.Fatal(err)
	}

	regex := regexp.MustCompile(`Pages:\s+(\d+)`)
	match := regex.FindStringSubmatch(text)

	if len(match) > 1 {
		pages := match[1]
		fmt.Println(pages)
	} else {
		fmt.Println("Número de páginas não encontrado.")
	}
}

func ConvertPDFToText(pdfPath string) (string, error) {
	cmd := exec.Command("pdfinfo", pdfPath, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
