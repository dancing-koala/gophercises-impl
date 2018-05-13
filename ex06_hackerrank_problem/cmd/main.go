package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var matcher = regexp.MustCompile("[A-Z]")

func camelcase(s string) int32 {
	matches := matcher.FindAllStringIndex(s, -1)

	if matches == nil {
		return 1
	}

	return int32(len(matches) + 1)
}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func keysToVal(keys []string) map[string]int32 {

	result := make(map[string]int32)

	for i, k := range keys {
		result[k] = int32(i)
	}

	return result
}

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	kvmap := keysToVal(alphabet)

	var buffer bytes.Buffer

	for _, char := range s {
		strChar := string(char)

		index, ok := kvmap[strings.ToLower(strChar)]

		if !ok {
			buffer.WriteString(strChar)
			continue
		}

		index = (index + k) % 26

		if matcher.MatchString(strChar) {
			buffer.WriteString(strings.ToUpper(alphabet[index]))
		} else {
			buffer.WriteString(alphabet[index])
		}
	}

	return buffer.String()
}

func main() {
	fmt.Println(camelcase("iLoveCodingInGolang"))
	fmt.Println(caesarCipher("abcdefgHIJklmnopqrstuvwxyz", 3))
}
