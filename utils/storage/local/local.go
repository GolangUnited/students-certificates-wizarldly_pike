package local

import (
	"fmt"
	"os"
	"path"
)

const (
	envTemplatesDir    = "TEMPLATES_DIR"
	envCertificatesDir = "CERTIFICATES_DIR"
)

// Оптимизация скорости. Структура для сохранения в памяти последнего запрошенного шаблона, возращает при повторных запросах.
type lastRequestTemplate struct {
	name string
	data []byte
}

type localStorage struct {
	templatesDir    string
	certificatesDir string
	lastTemplate    lastRequestTemplate
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
	ls.lastTemplate = lastRequestTemplate{}

	return ls, nil
}

func (l *localStorage) GetTemplate(fileName string) ([]byte, error) {
	// Возвращаем из памяти если уже запрашивали.
	if l.lastTemplate.name == fileName {
		return l.lastTemplate.data, nil
	}

	fullPath := path.Join(l.templatesDir, fileName)
	data, err := getFile(fullPath)
	if err != nil {
		return nil, err
	}

	// Обновляем в памяти последний запрошенный.
	l.lastTemplate.name = fileName
	l.lastTemplate.data = data

	return data, nil
}

func (l *localStorage) GetCertificate(fileName string) ([]byte, error) {
	fullPath := path.Join(l.certificatesDir, fileName)
	return getFile(fullPath)
}

func (l *localStorage) SaveTemplate(fileName string, data []byte) error {
	fullPath := path.Join(l.templatesDir, fileName)
	err := saveFile(fullPath, data)
	if err != nil {
		return err
	}

	// Обновляем в памяти последний сохраненный при обновлении.
	if l.lastTemplate.name == fileName {
		l.lastTemplate.data = data
	}

	return nil
}

func (l *localStorage) SaveCertificate(fileName string, data []byte) error {
	fullPath := path.Join(l.certificatesDir, fileName)
	return saveFile(fullPath, data)
}

func (l *localStorage) DeleteTemplate(fileName string) error {
	// Очищаем в памяти последний сохраненный при его удалении.
	if l.lastTemplate.name == fileName {
		l.lastTemplate.name = ""
		l.lastTemplate.data = nil
	}

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
