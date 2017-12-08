package collections

import (
	"fmt"
	"testing"
	"time"
)

type model struct {
	f1 string
	f2 int
	f3 int16
	f4 int32
	f5 int64
	f6 float32
	f7 float64
	f8 time.Time

	subs Enumerable
}

type sub struct {
	g1 string
	g2 int
	g3 int16
	g4 int32
	g5 int64
	g6 float32
	g7 float64
	g8 time.Time
}

func BuildCase() []interface{} {
	c := make([]interface{}, 0)
	c = append(c, &model{
		f1: "l1f1",
		f2: 12,
		f3: 13,
		f4: 14,
		f5: 16,
		f7: 17,
		f8: time.Now(),
		subs: NewEnumerableFromSource([]interface{}{
			&sub{
				g1: "l1f1g1",
				g2: 112,
				g3: 113,
				g4: 114,
				g5: 116,
				g7: 117,
				g8: time.Now(),
			},
			&sub{
				g1: "l2f1g1",
				g2: 212,
				g3: 213,
				g4: 214,
				g5: 216,
				g7: 217,
				g8: time.Now(),
			},
		}),
	})
	return c
}

func Test_enumerator(t *testing.T) {
	e := NewEnumeratorFromSource(BuildCase())
	ForEach(e, func(v interface{}) {
		fmt.Println(v.(*model).f1)
	})
}

func Test_enumerable(t *testing.T) {
	e := NewEnumerableFromSource(BuildCase())
	c := e.Count()
	if c != 1 {
		t.Fail()
	}
	if ok, _ := e.Any(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); !ok {
		t.Fail()
	}
	if c, _ = e.CountBy(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); c != 1 {
		t.Fail()
	}
	if _, err := e.First(func(v interface{}) bool {
		return v.(*model).f1 == "shouldbenil"
	}); err == nil {
		t.Fail()
	}
	if m, _ := e.First(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); m == nil {
		t.Fail()
	}
	if _, err := e.FirstOrDefault(func(v interface{}) bool {
		return v.(*model).f1 == "shouldbenil"
	}); err == nil {
		t.Fail()
	}
	if m, _ := e.FirstOrDefault(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); m == nil {
		t.Fail()
	}
	if _, err := e.Single(func(v interface{}) bool {
		return v.(*model).f1 == "shouldbenil"
	}); err == nil {
		t.Fail()
	}
	if m, _ := e.Single(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); m == nil {
		t.Fail()
	}
	if _, err := e.SingleOrDefault(func(v interface{}) bool {
		return v.(*model).f1 == "shouldbenil"
	}); err == nil {
		t.Fail()
	}
	if m, _ := e.SingleOrDefault(func(v interface{}) bool {
		return v.(*model).f1 == "l1f1"
	}); m == nil {
		t.Fail()
	}
}
