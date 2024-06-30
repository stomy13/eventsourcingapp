package student

type Database struct {
	events map[StudentId][]IEvent
}

func NewDatabase() *Database {
	return &Database{
		events: make(map[StudentId][]IEvent),
	}
}

func (d *Database) Append(event IEvent) {
	events := d.events[event.StreamId()]
	d.events[event.StreamId()] = append(events, event)
}

func (d *Database) GetStudent(studentId StudentId) *Student {
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
