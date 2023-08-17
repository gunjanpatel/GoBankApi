package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	bank "github.com/gunjanpatel/GoBank"
)

const API_DOMAIN = "localhost:8000"

// Custome struct to extend statement method
type CustomAccount struct {
	*bank.Account
}

var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Gunjan",
				Address: "Test 123, Denmark",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}

	accounts[1002] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Prital",
				Address: "Test 123, Denmark",
				Phone:   "(213) 555 0147",
			},
			Number: 1002,
		},
	}

	fmt.Printf("Starting server at: http://%v", API_DOMAIN)
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe(API_DOMAIN, nil))
}

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

	fmt.Fprintf(w, bank.Statement(account))
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	number, err := strconv.ParseFloat(numberqs, 64)

	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	}

	amount, err := strconv.ParseFloat(amountqs, 64)

	if err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
		return
	}

	account, ok := accounts[number]

	if !ok {
		fmt.Fprintf(w, "Account with number %v is not exist", number)
		return
	}

	err = account.Deposit(amount)

	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	fmt.Fprintf(w, bank.Statement(account))
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Missing account number")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Account Number is invalid!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Amount is invalid!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, bank.Statement(account))
			}
		}
	}
}

// Generate statement
func (a CustomAccount) Statement() string {
	jsonData, err := json.Marshal(a)

	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}
