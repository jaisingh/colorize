package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Colorize every other line

var colorFlag = flag.String("color", "", "foregroundColor+attributes:backgroundColor+attributes")

var color func(string) string

func init() {
	flag.Parse()
	if *colorFlag != "" {
		color = ansi.ColorFunc(*colorFlag)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count += 1

		line := scanner.Text()
		if count%2 == 0 && *colorFlag != "" {
			fmt.Println(color(line))
			continue
		}

		fmt.Println(line)
	}

}
