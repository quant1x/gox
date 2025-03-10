package mem

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	data := []byte("hello world")
	err := Lock(data)
	fmt.Println(err)
	err = Unlock(data)
	fmt.Println(err)
}

//func Test_goid(t *testing.T) {
//	id := goid()
//	fmt.Println(id)
//}
