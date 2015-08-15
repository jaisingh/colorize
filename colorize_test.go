package main

import (
	"fmt"
	"testing"

	"github.com/mgutz/ansi"
)

func TestColorize(t *testing.T) {
	for i := 0; i < 256; i++ {
		f := fmt.Sprintf("%d:default", i)
		b := fmt.Sprintf("default:%d", i)
		fmt.Println(ansi.Color(f, f), ansi.Color(b, b))
	}

}
