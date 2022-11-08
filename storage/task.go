package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"pueblomo/kanbancli/global"
	"pueblomo/kanbancli/model"

	"github.com/charmbracelet/bubbles/list"
)

type store struct {
	Todo       []list.Item
	InProgress []list.Item
	Done       []list.Item
}

type storeOut struct {
	Todo       []model.Task
	InProgress []model.Task
	Done       []model.Task
}

func CheckFileExists(project string) {
	if _, err := os.Stat(dir + "/" + global.StoragePath + "/" + project + "/" + global.StorageName); errors.Is(err, os.ErrNotExist) {
		errCreate := os.MkdirAll(dir+"/"+global.StoragePath+"/"+project, os.ModePerm)
		if errCreate != nil {
			log.Fatalln(errCreate)
		}
		file, err := os.Create(dir + "/" + global.StoragePath + "/" + project + "/" + global.StorageName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func WriteToFile(todo []list.Item, inProgress []list.Item, done []list.Item, project string) {
	store := store{Todo: todo, InProgress: inProgress, Done: done}
	file, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(dir+"/"+global.StoragePath+"/"+project+"/"+global.StorageName, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(project string) []list.Model {
	file, err := os.ReadFile(dir + "/" + global.StoragePath + "/" + project + "/" + global.StorageName)
	if err != nil {
		log.Fatal(err)
	}
	var models *[]list.Model
	if len(file) > 0 {
		var store *storeOut
		err = json.Unmarshal(file, &store)
		if err != nil {
			log.Fatal(err)
		}
		todoItems := []list.Item{}
		inProgressItems := []list.Item{}
		doneItems := []list.Item{}
		for _, v := range store.Todo {
			todoItems = append(todoItems, v)
		}
		for _, v := range store.InProgress {
			inProgressItems = append(inProgressItems, v)
		}
		for _, v := range store.Done {
			doneItems = append(doneItems, v)
		}
		todoList := list.New(todoItems, list.NewDefaultDelegate(), 0, 0)
		inProgressList := list.New(inProgressItems, list.NewDefaultDelegate(), 0, 0)
		doneList := list.New(doneItems, list.NewDefaultDelegate(), 0, 0)
		todoList.SetShowHelp(false)
		inProgressList.SetShowHelp(false)
		doneList.SetShowHelp(false)
		models = &[]list.Model{todoList, inProgressList, doneList}
	} else {
		defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
		defaultList.SetShowHelp(false)
		models = &[]list.Model{defaultList, defaultList, defaultList}
	}
	(*models)[global.Todo].Title = "ToDo"
	(*models)[global.InProgress].Title = "InProgress"
	(*models)[global.Done].Title = "Done"
	return *models
}
