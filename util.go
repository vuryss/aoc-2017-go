package main

import (
	"io/ioutil"
	"log"
)

func getInput(day string) string {
	byteContent, err := ioutil.ReadFile("inputs/" + day)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(byteContent)
}
