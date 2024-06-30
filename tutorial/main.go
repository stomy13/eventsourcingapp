package main

import (
	"fmt"
	"time"
	"tutorial/student"

	"github.com/google/uuid"
)

func main() {
	studentDatabase := student.NewDatabase()

	studentId := uuid.NewString()
	studentCreated := student.StudentCreated{
		StudentId:   student.StudentId(studentId),
		FullName:    "John Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentCreated)

	studentEnrolled := student.StudentEnrolled{
		StudentId: student.StudentId(studentId),
		CourseId:  "course-1",
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentEnrolled)

	studentUpdated := student.StudentUpdated{
		StudentId: student.StudentId(studentId),
		Email:     "john.doe.new@example.com",
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentUpdated)

	student := studentDatabase.GetStudent(student.StudentId(studentId))
	fmt.Println(student)
}
