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
	Languages    map[string]int
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

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("example.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var cv CV

	// Unmarshal our input YAML file into empty CV (var cv)
	err = yaml.Unmarshal(f, &cv)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Print out the new struct
	fmt.Printf("--- cv:%+v\n", cv)
	fmt.Printf("--- cv dump:\n%s\n\n", string(f))
}
