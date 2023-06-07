package main

import (
	"fmt"
	"log"
)

func hellow_world() {
	fmt.Println("Hello world")
}

func hello_name() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Richarlison", "Edson", "Udson"}
	messages, err := Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
