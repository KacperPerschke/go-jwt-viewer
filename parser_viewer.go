package main

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func parseAndShow(ts string) error {

	token, err := jwt.Parse(
		ts,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		},
	)

	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return errors.New("to nie wygląda na JWT")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return errors.New("skopany podpis")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return errors.New("przedatowany")
	} else {
		return fmt.Errorf(`Couldn't handle this token: %w`, err)
	}
	return nil
}
