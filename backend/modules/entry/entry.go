package entry

import "time"

type Entry struct {
	Id        string
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
