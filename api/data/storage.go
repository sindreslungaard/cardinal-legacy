package data

import (
	"cardinal/logger"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

const (
	dataFilename = "data.json"
	storagePath  = "./storage"
)

var mu sync.Mutex

func FilePath() string {
	return storagePath + "/" + dataFilename
}

func getOrCreate() Data {

	data, err := os.ReadFile(FilePath())

	if err != nil {
		logger.Info("No configurations found, setting up for first time use")

		config := Data{
			Servers:    make(map[string]Server),
			Containers: make(map[string]Container),
			Users:      make(map[string]User),
		}

		data, err = json.MarshalIndent(config, "", "")

		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(storagePath, os.ModeDir)

		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(FilePath(), data, os.ModePerm)

		if err != nil {
			panic(err)
		}

	}

	var d Data

	err = json.Unmarshal(data, &d)

	if err != nil {
		panic(err)
	}

	return d

}

func save(filePath string, d Data) {
	bytes, err := json.Marshal(d)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filePath, bytes, os.ModePerm)

	if err != nil {
		panic(err)
	}
}

func Copy() Data {
	mu.Lock()
	defer mu.Unlock()

	d := getOrCreate()
	return d
}

func Update(f func(Data) Data) {
	mu.Lock()
	defer mu.Unlock()

	d := getOrCreate()

	data := f(d)

	save(FilePath(), data)
}
