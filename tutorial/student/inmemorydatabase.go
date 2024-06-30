package student

import "context"

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

	student := d.GetStudent(event.StreamId())
	if student != nil {
		d.students[event.StreamId()] = student
	}

	return nil
}

func (d *InMemoryDatabase) GetStudent(studentId StudentId) *Student {
	events := d.events[studentId]
	if len(events) == 0 {
		return nil
	}

	student := &Student{}
	for _, event := range events {
		student.Apply(event)
	}
	return student
}

func (d *InMemoryDatabase) GetStudentView(studentId StudentId) *Student {
	return d.students[studentId]
}
