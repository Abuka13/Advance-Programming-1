package main

import (
	"Independent_tasks_1/CourseRegistry"
	"fmt"
	"os"
)

func main() {
	registry := CourseRegistry.Registry{
		Student: make(map[uint64]CourseRegistry.Student),
	}
	
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Student")
		fmt.Println("2. Enroll Course")
		fmt.Println("3. Remove Course")
		fmt.Println("4. List Students")
		fmt.Println("5. Course Statistics")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string
			fmt.Print("Enter ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			err := registry.AddStudent(CourseRegistry.Student{ID: id, Name: name})
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			var id uint64
			var course string
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Course Name: ")
			fmt.Scan(&course)
			err := registry.EnrollCourse(id, course)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			var id uint64
			var course string
			fmt.Print("Enter Student ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Course Name: ")
			fmt.Scan(&course)
			err := registry.RemoveCourses(id, course)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			students := registry.ListStudents()
			for _, s := range students {
				fmt.Printf("ID: %d | Name: %s | Courses: %v\n", s.ID, s.Name, s.Courses)
			}

		case 5:
			stats := registry.CoursesCount()
			for course, count := range stats {
				fmt.Printf("%s → %d\n", course, count)
			}

		case 6:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}
}
