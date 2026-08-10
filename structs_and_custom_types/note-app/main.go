package main

import (
	"fmt"

	"example.com/note/file"
)

func main() {
	fmt.Println("hello, what do you want to do?")
	file.CreateNoteFile()

	fmt.Println("1. Read my notes")
	fmt.Println("2. Write a note")

	var option int32
	fmt.Scan(&option)

	if option == 1 {
		file.GetAllNotes()
	}
	if option == 2 {
		file.WriteNoteToFile()
	}
}
