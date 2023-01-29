package main

import (
	"fmt"
	"github.com/mymmsc/gox/labs/aspect"
	"os"
	"strings"
)

func main() {
	aspect.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?")
}
