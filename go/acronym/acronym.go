// Package acronym convert a phrase to its acronym
package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate return acronym of string s
func Abbreviate(s string) string {
	var buffer bytes.Buffer
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})
	for _, word := range words {
		buffer.WriteByte(word[0])
	}
	return string(bytes.ToUpper(buffer.Bytes()))
}
