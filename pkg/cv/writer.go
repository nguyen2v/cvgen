package cv

var (
	// WriterPDF is default writer
	WriterPDF = 0
)

type Writer interface {
	// WriteFile is writing file to selected format
	WriteFile(cv *CV, templateName, file string)
}

// GetWriter returns valid writer
func GetWriter(writerType int) Writer {
	switch writerType {
	case WriterPDF:
		return pdfWriter{}
	}

	return nil
}