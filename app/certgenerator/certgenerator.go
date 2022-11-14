package certgenerator

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html/template"
	"strings"

	valid "github.com/go-ozzo/ozzo-validation/v4"
)

var shortTextFieldRule = []valid.Rule{valid.Required, valid.RuneLength(1, 50)}
var longTextFieldRule = []valid.Rule{valid.Required, valid.RuneLength(1, 250)}

type CertGenerator struct {
	data certData
}

type certData struct {
	CourseName       string
	CourseType       string
	CourseHours      string
	CourseDate       string
	CourseMentors    string
	StudentFirstname string
	StudentLastname  string
	QrCodeLink       template.URL
}

func (c *CertGenerator) ValidateData() error {
	return valid.ValidateStruct(&c.data,
		valid.Field(&c.data.CourseName, longTextFieldRule...),
		valid.Field(&c.data.CourseType, shortTextFieldRule...),
		valid.Field(&c.data.CourseHours, shortTextFieldRule...),
		valid.Field(&c.data.CourseDate, shortTextFieldRule...),
		valid.Field(&c.data.CourseMentors, longTextFieldRule...),
		valid.Field(&c.data.StudentFirstname, shortTextFieldRule...),
		valid.Field(&c.data.StudentLastname, shortTextFieldRule...),
	)
}

func (c *CertGenerator) GenerateCertHTML(templateHTMLData []byte) ([]byte, error) {
	tmpl := template.New("tmpl")

	tmpl, err := tmpl.Parse(string(templateHTMLData))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, c.data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *CertGenerator) CheckTemplateHTML(templateHTMLData []byte) error {
	_, err := c.GenerateCertHTML(templateHTMLData)

	return err
}

func (c *CertGenerator) GenerateID() string {
	data := []byte(c.getDataForIDGenerator())
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (c *CertGenerator) getDataForIDGenerator() string {
	return fmt.Sprintf("%s%s%s%s%s%s", c.data.CourseName, c.data.CourseType, c.data.CourseHours,
		c.data.CourseDate, c.data.StudentFirstname, c.data.StudentLastname)
}

func (c *CertGenerator) SetCourseName(courseName string) {
	c.data.CourseName = courseName
}

func (c *CertGenerator) SetCourseType(courseType string) {
	c.data.CourseType = courseType
}

func (c *CertGenerator) SetCourseHours(courseHours string) {
	c.data.CourseHours = courseHours
}

func (c *CertGenerator) SetCourseDate(courseDate string) {
	c.data.CourseDate = courseDate
}

func (c *CertGenerator) SetCourseMentors(courseMentors []string) {
	c.data.CourseMentors = strings.Join(courseMentors, ", ")
}

func (c *CertGenerator) SetStudentFirstname(studentFirstname string) {
	c.data.StudentFirstname = studentFirstname
}

func (c *CertGenerator) SetStudentLastname(studentLastname string) {
	c.data.StudentLastname = studentLastname
}

func (c *CertGenerator) SetQrCodeLink(imgData []byte) {
	htmlTagImagePngBase64 := "data:image/png;base64,"
	imgBase64String := base64.StdEncoding.EncodeToString(imgData)

	c.data.QrCodeLink = template.URL(fmt.Sprintf("%s%s", htmlTagImagePngBase64, imgBase64String))
}
