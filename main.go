package main

import (
	"errors"
	"fmt"
	"notes/dev/file"
	"notes/dev/note"
)

func main() {
	note, err := createNote()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.NewNote(note)

}

func input(label string) (*string, error) {
	fmt.Print(label)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return nil, errors.New("value cannot be empty")
	}
	return &value, nil
}

func createNote() (*note.Note, error) {
	title, err := input("Title: ")
	if err != nil || *title == "" {
		return nil, err
	}
	body, err := input("Body: ")
	if err != nil || *body == "" {
		return nil, err
	}
	return note.New(*title, *body), nil
}
