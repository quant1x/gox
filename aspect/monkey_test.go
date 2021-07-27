package aspect_test

import (
	"fmt"
	"github.com/mymmsc/gox/api"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/mymmsc/gox/aspect"
)

func no() bool  { return false }
func yes() bool { return true }

type Float64 float64

func (this Float64) GoString() string {
	return api.ToString(this)
}

func TestFloat64(t *testing.T) {
	var f Float64
	f = 123
	aspect.PatchInstanceMethod(reflect.TypeOf(f), "GoString", func(v Float64) string {
		return fmt.Sprintf("%.2f", float64(v))
	})
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%#v\n", f)
}

func TestTimePatch(t *testing.T) {
	before := time.Now()
	aspect.Patch(time.Now, func() time.Time {
		return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	})
	during := time.Now()
	assert(t, aspect.Unpatch(time.Now))
	after := time.Now()

	assert(t, time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC) == during)
	assert(t, before != during)
	assert(t, during != after)
}

func TestGC(t *testing.T) {
	value := true
	aspect.Patch(no, func() bool {
		return value
	})
	defer aspect.UnpatchAll()
	runtime.GC()
	assert(t, no())
}

func TestSimple(t *testing.T) {
	assert(t, !no())
	aspect.Patch(no, yes)
	assert(t, no())
	assert(t, aspect.Unpatch(no))
	assert(t, !no())
	assert(t, !aspect.Unpatch(no))
}

func TestGuard(t *testing.T) {
	var guard *aspect.PatchGuard
	guard = aspect.Patch(no, func() bool {
		guard.Unpatch()
		defer guard.Restore()
		return !no()
	})
	for i := 0; i < 100; i++ {
		assert(t, no())
	}
	aspect.Unpatch(no)
}

func TestUnpatchAll(t *testing.T) {
	assert(t, !no())
	aspect.Patch(no, yes)
	assert(t, no())
	aspect.UnpatchAll()
	assert(t, !no())
}

type s struct{}

func (s *s) yes() bool { return true }

func TestWithInstanceMethod(t *testing.T) {
	i := &s{}

	assert(t, !no())
	aspect.Patch(no, i.yes)
	assert(t, no())
	aspect.Unpatch(no)
	assert(t, !no())
}

type f struct{}

func (f *f) No() bool { return false }

func TestOnInstanceMethod(t *testing.T) {
	i := &f{}
	assert(t, !i.No())
	aspect.PatchInstanceMethod(reflect.TypeOf(i), "No", func(_ *f) bool { return true })
	assert(t, i.No())
	assert(t, aspect.UnpatchInstanceMethod(reflect.TypeOf(i), "No"))
	assert(t, !i.No())
}

func TestNotFunction(t *testing.T) {
	panics(t, func() {
		aspect.Patch(no, 1)
	})
	panics(t, func() {
		aspect.Patch(1, yes)
	})
}

func TestNotCompatible(t *testing.T) {
	panics(t, func() {
		aspect.Patch(no, func() {})
	})
}

func assert(t *testing.T, b bool, args ...interface{}) {
	t.Helper()
	if !b {
		t.Fatal(append([]interface{}{"assertion failed"}, args...))
	}
}

func panics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if v := recover(); v == nil {
			t.Fatal("expected panic")
		}
	}()
	f()
}
