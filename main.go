package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "embed"
)

func printUsage() {
	fmt.Println("Usage: asciifier [-f=<fontname> -p=<prefix>] <text>")
	flag.PrintDefaults()
}

//go:embed res/ansi-shadow.txt
var font string

type UnknownFontError struct {
	fontname string
}

func (e UnknownFontError) Error() string {
	return fmt.Sprintf("Unknown font: %s", e.fontname)
}

// In the future we could do a embed.fs with multiple fonts, that could be fun
func readFont(fontname string) (output map[rune][]string, err error) {

	// Currently we only have one, ansi-shadow
	if fontname != "ansi-shadow" {
		return nil, &UnknownFontError{fontname}
	}

	output = map[rune][]string{}

	// Font splits characters by empty lines, with the first line
	// containing the characters to map to this glyph and
	// the rest of the lines containing the glyph itself
	chars := strings.Split(font, "\n\n")

	for _, char := range chars {
		l := strings.Split(char, "\n")
		t := l[0]
		c := l[1:]

		for _, r := range t {
			output[rune(r)] = c
		}
	}

	return output, nil
}

func asciify(font map[rune][]string, prefix string, text string) (output string) {
	output = ""
	height := len(font['A'])

	for i := range height {
		output += prefix
		
		for _, r := range text {
			if glyph, ok := font[r]; ok {
				output += glyph[i]
			}
		}

		output += "\n"
	}

	return output
}

func main() {

	fontname := flag.String("f", "ansi-shadow", "font name")
	prefix := flag.String("p", "", "prefix for each line, for indentation or use as code comment")
	flag.Parse()
	
	text := strings.Join(flag.Args(), " ")

	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		return
	}

	font, err := readFont(*fontname)
	if err != nil {
		fmt.Println(err)
		return
	}

	output := asciify(font, *prefix, text)
	fmt.Print(output)
}

