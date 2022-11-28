package qrgenerator

import "github.com/skip2/go-qrcode"

const sizeImage = 128

type QrGenerator struct {
}

func (q *QrGenerator) GenerateQrPNG(sourceString string) ([]byte, error) {
	data, err := qrcode.Encode(sourceString, qrcode.Medium, sizeImage)
	if err != nil {
		return nil, err
	}
	return data, nil
}
