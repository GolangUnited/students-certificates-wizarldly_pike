package local

import (
	"os"
	"reflect"
	"testing"
)

var currentWorkDir string

func generateTestData(testTemplatesDir, testCertificatesDir string) (*localStorage, error) {
	err := os.Setenv(envTemplatesDir, testTemplatesDir)
	if err != nil {
		return nil, err
	}

	err = os.Setenv(envCertificatesDir, testCertificatesDir)
	if err != nil {
		return nil, err
	}
	return New()
}

func TestMain(m *testing.M) {
	currentDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	currentWorkDir = currentDir

	templatesDir, existsEnvTemplatesDir := os.LookupEnv(envTemplatesDir)
	certificatesDir, existsEnvCertificatesDir := os.LookupEnv(envCertificatesDir)

	code := m.Run()

	if existsEnvTemplatesDir {
		os.Setenv(envTemplatesDir, templatesDir)
	} else {
		os.Unsetenv(envTemplatesDir)
	}

	if existsEnvCertificatesDir {
		os.Setenv(envCertificatesDir, certificatesDir)
	} else {
		os.Unsetenv(envCertificatesDir)
	}

	os.Exit(code)
}

func TestNew_envTemplatesDir_fail(t *testing.T) {
	testTemplatesDir := ""
	testCertificatesDir := currentWorkDir
	_, err := generateTestData(testTemplatesDir, testCertificatesDir)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestNew_envCertificatesDir_fail(t *testing.T) {
	testTemplatesDir := currentWorkDir
	testCertificatesDir := ""
	_, err := generateTestData(testTemplatesDir, testCertificatesDir)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestNew_envDirs(t *testing.T) {
	testTemplatesDir := currentWorkDir
	testCertificatesDir := currentWorkDir
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
	testData, err := generateTestData(currentWorkDir, currentWorkDir)
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
		t.Error(err)
	}
}

func TestCertificatesOperations(t *testing.T) {
	testData, err := generateTestData(currentWorkDir, currentWorkDir)
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
		t.Error(err)
	}
}
