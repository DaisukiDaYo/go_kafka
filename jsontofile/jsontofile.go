package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("../kafka-error-logs/jsontofile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1)

	w := bufio.NewWriter(f)
	file, err := w.WriteString(string(res1B))
	if err != nil {
		panic(err)
	}
	fmt.Println("file", file)
	w.Flush()

}
