package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const num = 100000.5678

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("en %.2f \n", num)

	p = message.NewPrinter(language.German)
	p.Printf("de %.2f \n", num)

	fmt.Println()

	p = message.NewPrinter(language.Swedish)
	p.Printf("sv %.2f \n", num)

	p = message.NewPrinter(language.Icelandic)
	p.Printf("is %.2f \n", num)

	p = message.NewPrinter(language.Norwegian)
	p.Printf("no %.2f \n", num)

	p = message.NewPrinter(language.Finnish)
	p.Printf("fi %.2f \n", num)

	p = message.NewPrinter(language.Danish)
	p.Printf("da %.2f \n", num)
}
