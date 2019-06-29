package cv

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"strings"
)

type pdfWriter struct {
}

func(w pdfWriter) WriteFile(cv *CV, templateName, file string) {

	var writer bytes.Buffer
	t, _ := template.ParseFiles("templates/" + templateName + ".html")
	if err := t.Execute(&writer, cv); err != nil {
		panic(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(writer.String())))

	err = pdfg.Create()
	if err != nil {
		panic(err)
	}

	err = pdfg.WriteFile(file)
	if err != nil {
		panic(err)
	}
}
