package main

import (
	"bufio"
	"fmt"
	"notes/dev/note"
	"notes/dev/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type output interface {
	saver
	displayer
}

func main() {
	note, err := createTodo()
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func input(label string) (string, error) {
	fmt.Print(label)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}

	fmt.Println("Data saved successfully")
	return nil
}

func outputData(data output) {
	data.Display()
	saveData(data)
}

func createNote() (*note.Note, error) {
	title, err := input("Title: ")
	if err != nil || title == "" {
		return nil, err
	}
	body, err := input("Body: ")
	if err != nil || body == "" {
		return nil, err
	}
	return note.New(title, body), nil
}

func createTodo() (*todo.Todo, error) {
	text, err := input("Text: ")
	if err != nil || text == "" {
		return nil, err
	}
	return todo.New(text), nil
}

func printSomething(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println(value.(string))
	case int:
		fmt.Println(value.(int))
	default:
		fmt.Println("Unknown type")
	}
}
