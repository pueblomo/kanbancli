package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"pueblomo/kanbancli/global"
)

func CheckProjectsExist() {
	if _, err := os.Stat(dir + "/" + global.StoragePath + "/projects.json"); errors.Is(err, os.ErrNotExist) {
		errCreate := os.MkdirAll(dir+"/"+global.StoragePath, os.ModePerm)
		if errCreate != nil {
			log.Fatalln(errCreate)
		}
		file, err := os.Create(dir + "/" + global.StoragePath + "/projects.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func WriteProjectsToFile(projects []string) {
	file, err := json.Marshal(projects)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(dir+"/"+global.StoragePath+"/projects.json", file, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadProjects() []string {
	file, err := os.ReadFile(dir + "/" + global.StoragePath + "/projects.json")
	if err != nil {
		log.Fatal(err)
	}

	var projects *[]string
	if len(file) > 0 {
		err = json.Unmarshal(file, &projects)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		projects = &[]string{"Main"}
	}

	return *projects
}
