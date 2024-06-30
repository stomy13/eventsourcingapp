package student

type Database interface {
	Append(event IEvent)
	GetStudent(studentId StudentId) *Student
	GetStudentView(studentId StudentId) *Student
}
