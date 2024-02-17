package basic

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Duration int      `json:"duration"`
	Actors   []string `json:"actors"`
}

func Struct_to_json() {
	movie := Movie{
		Title:    "The Godfather",
		Year:     1972,
		Duration: 175,
		Actors:   []string{"aa", "bb", "cc"},
	}

	// 结构体 -> jsonstr
	jsonData, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshalling movie to JSON:", err)
		return
	}

	fmt.Println("JSON String:", string(jsonData))

	// jsonstr -> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonData, &myMovie)
	if err != nil {
		fmt.Println("Error marshalling movie to JSON:", err)
		return
	}

	fmt.Printf("Movie data: %v\n", myMovie)

}
