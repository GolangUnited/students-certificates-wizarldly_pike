package htmltopdf

import (
	"testing"
)

func TestNew_htmltopdf(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderHtmlToPdf(t *testing.T) {
	htp, err := New()
	if err != nil {
		t.Fatal(err)
	}

	testBytesHtml := []byte(`<html><body><h1 style="color:red;">Test html color<h1></body></html>`)
	_, err = htp.RenderHtmlToPdf(testBytesHtml)
	if err != nil {
		t.Error(err)
	}
}
