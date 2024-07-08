package file

import (
	"notes/dev/note"
	"os"
)

func NewNote(note *note.Note) {
	file, err := os.Create(note.Title + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(note.Json())
}
