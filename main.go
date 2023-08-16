package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gunjanpatel/go-bank"
)

var accounts = map[float64]*bank.Account{}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	number, err := strconv.ParseFloat(numberqs, 64)

	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	}

	account, ok := accounts[number]

	if !ok {
		fmt.Fprintf(w, "Account with number %v is not exist", number)
	}

	fmt.Fprintf(w, account.Statement())
}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Gunjan",
			Address: "Test 123, Denmark",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
