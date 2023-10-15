package main

func main() {
	tokenEncrypted, errR := readData()
	if errR != nil {
		panic(errR)
	}
	errP := parseAndShow(tokenEncrypted)
	if errP != nil {
		panic(errP)
	}
}
