package main

import (
	"fmt"
	"sync"
	"time"

	pgbar "gitee.com/quant1x/gox/progressbar"
)

func progressBarTest() {
	fmt.Printf("\n\n\n")
	b := pgbar.NewBar(1, "1st", 20000)
	b2 := pgbar.NewBar(2, "2st", 10000)
	b3 := pgbar.NewBar(3, "3st", 30000)
	b4 := pgbar.NewBar(4, "4st", 4000*3)
	b.SetSpeedSection(900, 100)
	b2.SetSpeedSection(900, 100)
	b3.SetSpeedSection(900, 100)

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add(1)
			time.Sleep(time.Second / 2000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			b2.Add(1)
			time.Sleep(time.Second / 1000)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 30000; i++ {
			b3.Add(1)
			time.Sleep(time.Second / 3000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4000; i++ {
			b4.Add(3)
			time.Sleep(time.Second / 300)
		}
	}()
	wg.Wait()
}

func main() {
	for i := 0; i < 100; i++ {
		progressBarTest()
		fmt.Printf("\n\n\n\n\n====================================================================================================")
	}
}
