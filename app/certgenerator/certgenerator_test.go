package certgenerator

import (
	"os"
	"reflect"
	"testing"
)

var testData = struct {
	certGenerator *CertGenerator
	goodTemplate  []byte
	failTemplate  []byte
	expectedCert  []byte
	expectedId    string
}{
	certGenerator: &CertGenerator{},
	goodTemplate: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName}}</p><p>{{.CourseType}}</p><p>{{.CourseHours}}</p><p>{{.CourseDate}}</p>
	<p>{{.CourseMentors}}</p><p>{{.StudentFirstname}}</p><p>{{.StudentLastname}}</p>
	</body></html>`),
	failTemplate: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName_Fail}}</p></body></html>`),
	expectedCert: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>Golang</p><p>Theory</p><p>35</p><p>25.01.2023</p>
	<p>Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling</p><p>Ivan</p><p>Ivanov</p>
	</body></html>`),
	expectedId: "612364afe471b3b1cc80083183fd381d",
}

func TestMain(m *testing.M) {
	testData.certGenerator.SetCourseName("Golang")
	testData.certGenerator.SetCourseType("Theory")
	testData.certGenerator.SetCourseHours("35")
	testData.certGenerator.SetCourseDate("25.01.2023")
	testData.certGenerator.SetCourseMentors([]string{"Pavel Gordiyanov", "Mikita Viarbovikau", "Sergey Shtripling"})
	testData.certGenerator.SetStudentFirstname("Ivan")
	testData.certGenerator.SetStudentLastname("Ivanov")

	os.Exit(m.Run())
}

func TestGenerateCertHTML_fail(t *testing.T) {
	generator := testData.certGenerator
	_, err := generator.GenerateCertHTML(testData.failTemplate)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestGenerateCertificate(t *testing.T) {
	generator := testData.certGenerator
	gotCertif, err := generator.GenerateCertHTML(testData.goodTemplate)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gotCertif, testData.expectedCert) {
		t.Errorf("%q and %q should be equal", "gotSertif", "expectedCert")
	}
}

func TestGenerateID(t *testing.T) {
	generator := testData.certGenerator

	actualId := generator.GenerateID()
	if actualId != testData.expectedId {
		t.Errorf("expected:%q,  actual:%q", testData.expectedId, actualId)
	}
}

func TestCheckTemplateHTML_fail(t *testing.T) {
	generator := testData.certGenerator
	err := generator.CheckTemplateHTML(testData.failTemplate)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestCheckTemplateHTML(t *testing.T) {
	generator := testData.certGenerator
	err := generator.CheckTemplateHTML(testData.goodTemplate)
	if err != nil {
		t.Error(err)
	}
}
