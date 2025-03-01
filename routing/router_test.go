package routing

import "testing"

func TestRouter_Add(t *testing.T) {
	r := NewRouter()

	r.Add("/hello")
	r.Add("/hello/world")
	r.Add("/foo/bar")

	if result := r.Dispatch("/hello/world"); result == false {
		t.Fatalf("expected to retrieve to be true, actual %+v", result)
	}
}

func TestRouter_PathNotShadowing(t *testing.T) {
	r := NewRouter()

	r.Add("/hello")
	r.Add("/hello/world")
	r.Add("/foo/bar")

	if result := r.Dispatch("/hello"); result == false {
		t.Fatalf("expected to retrieve to be true, actual %+v", result)
	}
}
