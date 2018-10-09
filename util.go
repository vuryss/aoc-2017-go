package main

import (
	"io/ioutil"
	"log"
)

type Point2 struct {
	x, y int
}

type Point3 struct {
	x, y, z int
}

func getInput(day string) string {
	byteContent, err := ioutil.ReadFile("inputs/" + day)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(byteContent)
}
