package threadlocal

import (
	"fmt"
	"sync"
)

func ExampleGo() {
	Init()
	var wg sync.WaitGroup
	wg.Add(2)

	Set(`key`, `main value`)
	Go(func() {
		Set(`key`, `value 1`)
		Go(func() {
			Set(`key`, `value 2`)
			v, _ := Get(`key`)
			fmt.Printf("g2 = %s\n", v)
			Delete(`key`)
			wg.Done()
		})
		v, _ := Get(`key`)
		fmt.Printf("g1 = %s\n", v)
		Delete(`key`)
		wg.Done()
	})
	wg.Wait()
	v, _ := Get(`key`)
	fmt.Printf("main = %v\n", v)
	// Unordered output:
	// g1 = value 1
	// g2 = value 2
	// main = main value
}
