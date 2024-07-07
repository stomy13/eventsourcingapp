package student

import (
	"encoding/json"
	"time"
)

type Student struct {
	Pk string
	Sk string

	Id          StudentId
	FullName    string
	Email       string
	DateOfBirth time.Time
	CoursesIds  []string
}

func NewStudentFromJson(jsonString string) (*Student, error) {
	var student Student
	err := json.Unmarshal([]byte(jsonString), &student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *Student) Apply(event IEvent) {
	event.apply(s)
}

func (s *Student) Json() (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type StudentId string

func (s StudentId) String() string {
	return string(s)
}
