package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"tutorial/student"

	"github.com/google/uuid"
)

func main() {
	useDynamoDBDatabase()
}

func useDynamoDBDatabase() {
	ctx := context.Background()
	studentDatabase := student.NewDynamoDBDatabase(ctx)

	studentId := student.StudentId(uuid.NewString())
	studentCreated := student.StudentCreated{
		StudentId:   studentId,
		FullName:    "John Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Event:       student.NewEvent("StudentCreated"),
	}
	studentDatabase.Append(ctx, studentCreated)

	studentEnrolled := student.StudentEnrolled{
		StudentId: studentId,
		CourseId:  "course-1",
		Event:     student.NewEvent("StudentEnrolled"),
	}
	studentDatabase.Append(ctx, studentEnrolled)

	studentUpdated := student.StudentUpdated{
		StudentId: studentId,
		Email:     "john.doe.new@example.com",
		Event:     student.NewEvent("StudentUpdated"),
	}
	studentDatabase.Append(ctx, studentUpdated)
}

func useInMemoryDatabase() {
	ctx := context.Background()
	studentDatabase := student.NewInMemoryDatabase()

	studentId := student.StudentId(uuid.NewString())
	studentCreated := student.StudentCreated{
		StudentId:   studentId,
		FullName:    "John Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Event:       student.NewEvent("StudentCreated"),
	}
	studentDatabase.Append(ctx, studentCreated)

	studentEnrolled := student.StudentEnrolled{
		StudentId: studentId,
		CourseId:  "course-1",
		Event:     student.NewEvent("StudentEnrolled"),
	}
	studentDatabase.Append(ctx, studentEnrolled)

	studentUpdated := student.StudentUpdated{
		StudentId: studentId,
		Email:     "john.doe.new@example.com",
		Event:     student.NewEvent("StudentUpdated"),
	}
	studentDatabase.Append(ctx, studentUpdated)

	student, err := studentDatabase.GetStudent(ctx, studentId)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(student)

	studentView, err := studentDatabase.GetStudentView(ctx, studentId)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(studentView)
}
