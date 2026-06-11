package main

import (
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreateAt    time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreateAt:    time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}
func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos
	todo := &t[index]

	todo.Completed = !todo.Completed

	if todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	return nil
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return fmt.Errorf("index out of range")
	}
	return nil
}
