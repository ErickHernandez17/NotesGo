package note

import "time"

type Note struct {
	Title     string
	Body      string
	CreatedAt string
}

func New(title, body string) *Note {
	return &Note{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

func (n *Note) String() string {
	return n.Title + "\n" + n.Body + "\n" + n.CreatedAt
}

func (n *Note) Json() string {
	return `{"title": "` + n.Title + `", "body": "` + n.Body + `", "created_at": "` + n.CreatedAt + `"}`
}
