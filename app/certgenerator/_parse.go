package certgenerator

import (
	"bytes"
	"html/template"
	"path"
	"path/filepath"
)

func ParseTemplate(p PdfData) (*bytes.Buffer, error) {
	if err := p.Validate(); err != nil {
		return nil, errors.Wrap()
	}
	absPath, _ := filepath.Abs("../utils/st  orage/Ϗ local/Templates/")
	path.Join()
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
