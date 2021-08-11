package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// Do nothing
	default:
		t.Errorf("The return value type is not http.Handler, but is %T", v)
	}

	// var myH1 http.Handler

	// h = NoSurf(myH1)

	// switch v := h.(type) {
	// case http.Handler:
	// 	// Do nothing
	// default:
	// 	t.Errorf("The return value type is not http.Handler, but is %T", v)
	// }

}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// Do nothing
	default:
		t.Errorf("The return value type is not http.Handler, but is %T", v)
	}
}
