package entry

import "time"

type Entry struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
