package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	envTemplatesDir    = "TEMPLATES_DIR"
	envCertificatesDir = "CERTIFICATES_DIR"
	envS3BucketName    = "S3_BUCKET_NAME"

	defaultBucketRegion = "eu-central-1"
)

type s3Storage struct {
	templatesDir    string
	certificatesDir string
	s3BucketName    string
	s3Client        *s3.Client
}

func New() (*s3Storage, error) {
	templatesDir := os.Getenv(envTemplatesDir)
	if templatesDir == "" {
		return nil, fmt.Errorf("environment variable %q not set", envTemplatesDir)
	}

	certificatesDir := os.Getenv(envCertificatesDir)
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

	return s3s, nil
}

func newS3Client(bucket string) (*s3.Client, error) {
	region, err := finBucketRegion(bucket)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	// Проверка доступности bucket.
	_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucket)})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func finBucketRegion(bucket string) (string, error) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(defaultBucketRegion))
	if err != nil {
		return "", err
	}

	// Поиск региона содержащий наш bucket.
	region, err := manager.GetBucketRegion(ctx, s3.NewFromConfig(cfg), bucket)
	if err != nil {
		var bnf manager.BucketNotFound
		if errors.As(err, &bnf) {
			return "", fmt.Errorf("unable to find bucket:%q Region", bucket)
		}
		return "", err
	}
	return region, nil
}

func (s *s3Storage) GetTemplate(fileName string) ([]byte, error) {
	fullPath := path.Join(s.templatesDir, fileName)
	return s.getFile(fullPath)

}

func (s *s3Storage) GetCertificate(fileName string) ([]byte, error) {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.getFile(fullPath)
}

// func (s *s3Storage) GetCertificatePath(fileName string) (string, error) {
// 	fullPath := path.Join(s.certificatesDir, fileName)

// 	if _, err := os.Stat(fullPath); errors.Is(err, fs.ErrNotExist) {
// 		return "", err
// 	}
// 	return fullPath, nil
// }

func (s *s3Storage) SaveTemplate(fileName string, data []byte) error {
	fullPath := path.Join(s.templatesDir, fileName)
	return s.saveFile(fullPath, data)

}

func (s *s3Storage) SaveCertificate(fileName string, data []byte) error {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.saveFile(fullPath, data)
}

func (s *s3Storage) DeleteTemplate(fileName string) error {
	fullPath := path.Join(s.templatesDir, fileName)
	return s.deleteFile(fullPath)

}

func (s *s3Storage) DeleteCertificate(fileName string) error {
	fullPath := path.Join(s.certificatesDir, fileName)
	return s.deleteFile(fullPath)
}

func (s *s3Storage) getFile(fullPath string) ([]byte, error) {
	ctx := context.TODO()
	obj := &s3.GetObjectInput{Bucket: aws.String(s.s3BucketName), Key: aws.String(fullPath)}
	resp, err := s.s3Client.GetObject(ctx, obj)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *s3Storage) saveFile(fullPath string, data []byte) error {
	ctx := context.TODO()
	obj := &s3.PutObjectInput{Bucket: aws.String(s.s3BucketName), Key: aws.String(fullPath), Body: bytes.NewBuffer(data)}
	_, err := s.s3Client.PutObject(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *s3Storage) deleteFile(fullPath string) error {
	if err := s.checkFile(fullPath); err != nil {
		return err
	}

	ctx := context.TODO()
	obj := &s3.DeleteObjectInput{Bucket: aws.String(s.s3BucketName), Key: aws.String(fullPath)}
	_, err := s.s3Client.DeleteObject(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *s3Storage) checkFile(fullPath string) error {
	ctx := context.TODO()
	obj := &s3.HeadObjectInput{Bucket: aws.String(s.s3BucketName), Key: aws.String(fullPath)}
	_, err := s.s3Client.HeadObject(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}
