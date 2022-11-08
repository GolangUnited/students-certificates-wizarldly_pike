package certgenerator

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Save(buf *bytes.Buffer) error {
	absPath, _ := filepath.Abs("../utils/storage/local/PDF/example.pdf")
	data, _ := ioutil.ReadAll(buf)
	err := os.WriteFile(absPath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
