package certgenerator

import (
	"os"
	"reflect"
	"testing"
)

var testData = struct {
	certGenerator CertGenerator
	goodTemplate  []byte
	failTemplate  []byte
	expected      []byte
}{
	certGenerator: CertGenerator{},
	goodTemplate: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName}}</p><p>{{.CourseType}}</p><p>{{.CourseHours}}</p><p>{{.CourseDate}}</p>
	<p>{{.CourseMentors}}</p><p>{{.StudentFirstname}}</p><p>{{.StudentLastname}}</p>
	</body></html>`),
	failTemplate: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName_Fail}}</p></body></html>`),
	expected: []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>Golang</p><p>Theory</p><p>35</p><p>25.01.2023</p>
	<p>Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling</p><p>Ivan</p><p>Ivanov</p>
	</body></html>`),
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

	if !reflect.DeepEqual(gotCertif, testData.expected) {
		t.Errorf("%q and %q should be equal", "gotSertif", "expectedCert")
	}
}
