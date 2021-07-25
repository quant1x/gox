package lang

import "testing"

func TestNew(t *testing.T) {
	var e interface{}
	e = New(404, "not found")
	exp, ok := e.(error)
	if !ok {
		t.Fatalf("类型转换失败")
	}
	t.Logf("%+v", exp)
}
