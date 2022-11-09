package certdata

import (
	"strings"
)

type Data struct {
	courseName       string
	courseType       string
	courseHours      string
	courseDate       string
	courseMentors    []string
	studentFirstname string
	studentLastname  string
}

func New() *Data {
	return &Data{}
}

func (c *Data) GetDataForTemplate() any {
	returnData := struct {
		CourseName       string
		CourseType       string
		CourseHours      string
		CourseDate       string
		CourseMentors    string
		StudentFirstname string
		StudentLastname  string
	}{
		CourseName:       c.courseName,
		CourseType:       c.courseType,
		CourseHours:      c.courseHours,
		CourseDate:       c.courseDate,
		CourseMentors:    strings.Join(c.courseMentors, ", "),
		StudentFirstname: c.studentFirstname,
		StudentLastname:  c.studentLastname,
	}

	return returnData
}

func (c *Data) GetDataForIDGenerator() string {
	return c.courseName + c.courseType + c.courseHours + c.courseDate + c.studentFirstname + c.studentLastname
}

func (c *Data) SetCourseName(courseName string) {
	c.courseName = courseName
}

func (c *Data) SetCourseType(courseType string) {
	c.courseType = courseType
}

func (c *Data) SetCourseHours(courseHours string) {
	c.courseHours = courseHours
}

func (c *Data) SetCourseDate(courseDate string) {
	c.courseDate = courseDate
}

func (c *Data) SetCourseMentors(courseMentors []string) {
	c.courseMentors = courseMentors
}

func (c *Data) SetStudentFirstname(studentFirstname string) {
	c.studentFirstname = studentFirstname
}

func (c *Data) SetStudentLastname(studentLastname string) {
	c.studentLastname = studentLastname
}
