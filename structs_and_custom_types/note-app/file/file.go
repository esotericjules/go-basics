package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"example.com/note/note"
)

const filename = "my-note.json"
const folder = "note-app"

type NoteFile struct {
	Notes []note.Note `json:"notes"`
}

func CreateNoteFile() {

	filePath := filepath.Join(folder, filename)

	_, err := os.Stat(filename)
	if err == nil {
		// fmt.Println("File already exists")
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Could not create file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Note file created at", filePath)
}

func WriteNoteToFile() {
	var newNote *note.Note

	noteTitle, inputErr := getUserInput("please enter the title of your note ")

	noteDescription, inputErr := getUserInput("please enter the description of your note")
	fmt.Println("noteDescription", noteDescription)

	if inputErr != nil {
		fmt.Println("problem reading input", inputErr)
	}

	newNote, err := note.New(strings.TrimSpace(noteTitle), strings.TrimSpace(noteDescription))
	fmt.Println("newNote", newNote)
	if err != nil {
		fmt.Println("problem creating a new note", err)
	}

	err = saveNote(newNote)
	if err != nil {
		fmt.Println("Could not save note:", err)
		return
	}
	fmt.Println("new note is ", newNote.Name, newNote.Text, newNote.CreatedAt)

}

func saveNote(newNote *note.Note) error {
	var noteFile NoteFile

	existingData, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if len(existingData) > 0 {
		err = json.Unmarshal(existingData, &noteFile)
		if err != nil {
			return err
		}
	}

	newNote.Id = len(noteFile.Notes) + 1
	noteFile.Notes = append(noteFile.Notes, *newNote)

	data, err := json.MarshalIndent(noteFile, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func getUserInput(msg string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var text strings.Builder

	fmt.Println(msg, "Press Enter on an empty line when finished")

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		if text.Len() > 0 {
			text.WriteString("\n")
		}

		text.WriteString(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Could not read input:", err)
		return "", err
	}

	result := text.String()
	return result, nil
}

func GetAllNotes() {
	// fetch notes from file
	var noteFile NoteFile
	existingData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error getting all notes", err)
	}
	if len(existingData) == 0 {
		fmt.Println("There are no notes")
	}

	// unmarshall the notes
	err = json.Unmarshal(existingData, &noteFile)

	for _, currentNote := range noteFile.Notes {
		fmt.Println("Note:", currentNote.Id)
		fmt.Println("Name:", currentNote.Name)
		fmt.Println("Text:", currentNote.Text)
		fmt.Println("==========")
	}

}
