package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const JSONOutIndent = "    "

func parseAndFormat(jwtS string) (string, error) {
	p := jwt.NewParser()
	token, _, err := p.ParseUnverified(
		jwtS,
		jwt.MapClaims{},
	)
	if err != nil {
		return EmptyString, err
	}

	headerFormatted, err := json.MarshalIndent(token.Header, "", JSONOutIndent)
	if err != nil {
		return EmptyString, err
	}
	claimsFormated, err := json.MarshalIndent(token.Claims, "", JSONOutIndent)
	if err != nil {
		return EmptyString, err
	}

	/* What does it tell to me?
	 * alg := fmt.Sprintf("%v", token.Header["alg"])
	 * m := jwt.GetSigningMethod(alg)
	 * fmt.Printf("\n\n→ %s ←\n\n", m)
	 */

	signatureFormated := "{\n    YYY… we don't parse for now.\n}"
	out := fmt.Sprintf(
		"===\nHeader →\n%s\nClaims →\n%s\nSignature →\n%s\n===\n",
		headerFormatted,
		claimsFormated,
		signatureFormated,
	)
	return out, nil
}
