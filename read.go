package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const EmptyString = ``

var errMultiLineString = errors.New(`string has more than one line`)

func byteToStrAndValidate(b []byte) (string, error) {
	s := strings.TrimSuffix(string(b), "\n") // Many editors puts an empty line at the end of file
	if strings.Contains(s, "\n") {
		return EmptyString, errMultiLineString
	}
	return s, nil
}

func readData() (string, error) {
	fCont, errRF := readFile()
	if errRF != nil {
		return EmptyString, errRF
	}

	inCont, errIN := readSTDIN()
	if errIN != nil {
		return EmptyString, errIN
	}

	if fCont == EmptyString && inCont == EmptyString {
		return EmptyString, errors.New(`I did not receive the data in the form of a file name nor as a stream on STDIN`)

	}

	if fCont != EmptyString && inCont != EmptyString {
		return EmptyString, errors.New(`I received the data both as a file and as a stream on STDIN`)
	}

	if fCont != EmptyString {
		return fCont, nil
	}

	return inCont, nil
}

func readFile() (string, error) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		return EmptyString, nil
	}
	if len(argsWithoutProg) > 1 {
		return EmptyString, errors.New(`I got more call arguments than one`)
	}

	fName := argsWithoutProg[0]
	fInfo, err := os.Stat(fName)
	if errors.Is(err, os.ErrNotExist) {
		return EmptyString, err
	}
	if fInfo.IsDir() {
		return EmptyString, fmt.Errorf(`%s is a directrory, but not a file`, fName)
	}
	if fInfo.Size() == 0 {
		return EmptyString, fmt.Errorf(`%s is empty`, fName)
	}

	buf, err := ioutil.ReadFile(fName)
	if err != nil {
		return EmptyString, err
	}

	s, err := byteToStrAndValidate(buf)
	if errors.Is(err, errMultiLineString) {
		return EmptyString, fmt.Errorf(`%s contains more than one line`, fName)
	}

	return s, nil
}

func readSTDIN() (string, error) {
	fh := os.Stdin
	fi, err := fh.Stat()
	itIsAPipe := func() bool {
		return (fi.Mode() & os.ModeNamedPipe) != 0
	}
	if !itIsAPipe() {
		return EmptyString, nil
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return EmptyString, err
	}
	s, err := byteToStrAndValidate(buf)
	if errors.Is(err, errMultiLineString) {
		return EmptyString, errors.New(`on STDIN I got more than one line`)
	}
	return s, nil
}
