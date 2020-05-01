package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

func main() {
	fmt.Println("under is sexy writing json...")

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)
	data := `{
		"Page": "2",
		"Fruits": ["apple", "peach", "pear"]
	}`
	var m1 response1
	err := json.Unmarshal([]byte(data), &m1)

	fmt.Println(m1)
	fmt.Println(err)

}
