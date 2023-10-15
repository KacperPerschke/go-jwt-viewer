package main

import "fmt"

func main() {
	tokenEncrypted, err := readData()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`%#v`, tokenEncrypted)
}
