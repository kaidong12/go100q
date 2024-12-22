package basic

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	name      string
	lat, long float32
}

func Struct_demo_1() {

	insight := location{"insight", 12.12345, 123.12345}
	fmt.Println(insight)

	opportunity1 := location{name: "opportunity1", lat: 23.12345, long: 123.12345}
	fmt.Println(opportunity1)

	opportunity2 := location{name: "opportunity2", lat: 23.12345, long: 123.12345}
	fmt.Println(opportunity2)

	locations := []location{
		location{"insight1", 12.12345, 123.12345},
		location{"insight2", 12.12345, 123.12345},
		location{name: "opportunity3", lat: 23.12345, long: 123.12345},
		location{lat: 23.12345, long: 123.12345},
	}
	fmt.Println(locations)
}

type location2 struct {
	Name      string
	Lat, Long float32
}

func Struct_demo_2() {

	opportunity1 := location{name: "opportunity1", lat: 23.12345, long: 123.12345}
	fmt.Println(opportunity1)
	bytes1, err := json.Marshal(opportunity1)
	exitOnError(err)
	fmt.Println("bytes1", bytes1)
	fmt.Println(string(bytes1))

	opportunity2 := location2{Name: "opportunity2", Lat: 23.12345, Long: 123.12345}
	fmt.Println(opportunity2)
	bytes2, err := json.Marshal(opportunity2)
	exitOnError(err)
	fmt.Println("bytes2", bytes2)
	fmt.Println(string(bytes2))

	locations := []location2{
		location2{"insight1", 12.12345, 123.12345},
		location2{"insight2", 12.12345, 123.12345},
		location2{Name: "opportunity3", Lat: 23.12345, Long: 123.12345},
		location2{Lat: 23.12345, Long: 123.12345},
	}
	fmt.Println(locations)
	for _, loc := range locations {
		fmt.Println("========In loop========")
		bytesn, err := json.Marshal(loc)
		exitOnError(err)
		fmt.Println(string(bytesn))
	}
}

type location3 struct {
	Name string  `json:"name"`
	Lat  float32 `json:"latitude"`
	Long float32 `json:"longitude"`
}

func Struct_demo_3() {

	opportunity3 := location3{Name: "opportunity2", Lat: 23.12345, Long: 123.12345}
	fmt.Println(opportunity3)
	bytes3, err := json.Marshal(opportunity3)
	exitOnError(err)
	fmt.Println("bytes3", bytes3)
	fmt.Println(string(bytes3))

	locations := []location3{
		location3{"insight1", 12.12345, 123.12345},
		location3{"insight2", 12.12345, 123.12345},
		location3{Name: "opportunity3", Lat: 23.12345, Long: 123.12345},
		location3{Lat: 23.12345, Long: 123.12345},
	}
	fmt.Println(locations)
	for _, loc := range locations {
		fmt.Println("========In loop========")
		bytesN, err := json.Marshal(loc)
		exitOnError(err)
		fmt.Println(string(bytesN))
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
