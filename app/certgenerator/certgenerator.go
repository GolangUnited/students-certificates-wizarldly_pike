package certgenerator

import (
	"bytes"
	"html/template"
	"strings"
)

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

// func (c *CertGenerator) getDataForIDGenerator() string {
// 	return fmt.Sprintf("%s%s%s%s%s%s", c.data.CourseName, c.data.CourseType, c.data.CourseHours,
// 		c.data.CourseDate, c.data.StudentFirstname, c.data.StudentLastname)
// }

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
