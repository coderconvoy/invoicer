package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

const (
	CASH = 1 << iota
	BACS
)

type Job struct {
	Description string
	Rate        int //Cost per unit
	Units       string
}

type Invoice struct {
	Client  string
	Address string
	Method  int
	Date    time.Time
	Jobs    []Job
}

//Ask the user for details of invoice they want.
func BuildInvoice(prefix string, num int) (Invoice, error) {
	return Invoice{}, nil
}

func LoadInvoices(fname string) ([]Invoice, error) {
	res := []Invoice{}
	dt, err := ioutil.ReadFile(fname)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(dt, &res)
	return res, err
}

func SaveInvoices(invoices []Invoice, fname string) error {
	dt, err := json.Marshal(invoices)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fname, dt, 0777)
	return err
}
