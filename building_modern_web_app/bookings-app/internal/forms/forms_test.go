package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	// r := httptest.NewRequest("POST", "/some-url", nil)
	postedData := url.Values{}
	form := New(postedData)

	if !form.Valid() {
		t.Error("form is invalid but it should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	// r := httptest.NewRequest("POST", "/some-url", nil)
	postedData := url.Values{}
	form := New(postedData)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form is valid but it should be invalid")
	}
	if form.Errors.Get("a") == "" {
		t.Error("form does not have error on a field but it should have")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	// r = httptest.NewRequest("POST", "/some-url", nil)
	// r.PostForm = postedData
	form = New(postedData)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form is invalid but it should be valid")
	}
	if form.Errors.Get("a") != "" {
		t.Error("form has an error on a field but it should not have")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	if form.Has("a") {
		t.Error("form has a field but it should not have")
	}

	postedData.Add("a", "value")
	form = New(postedData)
	if !form.Has("a") {
		t.Error("form does not have a field but it should have")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "invalid_21_MinLength")
	form := New(postedData)
	if form.MinLength("a", 21) {
		t.Error("MinLength returned true but should be false")
	}

	postedData.Add("b", "valid_18_MinLength")
	form = New(postedData)
	if !form.MinLength("b", 18) {
		t.Error("MinLength returned false but should be true")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("invalid_email", "a@bcd")
	form := New(postedData)
	form.IsEmail("invalid_email")
	if form.Valid() {
		t.Error("Valid email but should be invalid")
	}
	postedData.Add("valid_email", "a@bcd.com")
	form = New(postedData)
	form.IsEmail("valid_email")
	if !form.Valid() {
		t.Error("Invalid email but should be valid")
	}
}
