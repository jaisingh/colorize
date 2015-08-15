package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Colorize output

var colorFlag = flag.String("color", "", "foregroundColor+attributes:backgroundColor+attributes")
var altFlag = flag.Bool("alt", false, "Alternate every other line")

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
		line := scanner.Text()

		if *colorFlag != "" {
			if !*altFlag {
				fmt.Println(color(line))
				continue
			}

			count += 1
			if *altFlag && count%2 == 0 {
				fmt.Println(color(line))
				continue
			}
		}
		fmt.Println(line)
	}

}
