package main

import (
	"io"
	"flag"
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: asciifier [-font=<foo>] <text>")
	fmt.Println("   - available fonts: ")
	fmt.Println("     - foo")
	fmt.Println("     - bar")
	fmt.Println("     - baz")
}

func readFont(filename string) (output map[rune]string, err error) {
	output = map[rune]string{}

	file, err := os.Open("res/"+filename)
	if err != nil {
		return output, err
	}
	defer file.Close()


	buf, err := io.ReadAll(file)
	if err != nil {
		return output, err
	}

	bufs := string(buf)
	fmt.Println("file content:")
	fmt.Print(bufs)

	return output, nil

}


func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage()
		return
	}

	font := flag.String("font", "foo", "font name")
	flag.Parse()

	text := strings.Join(flag.Args(), " ")

	glyphs, err := readFont("ansi-shadow.txt")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	fmt.Println(glyphs)



	switch *font {
	case "foo":
		fmt.Println("foo font:", text)
	case "bar":
		fmt.Println("bar font:", text)
	case "baz":
		fmt.Println("baz font:", text)
	default:
		fmt.Println("Unknown font:", font)
		printUsage()
			
	}


}


