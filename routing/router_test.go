package routing

import "testing"

func TestRouter_Add(t *testing.T) {
	r := NewRouter()

	r.Add("GET", "/hello")
	r.Add("GET", "/hello/world")
	r.Add("GET", "/foo/bar")

	if result := r.Dispatch("GET", "/hello/world"); result == false {
		t.Fatalf("expected to retrieve to be true, actual %+v", result)
	}
}

func TestRouter_PathNotShadowing(t *testing.T) {
	r := NewRouter()

	r.Add("GET", "/hello")
	r.Add("GET", "/hello/world")
	r.Add("GET", "/foo/bar")

	if result := r.Dispatch("GET", "/hello"); result == false {
		t.Fatalf("expected to retrieve to be true, actual %+v", result)
	}
}
func TestRouter_PerMethodTree(t *testing.T) {
	r := NewRouter()

	r.Add("GET", "/hello")
	r.Add("GET", "/hello/world")

	if result := r.Dispatch("POST", "/hello"); result == true {
		t.Fatalf("expected to retrieve to be true, actual %+v", result)
	}
}
