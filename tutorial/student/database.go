package student

import "context"

type Database interface {
	Append(ctx context.Context, event IEvent) error
	GetStudent(ctx context.Context, studentId StudentId) (*Student, error)
	GetStudentView(ctx context.Context, studentId StudentId) (*Student, error)
}
