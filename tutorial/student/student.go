package student

import "time"

type Student struct {
	Id          StudentId
	FullName    string
	Email       string
	DateOfBirth time.Time
	CoursesIds  []string
}

type StudentId string

func (s StudentId) String() string {
	return string(s)
}

func (s *Student) Apply(event IEvent) {
	event.apply(s)
}
