package main

import (
	"encoding/json"
	"fmt"
	"strings"

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

	alg := fmt.Sprintf("%s", token.Header["alg"])
	algFS := strings.Replace(alg, "HS", "HMACSHA", -1)
	signatureFormated := strings.Join(
		[]string{
			fmt.Sprintf(`%s(`, algFS),
			`    base64UrlEncode(header) + "." + base64UrlEncode(payload),`,
			`    your-256-bit-secret`,
			`)`,
		},
		"\n",
	)

	out := fmt.Sprintf(
		"===\nHeader →\n%s\nClaims →\n%s\nSignature →\n%s\n===\n",
		headerFormatted,
		claimsFormated,
		signatureFormated,
	)
	return out, nil
}
