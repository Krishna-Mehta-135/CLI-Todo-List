package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

//We create a new type called Todos — which is just a slice of Todo structs. Now methods can be attached to Todos
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	//append todo in the todos slice
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error{
	//index is invalid so give error
	if index < 0 || index >= len(*todos){
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err 
	}

	return nil   //index is valid
}

func (todos *Todos) delete(index int) error{
	t := *todos

	if err := t.validateIndex(index) ; err != nil{   //validate index and give error if error isnt empty
		return err
	}

	*todos = append(t[:index], t[index+1:]... )  //delete todo

	return nil
}