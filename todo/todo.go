package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

func New(text string) *Todo {

	return &Todo{
		Text: text,
	}
}

func (t *Todo) String() string {
	return `{"text": "` + t.Text + `"}`
}

func (t *Todo) Json() string {
	return `{"text": "` + t.Text + `"}`
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func (t *Todo) Display() {
	fmt.Println("Todo: ", t.Text)
}
