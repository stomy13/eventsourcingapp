package student

import (
	"context"
	"errors"
)

var _ Database = (*InMemoryDatabase)(nil)

type InMemoryDatabase struct {
	events   map[StudentId][]IEvent
	students map[StudentId]*Student
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		events:   make(map[StudentId][]IEvent),
		students: make(map[StudentId]*Student),
	}
}

func (d *InMemoryDatabase) Append(ctx context.Context, event IEvent) error {
	events := d.events[event.StreamId()]
	d.events[event.StreamId()] = append(events, event)

	student, err := d.GetStudent(ctx, event.StreamId())
	if err != nil {
		return err
	}
	if student != nil {
		d.students[event.StreamId()] = student
	}

	return nil
}

func (d *InMemoryDatabase) GetStudent(ctx context.Context, studentId StudentId) (*Student, error) {
	events := d.events[studentId]
	if len(events) == 0 {
		return nil, errors.New("student not found")
	}

	student := &Student{}
	for _, event := range events {
		student.Apply(event)
	}
	return student, nil
}

func (d *InMemoryDatabase) GetStudentView(ctx context.Context, studentId StudentId) (*Student, error) {
	student := d.students[studentId]
	if student == nil {
		return nil, errors.New("student not found")
	}
	return student, nil
}
