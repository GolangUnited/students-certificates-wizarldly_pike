package local

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
)

const (
	envTemplatesDir    = "TEMPLATES_DIR"
	envCertificatesDir = "CERTIFICATES_DIR"
)

type localStorage struct {
	templatesDir    string
	certificatesDir string
}

func New() (*localStorage, error) {
	templatesDir := os.Getenv(envTemplatesDir)
	if templatesDir == "" {
		return nil, fmt.Errorf("environment variable %q not set", envTemplatesDir)
	}
	err := os.MkdirAll(templatesDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	certificatesDir := os.Getenv(envCertificatesDir)
	if certificatesDir == "" {
		return nil, fmt.Errorf("environment variable %q not set", envCertificatesDir)
	}
	err = os.MkdirAll(certificatesDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	ls := &localStorage{}
	ls.templatesDir = templatesDir
	ls.certificatesDir = certificatesDir

	return ls, nil
}

func (l *localStorage) GetTemplate(fileName string) ([]byte, error) {
	fullPath := path.Join(l.templatesDir, fileName)
	return getFile(fullPath)

}

func (l *localStorage) GetCertificate(fileName string) ([]byte, error) {
	fullPath := path.Join(l.certificatesDir, fileName)
	return getFile(fullPath)
}

func (l *localStorage) GetCertificatePath(fileName string) (string, error) {
	fullPath := path.Join(l.certificatesDir, fileName)

	if _, err := os.Stat(fullPath); errors.Is(err, fs.ErrNotExist) {
		return "", err
	}
	return fullPath, nil
}

func (l *localStorage) SaveTemplate(fileName string, data []byte) error {
	fullPath := path.Join(l.templatesDir, fileName)
	return saveFile(fullPath, data)

}

func (l *localStorage) SaveCertificate(fileName string, data []byte) error {
	fullPath := path.Join(l.certificatesDir, fileName)
	return saveFile(fullPath, data)
}

func (l *localStorage) DeleteTemplate(fileName string) error {
	fullPath := path.Join(l.templatesDir, fileName)
	return deleteFile(fullPath)

}

func (l *localStorage) DeleteCertificate(fileName string) error {
	fullPath := path.Join(l.certificatesDir, fileName)
	return deleteFile(fullPath)
}

func getFile(fullPath string) ([]byte, error) {
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func saveFile(fullPath string, data []byte) error {
	err := os.WriteFile(fullPath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(fullPath string) error {
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}