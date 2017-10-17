package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("f_input_file")
	check(err)
	defer f.Close()
	for i := 1; i <= 1000; i++ {
		f.WriteString(fmt.Sprintf("page %d hello world\f", i))
	}

	f, err = os.Create("l_input_file")
	check(err)
	defer f.Close()
	for i := 1; i <= 1000; i++ {
		f.WriteString(fmt.Sprintf("line %d hello world\n", i))
	}

}
