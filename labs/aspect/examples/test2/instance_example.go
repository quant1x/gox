package main

import (
	"fmt"
	"github.com/mymmsc/gox/labs/aspect"
	"net"
	"net/http"
	"reflect"
)

func main() {
	var d *net.Dialer
	aspect.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}
