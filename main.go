package main

import (
	"fmt"
)

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todos)
	if err != nil {
		fmt.Println("Warning: Could not load todos from storage. Starting fresh todos.")
	}

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	err = storage.Save(todos)
	if err != nil {
		fmt.Printf("Error saving todos in storage: %v\n", err)
	}
}
