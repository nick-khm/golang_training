package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/certificate/cert"
)

type PdfSaver struct {
	Output string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		Output: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// background
	background(pdf)

	header(pdf, &cert)
	pdf.Ln(30)

	// body
	pdf.SetFont("Helvetica", "I", 30)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// body - student name
	pdf.SetFont("Times", "B", 50)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 25)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	footer(pdf, &cert)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.Output, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/bg.png", 0, 0, pageWidth, pageHeight, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename, x+margin, 20, imageWidth, 0, false, opts, 0, "")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename, x-margin, 20, imageWidth, 0, false, opts, 0, "")
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	filename := "img/certified.png"
	imageWidth := 70.0
	imageHeight := 70.0
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions(filename, pageWidth-imageWidth-40, pageHeight-imageHeight-10, imageWidth, imageHeight, false, opts, 0, "")
}
