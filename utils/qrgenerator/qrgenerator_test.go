package qrgenerator

import "testing"

func TestGenerateQrPNG(t *testing.T) {
	qrGen := QrGenerator{}

	testStringLink := "https://example.com/certificate1234.pdf"

	data, err := qrGen.GenerateQrPNG(testStringLink)
	if err != nil {
		t.Error(err)
	}

	if len(data) == 0 {
		t.Error("len(data): must be greater than 0")
	}
}
