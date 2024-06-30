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
		StudentID:   student.StudentID(studentId),
		FullName:    "John Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentCreated)

	studentEnrolled := student.StudentEnrolled{
		StudentID: student.StudentID(studentId),
		CourseID:  "course-1",
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentEnrolled)

	studentUpdated := student.StudentUpdated{
		StudentID: student.StudentID(studentId),
		Email:     "john.doe.new@example.com",
		Event: student.Event{
			CreatedAtUtc: time.Now().UTC(),
		},
	}
	studentDatabase.Append(studentUpdated)

	fmt.Println(studentDatabase)
}
