package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	envTemplatesDir    = "TEMPLATES_DIR"
	envCertificatesDir = "CERTIFICATES_DIR"
	envS3BucketName    = "S3_BUCKET_NAME"
	envS3Endpoint      = "S3_ENDPOINT"
	envAccessKeyID     = "ACCESS_KEY_ID"
	envSecretAccessKey = "SECRET_ACCESS_KEY"
)

// Оптимизация скорости. Структура для сохранения в памяти последнего запрошенного шаблона, возращает при повторных запросах.
type lastRequestTemplate struct {
	name string
	data []byte
}

type s3Storage struct {
	templatesDir    string
	certificatesDir string
	s3BucketName    string
	s3Client        *minio.Client
	lastTemplate    lastRequestTemplate
}

func New() (*s3Storage, error) {
	templatesDir := filepath.Base(os.Getenv(envTemplatesDir))
	if templatesDir == "" {
		return nil, fmt.Errorf("environment variable %q not set", envTemplatesDir)
	}

	certificatesDir := filepath.Base(os.Getenv(envCertificatesDir))
	if certificatesDir == "" {
		return nil, fmt.Errorf("environment variable %q not set", envCertificatesDir)
	}

	s3BucketName := os.Getenv(envS3BucketName)
	if s3BucketName == "" {
		return nil, fmt.Errorf("environment variable %q not set", envS3BucketName)
	}

	s3Client, err := newS3Client(s3BucketName)
	if err != nil {
		return nil, err
	}

	s3s := &s3Storage{}
	s3s.templatesDir = templatesDir
	s3s.certificatesDir = certificatesDir
	s3s.s3BucketName = s3BucketName
	s3s.s3Client = s3Client
	s3s.lastTemplate = lastRequestTemplate{}

	return s3s, nil
}

func newS3Client(bucket string) (*minio.Client, error) {

	s3Endpoint := os.Getenv(envS3Endpoint)
	if s3Endpoint == "" {
		return nil, fmt.Errorf("environment variable %q not set", envS3Endpoint)
	}

	accessKeyID := os.Getenv(envAccessKeyID)
	if accessKeyID == "" {
		return nil, fmt.Errorf("environment variable %q not set", envAccessKeyID)
	}

	secretAccessKey := os.Getenv(envSecretAccessKey)
	if secretAccessKey == "" {
		return nil, fmt.Errorf("environment variable %q not set", envSecretAccessKey)
	}

	useSSL := true
	options := &minio.Options{Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""), Secure: useSSL}
	client, err := minio.New(s3Endpoint, options)
	if err != nil {
		return nil, err
	}

	// Проверяем доступ к Bucket.
	ctx := context.Background()
	exist, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("bucket:%q not exist", bucket)
	}

	return client, nil
}

func (s *s3Storage) GetTemplate(fileName string) ([]byte, error) {
	// Возвращаем из памяти если уже запрашивали.
	if s.lastTemplate.name == fileName {
		return s.lastTemplate.data, nil
	}

	fullPath := path.Join(s.templatesDir, fileName)
	data, err := s.getFile(fullPath)
	if err != nil {
		return nil, err
	}

	// Обновляем в памяти последний запрошенный.
	s.lastTemplate.name = fileName
	s.lastTemplate.data = data

	return data, nil
}

func (s *s3Storage) GetCertificate(fileName string) ([]byte, error) {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.getFile(fullPath)
}

func (s *s3Storage) SaveTemplate(fileName string, data []byte) error {
	fullPath := path.Join(s.templatesDir, fileName)
	err := s.saveFile(fullPath, data)
	if err != nil {
		return err
	}

	// Обновляем в памяти последний сохраненный при обновлении.
	if s.lastTemplate.name == fileName {
		s.lastTemplate.data = data
	}

	return nil
}

func (s *s3Storage) SaveCertificate(fileName string, data []byte) error {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.saveFile(fullPath, data)
}

func (s *s3Storage) DeleteTemplate(fileName string) error {
	// Очищаем в памяти последний сохраненный при его удалении.
	if s.lastTemplate.name == fileName {
		s.lastTemplate.name = ""
		s.lastTemplate.data = nil
	}

	fullPath := path.Join(s.templatesDir, fileName)
	return s.deleteFile(fullPath)

}

func (s *s3Storage) DeleteCertificate(fileName string) error {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.deleteFile(fullPath)
}

func (s *s3Storage) getFile(fullPath string) ([]byte, error) {
	ctx := context.Background()

	obj, err := s.s3Client.GetObject(ctx, s.s3BucketName, fullPath, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer obj.Close()
	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *s3Storage) saveFile(fullPath string, data []byte) error {
	ctx := context.Background()
	_, err := s.s3Client.PutObject(ctx, s.s3BucketName, fullPath, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *s3Storage) deleteFile(fullPath string) error {
	// Проверяем существует/доступен ли файл.
	if err := s.checkFile(fullPath); err != nil {
		return err
	}

	ctx := context.Background()
	err := s.s3Client.RemoveObject(ctx, s.s3BucketName, fullPath, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *s3Storage) checkFile(fullPath string) error {
	ctx := context.Background()
	_, err := s.s3Client.StatObject(ctx, s.s3BucketName, fullPath, minio.StatObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
