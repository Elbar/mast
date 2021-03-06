package si

import "testing"
import "fmt"

func TestEq01(t *testing.T) {
	type pair struct {
		x *state
		y *state
	}

	s := &state{}

	crs := []struct {
		call pair
		resp bool
	}{
		{pair{x: s, y: s}, true},
		{pair{x: nil, y: nil}, false},
		{pair{x: nil, y: &state{}}, false},
		{pair{x: &state{}, y: nil}, false},
		{pair{&state{ID: 1}, &state{ID: 2}}, true},
		{pair{&state{IsFinal: true}, &state{IsFinal: false}}, false},
		{pair{&state{Tail: map[int]bool{1: true}}, &state{Tail: map[int]bool{1: true}}}, true},
	}
	for _, cr := range crs {
		if rst := cr.call.x.eq(cr.call.y); rst != cr.resp {
			t.Errorf("got %v, expected %v, %v\n", rst, cr.resp, cr)
		}
		if rst := cr.call.y.eq(cr.call.x); rst != cr.resp {
			t.Errorf("got %v, expected %v, %v\n", rst, cr.resp, cr)
		}
	}
}

func TestEq02(t *testing.T) {
	x := &state{ID: 1}
	y := &state{ID: 2}
	a := &state{
		Trans: map[byte]*state{1: x, 2: y},
	}
	b := &state{
		Trans: map[byte]*state{1: x, 2: y},
	}
	c := &state{
		Trans: map[byte]*state{1: y, 2: y},
	}
	d := &state{
		Trans: map[byte]*state{1: x, 2: y, 3: x},
	}
	if rst, exp := a.eq(b), true; rst != exp {
		t.Errorf("got %v, expected %v\n", rst, exp)
	}
	if rst, exp := a.eq(c), false; rst != exp {
		t.Errorf("got %v, expected %v\n", rst, exp)
	}
	if rst, exp := a.eq(c), false; rst != exp {
		t.Errorf("got %v, expected %v\n", rst, exp)
	}
	if rst, exp := a.eq(d), false; rst != exp {
		t.Errorf("got %v, expected %v\n", rst, exp)
	}

}

func TestString01(t *testing.T) {
	crs := []struct {
		call *state
		resp string
	}{
		{nil, "<nil>"},
	}
	for _, cr := range crs {
		if rst := cr.call.String(); rst != cr.resp {
			t.Errorf("got %v, expected %v, %v\n", rst, cr.resp, cr)
		}
	}
	r := &state{}
	s := state{
		ID:    1,
		Trans: map[byte]*state{1: nil, 2: r},
		//Output:  map[byte]int{3: 1, 4: 2},
		Tail:    intSet{33: true},
		IsFinal: true,
		Prev:    []*state{nil, r},
	}
	fmt.Println(s.String())
}
