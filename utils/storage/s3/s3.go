package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gus_certificates/utils/storage"
	"io"
	"log"
	"net/url"
	"time"
)

type minioStorage struct {
	client *minio.Client
	//logger
}

func NewStorage(endpoint, accessKey, secretKey string) (storage.Storage, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("Error!")
	}

	return &minioStorage{
		client: minioClient,
	}, nil
}

func (m minioStorage) GetTemplate(templateName string) ([]byte, error) {
	ctx := context.Background()
	bucketName := "templates"
	objectName := templateName

	obj, err := m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("Template not getted!")
	}

	log.Printf("Successfully downladed %s", objectName)

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("Template info not returned!")
	}

	buffer := make([]byte, objectInfo.Size)

	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Template not readed!")
	}

	return buffer, nil
}

func (m minioStorage) GetCertificate(certificateName string) ([]byte, error) {
	ctx := context.Background()
	bucketName := "certificates"
	objectName := certificateName

	obj, err := m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("Certificate not getted!")
	}

	log.Printf("Successfully downladed %s", objectName)

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("Certificate info not returned!")
	}

	buffer := make([]byte, objectInfo.Size)

	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("Certificate not readed!")
	}
	return buffer, nil
}

func (m minioStorage) GetCertificatePath(certificateName string) (string, error) {
	ctx := context.Background()
	bucketName := "certificates"
	objectName := certificateName

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	url, err := m.client.PresignedGetObject(ctx, bucketName, objectName, time.Second*60, reqParams)
	if err != nil {
		return "", fmt.Errorf("URL not generated!")
	}

	log.Printf("URL generated!")

	return url.String(), nil
}

func (m minioStorage) SaveTemplate(templateName string, data []byte) error {
	ctx := context.Background()
	bucketName := "templates"
	objectName := templateName

	_, err := m.client.PutObject(ctx, bucketName, objectName, bytes.NewReader(data), -1, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("Template not saved!")
	}

	log.Printf("Successfully saved %s", objectName)

	return nil
}

func (m minioStorage) SaveCertificate(certificateName string, data []byte) error {
	ctx := context.Background()
	bucketName := "certificates"
	objectName := certificateName

	_, err := m.client.PutObject(ctx, bucketName, objectName, bytes.NewReader(data), -1, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("Certificate not saved!")
	}

	log.Printf("Successfully saved %s", objectName)

	return nil
}

func (m minioStorage) DeleteTemplate(templateName string) error {
	ctx := context.Background()
	bucketName := "templates"
	objectName := templateName

	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("Template not removed!")
	}

	log.Printf("Successfully removed %s", objectName)

	return nil
}

func (m minioStorage) DeleteCertificate(certificateName string) error {
	ctx := context.Background()
	bucketName := "certificates"
	objectName := certificateName

	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("Certificate not removed!")
	}

	log.Printf("Successfully removed %s", objectName)

	return nil
}
