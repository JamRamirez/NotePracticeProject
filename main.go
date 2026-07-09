package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.SaveToJson()

	if err != nil {
		fmt.Println("Saving the note Failed")
		return
	}

	fmt.Println("Note has been save on JSON File")

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

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")
	return title, content

}
