package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func OpenFile(path string, target interface{}) interface{} {
	file, err := os.Open(path)

	if err != nil {
		log.Fatalln("error opening ", path, err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalln("error read file ", path, err)
	}

	err = json.Unmarshal(data, &target)

	if err != nil {
		log.Fatalln("error unmarshal data", path, err)
	}

	return target

}

func SaveFile(path string, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("failed to save to ", path, err)
	}

	ioutil.WriteFile(path, jsonData, 0644)
}
