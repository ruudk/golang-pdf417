package main

import (
	"fmt"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/barcode"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	barcode.Barcode(
		pdf,
		barcode.RegisterPdf417(pdf, "Hello world", 2, 2),
		20,
		0,
		150,
		30,
		false,
	)

	err := pdf.OutputFileAndClose("hello.pdf")

	if err != nil {
		fmt.Println("Something went wrong")
	}
}
