package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-pdf/fpdf"
)

func combine(template string, company string, position string, team string) string {
	template = strings.Replace(template, "{position}", position, -1)
	template = strings.Replace(template, "{team}", team, -1)
	template = strings.Replace(template, "{company}", company, -1)
	return template
}

func get_template() string {
	// Read template from file.
	file, err := os.Open("template")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var template string
	for scanner.Scan() {
		template += scanner.Text() + "\n"
	}

	return template
}

func output_pdf(content string, filename string) error {
	log.Println("Outputing cover letter...")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	page_width, _ := pdf.GetPageSize()
	_, _, r, _ := pdf.GetMargins()
	width := page_width - 2*r
	height := 5.0
	pdf.MultiCell(width, height, content, "", "", false)
	err := pdf.OutputFileAndClose(filename)
	return err
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the position name: ")
	scanner.Scan()
	position := scanner.Text()

	fmt.Print("Enter the team name: ")
	scanner.Scan()
	team := scanner.Text()

	fmt.Print("Enter the company name: ")
	scanner.Scan()
	company := scanner.Text()

	template := get_template()
	content := combine(template, company, position, team)

	err := output_pdf(content, "output.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
