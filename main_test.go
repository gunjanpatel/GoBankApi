package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatementHandler(t *testing.T) {
	// Create a new request for the /statement endpoint with a query parameter
	req, err := http.NewRequest("GET", "/statement?number=1001", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create an http.Handler by calling the statement function directly
	handler := http.HandlerFunc(statement)

	// Serve the request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Account{Number: 1001, Name: Gunjan, Address: Test 123, Denmark, Phone: (213) 555 0147}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestStatementHandlerMissingAccountNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/statement", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(statement)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Account number is missing!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestStatementHandlerInvalidAccountNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/statement?number=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(statement)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Invalid account number!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
