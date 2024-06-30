package student

import "time"

type IEvent interface {
	StreamID() StudentID
}

type Event struct {
	CreatedAtUtc time.Time
}

type StudentCreated struct {
	Event
	StudentID   StudentID
	FullName    string
	Email       string
	DateOfBirth time.Time
}

func (e StudentCreated) StreamID() StudentID {
	return e.StudentID
}

type StudentUpdated struct {
	Event
	StudentID   StudentID
	FullName    string
	Email       string
	DateOfBirth time.Time
}

func (e StudentUpdated) StreamID() StudentID {
	return e.StudentID
}

type StudentDeleted struct {
	Event
	StudentID StudentID
}

func (e StudentDeleted) StreamID() StudentID {
	return e.StudentID
}

type StudentEnrolled struct {
	Event
	StudentID StudentID
	CourseID  string
}

func (e StudentEnrolled) StreamID() StudentID {
	return e.StudentID
}

type StudentUnEnrolled struct {
	Event
	StudentID StudentID
	CourseID  string
}

func (e StudentUnEnrolled) StreamID() StudentID {
	return e.StudentID
}
