package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Structure for goals
type goal struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		return
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Could not read data file")
	}
	var goals []goal
	json.Unmarshal(bytes, &goals)

	new := flag.Bool("n", false, "Create new goal")
	list := flag.Bool("l", false, "List all goals")
	flag.Parse()

	if *new {
		newGoal(goals)
	} else if *list {
		listgoals(goals)
	}
}

// newGoal creates a new goal by prompting the user for a title and description.
//
// It reads the user input from the standard input and creates a new `goal` struct
// with the provided title, description, and the current date and time as the creation date.
// The `goal` struct is then appended to a slice of `goal` structs.
// After that, the slice is marshaled into JSON format and written to a file named "data.json".
//
// No parameters are passed to this function.
// This function does not return any values.
func newGoal(goals []goal) {
	fmt.Print("Goal title: ")
	r := bufio.NewReader(os.Stdin)
	in, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input for Title")
	}
	fmt.Print("Goal Description:\n")
	in2, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input for Title")
	}
	newGoal := goal{
		Title:       in,
		Description: in2,
		DateCreated: time.Now(),
	}

	goals = append(goals, goal{newGoal.Title, newGoal.Description, newGoal.DateCreated})
	g, err := json.MarshalIndent(goals, "", "   ")
	if err != nil {
		log.Fatal("Something man IDK")
	}
	err = os.WriteFile("data.json", g, 0666)
	if err != nil {
		log.Fatal("Could not write goal to data file")
	}
}

func listgoals(goals []goal) {
	for i := range goals {
		fmt.Printf("Goal %d => Title: %s ->Description: %s ->Time: %v\n\n", i, goals[i].Title, goals[i].Description, goals[i].DateCreated)
	}
}
