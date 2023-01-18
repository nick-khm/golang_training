package html

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"training.go/certificate/cert"
)

type HtmlSaver struct {
	Output string
}

func New(outputdir string) (*HtmlSaver, error) {
	var p *HtmlSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &HtmlSaver{
		Output: outputdir,
	}
	return p, nil
}

func (p *HtmlSaver) Save(cert cert.Cert) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(p.Output, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, cert)
	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)
	return nil
}

var tpl = `
<!DOCTYPE html>
<html lang="en" data-layout="responsive">
  <head>
	<title>{{.LabelTitle}}</title>
  </head>
  <body>
	<h1>{{.LabelCompletion}}</h1>
		
  </body>
  </html>

`
