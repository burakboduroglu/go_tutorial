package main

import (
	"os"
)

func main() {
	const filePath = "file.txt"

	writeToFile(filePath, "Hello, World!")
	readFromFile(filePath)
}

// nil means no value
// panic is a built-in function that stops the ordinary flow of control and begins panicking.
// When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.

// Write to a file function
// 0644 is the permission for the file
// []byte(text) converts the text to a byte array
func writeToFile(name string, text string) {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
	println("***Text written successfully.***")
}

// Read from a file function
func readFromFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	text := string(data)
	println("Text read: " + text)
	println("***Text read successfully.***")
}
