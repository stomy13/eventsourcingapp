package student

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/samber/lo"
)

type IEvent interface {
	StreamId() StudentId
	Sk() string
	apply(*Student)
	Json() (string, error)
}

type Event struct {
	Type         string
	CreatedAtUtc time.Time
}

func NewEvent(eventType string) Event {
	now := time.Now().UTC()
	return Event{
		Type:         eventType,
		CreatedAtUtc: now,
	}
}

func NewEventFromJson(jsonStr string) (Event, error) {
	event := Event{}
	err := json.Unmarshal([]byte(jsonStr), &event)
	return event, err
}

func (e Event) Sk() string {
	return e.CreatedAtUtc.Format(time.RFC3339Nano)
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

func (e StudentCreated) Json() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
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

func (e StudentUpdated) Json() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
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

func (e StudentDeleted) Json() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
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

func (e StudentEnrolled) Json() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
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

func (e StudentUnEnrolled) Json() (string, error) {
	b, err := json.Marshal(e)
	return string(b), err
}
