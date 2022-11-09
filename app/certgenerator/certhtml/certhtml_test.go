package certhtml

import (
	"gus_certificates/app/types/certdata"
	"os"
	"reflect"
	"testing"
)

var certGenerator *certhtml
var dataForCert *certdata.Data
var templateCorrect []byte
var expectedCert []byte
var templateFail []byte

func TestMain(m *testing.M) {
	certGenerator = New()
	templateCorrect = []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName}}</p><p>{{.CourseType}}</p><p>{{.CourseHours}}</p><p>{{.CourseDate}}</p>
	<p>{{.CourseMentors}}</p><p>{{.StudentFirstname}}</p><p>{{.StudentLastname}}</p>
	</body></html>`)
	expectedCert = []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>Golang</p><p>Theory</p><p>35</p><p>25.01.2023</p>
	<p>Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling</p><p>Ivan</p><p>Ivanov</p>
	</body></html>`)
	templateFail = []byte(`<html><body><h1 style="color:red;">Test html color<h1>
	<p>{{.CourseName_Fail}}</p></body></html>`)

	dataForCert = certdata.New()
	dataForCert.SetCourseName("Golang")
	dataForCert.SetCourseType("Theory")
	dataForCert.SetCourseHours("35")
	dataForCert.SetCourseDate("25.01.2023")
	dataForCert.SetCourseMentors([]string{"Pavel Gordiyanov", "Mikita Viarbovikau", "Sergey Shtripling"})
	dataForCert.SetStudentFirstname("Ivan")
	dataForCert.SetStudentLastname("Ivanov")

	os.Exit(m.Run())
}

func TestGenerateCertificate_fail(t *testing.T) {
	_, err := certGenerator.GenerateCertificate(dataForCert, templateFail)
	if err == nil {
		t.Error("err must not be nil")
	}
}

func TestGenerateCertificate(t *testing.T) {
	gotCertif, err := certGenerator.GenerateCertificate(dataForCert, templateCorrect)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(gotCertif, expectedCert) {
		t.Errorf("%q and %q should be equal", "gotSertif", "expectedCert")
	}
}
