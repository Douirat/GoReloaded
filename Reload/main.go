package main

import (
	"Reload"
	"fmt"
	"os"
)

// Function to meke sure the files input recieved from the os are valid:
func validArgs(args []string) bool {
	if args[0] != "sample.txt" || args[1] != "result.txt" {
		return false
	}
	return true
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 2 || len(arguments) > 2 {
		fmt.Println("Not enough arguments!")
		return
	} else {
		if validArgs(arguments) {
			file_sample := Reload.NewFile("../"+arguments[0], "")
			file_sample.Read()
			if len(file_sample.Content) < 1 {
				fmt.Println("Your file is empty!")
				return
			} else {
				Data := Reload.NewInput(file_sample.Content)
				new_text := Reload.NewText()
				new_text.DivideInput(Data)
				if new_text.IsEmpty() {
					fmt.Println("Not enough arguments!!!")
				} else {
					new_text.RockAndRoll()
					text := new_text.First.Dequeue()
					result_file := Reload.NewFile("../result.txt", text)
					fmt.Println(result_file.Content)
					result_file.CopyData()
					fmt.Println(new_text.Size)
				}
			}
		}
	}
}
