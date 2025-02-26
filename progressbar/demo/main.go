package main

import (
	"gitee.com/quant1x/gox/logger"
	"sync"
	"time"

	pgbar "gitee.com/quant1x/gox/progressbar"
)

func v1ProgressBarTest() {
	b := pgbar.NewBar(1, "1st", 20000)
	//b2 := pgbar.NewBar(2, "2st", 10000)
	//b3 := pgbar.NewBar(3, "3st", 30000)
	//b4 := pgbar.NewBar(4, "4st", 4000)
	b.SetSpeedSection(900, 100)
	//b2.SetSpeedSection(900, 100)
	//b3.SetSpeedSection(900, 100)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add(1)
			time.Sleep(time.Second / 2000)
		}
	}()

	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 10000; i++ {
	//		b2.Add(1)
	//		time.Sleep(time.Second / 1000)
	//	}
	//}()
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 30000; i++ {
	//		b3.Add(1)
	//		time.Sleep(time.Second / 3000)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 4000; i++ {
	//		b4.Add(1)
	//		time.Sleep(time.Second / 300)
	//	}
	//}()
	wg.Wait()
}

func v2ProgressBarTest() {
	//fmt.Printf("\n\n\n")
	b := pgbar.NewBar(1, "1st", 20000)
	b2 := pgbar.NewBar(2, "2st", 10000)
	b3 := pgbar.NewBar(3, "3st", 30000)
	b4 := pgbar.NewBar(4, "4st", 4000)
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
			b4.Add(1)
			time.Sleep(time.Second / 300)
		}
	}()
	wg.Wait()
}

func main() {
	logger.InitLogger("/opt/logs/test", logger.INFO)
	for i := 0; i < 100; i++ {
		v2ProgressBarTest()
		//fmt.Printf("\n\n\n\n\n====================================================================================================")
	}
}
