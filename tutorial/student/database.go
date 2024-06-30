package student

type Database struct {
	events map[StudentID][]IEvent
}

func NewDatabase() *Database {
	return &Database{
		events: make(map[StudentID][]IEvent),
	}
}

func (d *Database) Append(event IEvent) {
	events := d.events[event.StreamID()]
	d.events[event.StreamID()] = append(events, event)
}
