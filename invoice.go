package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/coderconvoy/money"
)

const (
	CASH = 1 << iota
	BACS
)

type Job struct {
	Description string
	Rate        money.M
	NumUnits    money.M
	UnitType    string
}

type Invoice struct {
	Client  string
	Address string
	Methods int
	Date    time.Time
	Jobs    []Job
	PaidIn  int
}

func BuildJob(def Job) Job {
	res := Job{}
	res.Description = askString("Job Description", def.Description, nil)
	res.UnitType = askString("Unit Type", def.UnitType, nil)
	if res.UnitType != "-" {
		res.NumUnits = askMoney("Number of units", def.NumUnits.String())
		res.Rate = askMoney("Rate £/"+res.UnitType, def.Rate.String())
		return res
	}
	res.Rate = askMoney("Agreed Price?", def.Rate.String())
	return res
}

func (j Job) Cost() money.M {
	if j.UnitType == "-" {
		return j.Rate
	}
	return (j.Rate * j.NumUnits) / 100
}

func (j Job) String() string {
	if j.UnitType == "-" {
		return j.Description + ": Agreed Price = £" + j.Rate.String()
	}
	return j.Description + " " + j.NumUnits.String() + " * " + j.UnitType + " at £" + j.Rate.String() + "/" + j.UnitType + " = £" + j.Cost().String()
}

//Ask the user for details of invoice they want.
func BuildInvoice(prefix string, num int, def Invoice, oldDate bool) (Invoice, error) {
	res := Invoice{}
	res.Client = askString("Client Name?", def.Client, nil)
	res.Address = askString("Client Address", def.Address, nil)

	if oldDate {
		res.Date = askDate("Date?", def.Date)
	} else {
		res.Date = askDate("Date?", time.Now())
	}

	for i := 0; i < 10; i++ {
		if len(def.Jobs) > i {
			res.Jobs = append(res.Jobs, BuildJob(def.Jobs[i]))
		}
		res.Jobs = append(res.Jobs, BuildJob(Job{}))

		if !askBool("Would you like to add another job?", false) {
			break
		}
	}

	return res, nil
}

func (iv Invoice) String() string {
	res := fmt.Sprintf("Client:%s\nAddress:%s\nDate:%s\n", iv.Client, iv.Address, iv.Date.Format("02/01/06"))
	var cost money.M
	for _, v := range iv.Jobs {
		cost += v.Cost()
		res += v.String() + "\n"
	}
	res += "Total Cost: £" + cost.String()

	return res
}

func LoadInvoices(fname string) ([]Invoice, error) {
	jdata, err := ioutil.ReadFile(fname)
	var res []Invoice

	if err != nil {
		fmt.Println("Could not load file '" + fname + "'")
		cont := askBool("Do you want to create it?", true)
		if !cont {
			return res, errors.New("Could not read file: Stopping")
		}
		return res, nil
	}

	err = json.Unmarshal(jdata, &res)
	if err != nil {
		fmt.Println("Could not marshal Json Data from '" + fname + "'")
		cont := askBool("Continue as new?", true)
		if !cont {
			return res, errors.New("Could not marshal json from : " + fname)
		}
	}
	return res, nil

}

func SaveInvoices(invoices []Invoice, fname string) error {
	dt, err := json.Marshal(invoices)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fname, dt, 0777)
	return err
}
