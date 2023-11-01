package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jsonOutIndent   = "    "
	regexTemplate   = `(?m)^( +"%s": )\d+,$`
	replaceTemplate = `${1}"%+v",`
)

var (
	expRe *regexp.Regexp = regexp.MustCompile(fmt.Sprintf(regexTemplate, "exp"))
	iatRe *regexp.Regexp = regexp.MustCompile(fmt.Sprintf(regexTemplate, "iat"))
)

func parseAndFormat(jwtS string) (string, error) {
	p := jwt.NewParser()
	token, _, err := p.ParseUnverified(
		jwtS,
		jwt.MapClaims{},
	)
	if err != nil {
		return EmptyString, err
	}

	headerFormatted, err := json.MarshalIndent(token.Header, "", jsonOutIndent)
	if err != nil {
		return EmptyString, err
	}

	bs, err := json.MarshalIndent(token.Claims, "", jsonOutIndent)
	if err != nil {
		return EmptyString, err
	}
	claimsFormated := string(bs)
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return EmptyString, err
	}
	if exp != nil {
		tmpClaims := expRe.ReplaceAllString(
			claimsFormated,
			fmt.Sprintf(replaceTemplate, exp),
		)
		claimsFormated = tmpClaims
	}
	iat, err := token.Claims.GetIssuedAt()
	if err != nil {
		return EmptyString, err
	}
	if iat != nil {
		tmpClaims := iatRe.ReplaceAllString(
			claimsFormated,
			fmt.Sprintf(replaceTemplate, iat),
		)
		claimsFormated = tmpClaims
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
