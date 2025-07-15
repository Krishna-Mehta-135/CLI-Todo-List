package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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
		err := errors.New("invalid index")
		fmt.Println(err)
		return err 
	}

	return nil   //index is valid
}

func (todos *Todos) delete(index int) error{

	if err := todos.validateIndex(index) ; err != nil{   //validate index and give error if error isnt empty
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)  //delete todo

	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index] // pointer to the original item

	todo.Completed = !todo.Completed

	if todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	return nil
}


func (todos *Todos) edit(index int, newTitle string) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index]
todo.Title = newTitle

if todo.Completed {
	todo.Completed = false
	todo.CompletedAt = nil
}

	return nil
}

func (todos *Todos) print(){
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos{
		completed:= "❌"
		completedAt:= ""

		if t.Completed{
			completed = "✅"
			if t.CompletedAt != nil{
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}