package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type CV struct {
	Name         string
	Address      string
	Phone        string
	WorkFromHome bool
	Mail         string
	Internships  []string
	Education    Education
	Languages    map[string]string
}

type Education struct {
	Msc        float32
	Bsc        float32
	Xii        float32
	Courses    []Course
	References string
}

type Course struct {
	Name  string
	Score int
}

func (file *CV) Parse(data []byte) error {
	return yaml.Unmarshal(data, file)
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cv CV

	// Unmarshal our input YAML file into empty CV (var cv)
	if err := cv.Parse(f); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	fmt.Printf("%+v\n", cv)
}
