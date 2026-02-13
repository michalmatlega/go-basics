package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

//type displayer interface {
//	Display()
//}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("mehumehu")
	printSomething(true)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todoItem)

	if err != nil {
		return
	}

	fmt.Println("Saving todoItem succeeded!")

	err = outputData(userNote)

	//userNote.Display()
	//err = saveData(userNote)

	if err != nil {
		return
	}
}

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

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}

	//switch value.(type) {
	//case int:
	//	fmt.Println("Integer:", value)
	//case string:
	//	fmt.Println("String:", value)
	//case float64:
	//	fmt.Println("Float:", value)
	//default:
	//	fmt.Println("Unknown type!")
	//}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving todo failed.")
		return err
	}

	fmt.Println("Saving todo succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//var value string
	//fmt.Scanln(&value)	//can not into long input

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
