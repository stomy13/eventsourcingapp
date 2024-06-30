package student

type Student struct {
	StudentID   StudentID
	FullName    string
	Email       string
	DateOfBirth string
}

type StudentID string
