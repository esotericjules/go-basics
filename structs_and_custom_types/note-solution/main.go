package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interfaces/todo"

	"example.com/note-solution/note"
)

// an interface is a contract that guarantees
// that a certain value (typically a struct in go) has a certain method
// it can be used anywhere a type can be used
// it is not about defininig logic but establishing that a contract
// exists. it also extablishes the types a method parameters accepts
// An interface can define several methods in it.
// But if you have only one method in an interface,
// the convention is that the name of the interface should
// start with the name of the method + er
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething("HELLO")
	title, content := getNoteData()
	todoText := getUserInput("Enter your todo text")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()
	// err = saveData(todo)

	err = outputData(todo)

	if err != nil {
		return
	}

	// userNote.Display()

	// usernote can be used as a value for the saveData fn
	// which expects a parameter that adheres to the saver interface
	// because the note package has a Save method defined in it
	// err = saveData(userNote)

	outputData(userNote)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// using the saver interface here as a type for data
// means that whatever type of value data will be
// it will be of some type that adheres to that saver interface,
// meaning that it has signed the contract the interface provides

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note suceeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

// any type
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stingVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stingVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

// generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
