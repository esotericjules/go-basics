package note

import (
	"errors"
	"time"
)

type Note struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func New(name, text string) (*Note, error) {
	if name == "" || text == "" {
		return nil, errors.New("All inputs are required")
	}

	return &Note{
		Name:      name,
		Text:      text,
		CreatedAt: time.Now(),
	}, nil
}
