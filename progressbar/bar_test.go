package progressbar

import (
	"sync"
	"testing"
	"time"
)

func TestBar(t *testing.T) {

	b := NewBar(1, "1st", 20000)
	b2 := NewBar(2, "2st", 10000)
	b3 := NewBar(3, "3st", 30000)

	b.SetSpeedSection(900, 100)
	b2.SetSpeedSection(900, 100)
	b3.SetSpeedSection(900, 100)
	b4 := NewBar(0, "4st", 4000*3)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 2000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			b2.Add()
			time.Sleep(time.Second / 1000)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 30000; i++ {
			b3.Add()
			time.Sleep(time.Second / 3000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4000; i++ {
			b4.Add(3)
			time.Sleep(time.Second / 300)
		}
		b4.Wait()
	}()
	wg.Wait()
}

func TestBar1(t *testing.T) {

	b := NewBar(1, "1st", 20000)

	//b.SetSpeedSection(900, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 2000)
		}
	}()

	wg.Wait()
}
