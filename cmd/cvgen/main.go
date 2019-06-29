package main

import (
	"cvgen/pkg/cv"
	"flag"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[cvgen] ", log.LstdFlags)

func main() {
	var cvFile string
	var template string
	var format string
	var outputName string

	wkhtmltopdf.SetPath("./utils/wkhtmltopdf/")

	flag.StringVar(&cvFile, "f", "example_cv.yaml", "generates CV from file")
	flag.StringVar(&template, "t", "basic", "selects template")
	flag.StringVar(&format, "form", "pdf", "selects output format")
	flag.StringVar(&outputName, "o", "cv.pdf", "output name")
	flag.Parse()

	logger.Println("cvgen v0.1 - CV Generator")
	logger.Println("Reading CV")

	cvData := cv.ReadFromYamlFile(cvFile)
	logger.Println("Generating document")
	switch format {
	case "pdf":
		cv.GetWriter(cv.WriterPDF).WriteFile(cvData, template, outputName)
	default:
		logger.Fatalf("Unknown format: %s", format)
	}
	logger.Println("Done.")
}
