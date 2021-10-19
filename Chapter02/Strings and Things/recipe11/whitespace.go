package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {

	stringToTrim := "\t\t\n   Go \tis\t Awesome \t\t"
	trimResult := strings.TrimSpace(stringToTrim)
	fmt.Println(trimResult)

	stringWithSpaces := "\t\t\n   Go \tis\n Awesome \t\t"
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(stringWithSpaces, " ")
	fmt.Println(replace)

	needSpace := "need space"
	fmt.Println(pad(needSpace, 14, "CENTER"))
	fmt.Println(pad(needSpace, 14, "LEFT"))
}

func pad(input string, padLen int, align string) string {
	inputLen := len(input)

	if inputLen >= padLen {
		return input
	}

	repeat := padLen - inputLen
	var output string
	switch align {
	case "RIGHT":
		output = fmt.Sprintf("%[1]*s", -padLen, input)
	case "LEFT":
		output = fmt.Sprintf("%[1]*s", padLen, input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("%[1]*s%[3]*s", left, input, right, "")
	}
	return output
}
