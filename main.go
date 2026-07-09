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
	SaveToJson() error
}

type outputable interface {
	saver
	Display()
}

// type displayer interface {
// 	Display()
// }

// type outputable interface {
// 	SaveToJson() error
// 	Display()
// }

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

}

// These two Value1 and Value 2 are basically the same with eachother "interface{}" == any
func printSomething(value interface{}, value2 any) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("Float: ", value)
	default:
		fmt.Println("Any Value: ", value)

	}
	fmt.Println(value)
}

func saveDataToJson(data saver) error {
	err := data.SaveToJson()

	if err != nil {
		fmt.Println("Saving Data the note Failed")
		return err
	}
	fmt.Println("DATA has been save on JSON File")
	return nil
}

func outputData(data outputable) error {
	data.Display()
	return saveDataToJson(data)
}

func getUserInput(promt string) string {
	fmt.Print(promt)
	var value string

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func getTodoData() string {
	return getUserInput("TODO text: ")
}

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")
	return title, content

}
