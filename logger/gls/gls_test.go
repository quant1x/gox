package gls

import (
	"sync"
	"testing"
)

func TestGls(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go WithGls(func() {
		if nil != Get("hello") {
			t.Errorf("1...\n")
			return
			//t.Fail()
		}
		Set("hello", "world")
		if "world" != Get("hello") {
			t.Errorf("2...\n")
			return
			//t.Fail()
		}
		if !IsGlsEnabled(GoID()) {
			t.Errorf("3...\n")
			return
			//t.Fail()
		}
		wg.Done()
	})()
	wg.Wait()
	if IsGlsEnabled(GoID()) {
		t.Fatalf("4...\n")
		//t.Fail()
	}
	if nil != Get("hello") {
		t.Fatalf("5...\n")
		//t.Fail()
	}
	//SetIndex("hello", "world") // will panic
}

func TestNestedGls(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go WithGls(func() {
		Set("hello", "world")
		go WithGls(func() {
			if "world" != Get("hello") {
				t.Errorf("6...\n")
				return
				//t.Fail()
			}
			wg.Done()
		})()
	})()
	wg.Wait()
}
