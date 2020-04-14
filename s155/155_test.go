package s155

import "testing"

func TestConstructor(t *testing.T) {
	a:=Constructor()
	a.Push(2)
	a.Push(0)

	a.Push(3)
	a.Push(0)

	a.Pop()
	a.Pop()
	a.Pop()
	t.Log(a.GetMin())


}
