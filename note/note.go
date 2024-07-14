package note

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

func (n *Note) Save() error {
	fileName := n.Title + `.json`

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
func (n *Note) Display() {
	fmt.Println(n.Title, n.Body)
}
