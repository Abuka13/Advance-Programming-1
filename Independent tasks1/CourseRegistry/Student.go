package CourseRegistry

import (
	"fmt"
)

type Student struct {
	ID      uint64
	Name    string
	Courses []string
}

type Registry struct {
	Student map[uint64]Student
}

func (r *Registry) AddStudent(student Student) error {

	if student.Name == "" {
		return fmt.Errorf("student name cannot be empty")
	}
	if _, ok := r.Student[student.ID]; ok {
		return fmt.Errorf("id already exists")
	}

	r.Student[student.ID] = student
	return nil
}

func (r *Registry) EnrollCourse(studentID uint64, course string) error {
	student, ok := r.Student[studentID]
	if !ok {
		return fmt.Errorf("student not found")
	}
	if course == "" {
		return fmt.Errorf("course name cannot be empty")
	}
	for _, c := range student.Courses {
		if c == course {
			return fmt.Errorf("student already enrolled in the course")
		}
	}
	student.Courses = append(student.Courses, course)
	r.Student[studentID] = student
	return nil

}
func (r *Registry) RemoveCourses(studentID uint64, course string) error {
	student, ok := r.Student[studentID]
	if !ok {
		return fmt.Errorf("student not found")
	}
	flag := false
	var newCourses []string
	for _, c := range student.Courses {
		if c == course {
			flag = true
			continue
		}
		newCourses = append(newCourses, c)
	}
	if !flag {
		return fmt.Errorf("student not enrolled in the course")
	}
	student.Courses = newCourses
	r.Student[studentID] = student
	return nil
}
func (r *Registry) ListStudents() []Student {
	var s []Student
	for _, student := range r.Student {
		s = append(s, student)
	}
	return s
}

func (r *Registry) CoursesCount() map[string]int {
	n := make(map[string]int)
	for _, student := range r.Student {
		for _, course := range student.Courses {
			n[course]++
		}
	}
	return n
}
