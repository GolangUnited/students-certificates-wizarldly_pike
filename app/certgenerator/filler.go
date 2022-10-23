package certgenerator

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func ParseTemplate(p PdfData) (*bytes.Buffer, error) {
	absPath, _ := filepath.Abs("../utils/storage/local/Templates/")
	tmpl, err := template.ParseFiles(absPath + "/" + p.Template)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
