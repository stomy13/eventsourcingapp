package student

import "context"

type Database interface {
	Append(ctx context.Context, event IEvent) error
	GetStudent(studentId StudentId) *Student
	GetStudentView(studentId StudentId) *Student
}
