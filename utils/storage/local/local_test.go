package local

import (
	"os"
	"reflect"
	"testing"
)

func generateTestData(testTemplatesDir, testCertificatesDir string) (*localStorage, error) {

	templatesDir, ok := os.LookupEnv(envTemplatesDir)
	if ok {
		defer os.Setenv(envTemplatesDir, templatesDir)
	} else {
		defer os.Unsetenv(envTemplatesDir)
	}

	err := os.Setenv(envTemplatesDir, testTemplatesDir)
	if err != nil {
		return nil, err
	}

	certificatesDir, ok := os.LookupEnv(envCertificatesDir)
	if ok {
		defer os.Setenv(envCertificatesDir, certificatesDir)
	} else {
		defer os.Unsetenv(envCertificatesDir)
	}

	err = os.Setenv(envCertificatesDir, testCertificatesDir)
	if err != nil {
		return nil, err
	}
	return New()
}

func TestNew_envTemplatesDir_fail(t *testing.T) {
	testTemplatesDir := ""
	testCertificatesDir := "/test/Certificates"
	_, err := generateTestData(testTemplatesDir, testCertificatesDir)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestNew_envCertificatesDir_fail(t *testing.T) {
	testTemplatesDir := "/test/Templates"
	testCertificatesDir := ""
	_, err := generateTestData(testTemplatesDir, testCertificatesDir)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestNew_envDirs(t *testing.T) {
	testTemplatesDir := "/test/Templates"
	testCertificatesDir := "/test/Certificates"
	testData, err := generateTestData(testTemplatesDir, testCertificatesDir)
	if err != nil {
		t.Fatal(err)
	}

	if testData.templatesDir != testTemplatesDir {
		t.Errorf("templatesDir expected:%q, actual:%q", testTemplatesDir, testData.templatesDir)
	}

	if testData.certificatesDir != testCertificatesDir {
		t.Errorf("certificatesDir expected:%q, actual:%q", testCertificatesDir, testData.certificatesDir)
	}
}

func TestTemplatesOperations(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testData, err := generateTestData(currentDir, currentDir)
	if err != nil {
		t.Fatal(err)
	}

	testFileName := "testDataTemplates.tmp"
	testBytes := []byte("Test Templates Operations")
	err = testData.SaveTemplate(testFileName, testBytes)
	if err != nil {
		t.Fatal(err)
	}

	readBytes, err := testData.GetTemplate(testFileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testBytes, readBytes) {
		t.Error("testBytes is not equal to readBytes")
	}

	err = testData.DeleteTemplate(testFileName)
	if err != nil {
		t.Error(err, currentDir)
	}
}

func TestCertificatesOperations(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testData, err := generateTestData(currentDir, currentDir)
	if err != nil {
		t.Fatal(err)
	}

	testFileName := "testDataCertificates.tmp"
	testBytes := []byte("Test Certificates Operations")
	err = testData.SaveCertificate(testFileName, testBytes)
	if err != nil {
		t.Fatal(err)
	}

	readBytes, err := testData.GetCertificate(testFileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testBytes, readBytes) {
		t.Error("testBytes is not equal to readBytes")
	}

	err = testData.DeleteCertificate(testFileName)
	if err != nil {
		t.Error(err, currentDir)
	}
}
