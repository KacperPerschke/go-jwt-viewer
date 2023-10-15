package main

import "fmt"

func main() {
	tokenEncrypted, err := readData()
	if err != nil {
		panic(err)
	}
	toShow, err := parseAndFormat(tokenEncrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println(toShow)
}
