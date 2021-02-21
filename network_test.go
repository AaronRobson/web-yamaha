package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//https://blog.questionable.services/article/testing-http-handlers-go/

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if ctype := rr.Header().Get("Content-Type"); ctype != expectedContentType {
		t.Errorf("content type header does not match: got %v want %v", ctype, expectedContentType)
	}

	if err := ValidateHTML(rr.Body.String()); err != nil {
		t.Errorf("HTML is invalid %v", err)
	}
}

func TestPing(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}

	// Check the response body is what we expect.
	expected := `true`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func ValidateHTML(value string) (err error) {
	//https://stackoverflow.com/a/52410528/1011785
	r := strings.NewReader(value)
	d := xml.NewDecoder(r)
	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	d.Strict = true
	d.AutoClose = xml.HTMLAutoClose
	d.Entity = xml.HTMLEntity
	for {
		_, err := d.Token()
		switch err {
		case io.EOF:
			return nil
		case nil:
		default:
			return err
		}
	}
}
