package student

type Database struct {
	events   map[StudentId][]IEvent
	students map[StudentId]*Student
}

func NewDatabase() *Database {
	return &Database{
		events:   make(map[StudentId][]IEvent),
		students: make(map[StudentId]*Student),
	}
}

func (d *Database) Append(event IEvent) {
	events := d.events[event.StreamId()]
	d.events[event.StreamId()] = append(events, event)

	student := d.GetStudent(event.StreamId())
	if student != nil {
		d.students[event.StreamId()] = student
	}
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

func (d *Database) GetStudentView(studentId StudentId) *Student {
	return d.students[studentId]
}
