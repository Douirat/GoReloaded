package main

import (
	"fmt"
	"os"
	"reload"
)

func main() {
	Arguments := os.Args[1:]
	file_name := "../sample.txt"
	// result_file := "../result.txt"
	if len(Arguments) < 1 || len(Arguments) > 1 {
		fmt.Println("Not enough arguments!!!")
		return
	}
	the_file := reload.NewFile(file_name, Arguments[0])
	var newFile reload.Filer = the_file
	var implementation reload.File_Handler = the_file
	newFile.WriteInto()
	implementation.Determine()
	// data := newFile.CopyData()
	// fmt.Println(data)
}

