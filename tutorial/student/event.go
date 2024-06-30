package student

import (
	"slices"
	"time"

	"github.com/samber/lo"
)

type IEvent interface {
	StreamId() StudentId
	apply(*Student)
}

type Event struct {
	CreatedAtUtc time.Time
}

type StudentCreated struct {
	Event
	StudentId   StudentId
	FullName    string
	Email       string
	DateOfBirth time.Time
}

func (e StudentCreated) StreamId() StudentId {
	return e.StudentId
}

func (e StudentCreated) apply(s *Student) {
	s.Id = e.StudentId
	s.FullName = e.FullName
	s.Email = e.Email
	s.DateOfBirth = e.DateOfBirth
}

type StudentUpdated struct {
	Event
	StudentId   StudentId
	FullName    string
	Email       string
	DateOfBirth time.Time
}

func (e StudentUpdated) StreamId() StudentId {
	return e.StudentId
}

func (e StudentUpdated) apply(s *Student) {
	if e.FullName != "" {
		s.FullName = e.FullName
	}
	if e.Email != "" {
		s.Email = e.Email
	}
	if !e.DateOfBirth.IsZero() {
		s.DateOfBirth = e.DateOfBirth
	}
}

type StudentDeleted struct {
	Event
	StudentId StudentId
}

func (e StudentDeleted) StreamId() StudentId {
	return e.StudentId
}

func (e StudentDeleted) apply(s *Student) {
	// TBI
}

type StudentEnrolled struct {
	Event
	StudentId StudentId
	CourseId  string
}

func (e StudentEnrolled) StreamId() StudentId {
	return e.StudentId
}

func (e StudentEnrolled) apply(s *Student) {
	if !slices.Contains(s.CoursesIds, e.CourseId) {
		s.CoursesIds = append(s.CoursesIds, e.CourseId)
	}
}

type StudentUnEnrolled struct {
	Event
	StudentId StudentId
	CourseId  string
}

func (e StudentUnEnrolled) StreamId() StudentId {
	return e.StudentId
}

func (e StudentUnEnrolled) apply(s *Student) {
	s.CoursesIds = lo.Filter(s.CoursesIds, func(courseId string, index int) bool {
		return courseId != e.CourseId
	})
}
