package main

import (
	"fmt"
	"log"
	"path"

	"github.com/coderconvoy/lazyf"
)

func Interactive() {
	_jsloc := lazyf.FlagString("js", "invoices.json", "js-list", "The location of the json file for the list")

	_rt := lazyf.FlagString("rp", "", "root-path", "The Root folder for relative locations")
	_pre := lazyf.FlagString("pre", "", "id-prefix", "Prefix for invoice ID's")

	_, cpath := lazyf.FlagLoad("cf", "conf.lz", "{HOME}/.config/invoices/conf.lz")

	jsloc := lazyf.PlusPathEnv(path.Dir(cpath), *_rt, *_jsloc)

	invs, err := LoadInvoices(jsloc)

	if err != nil {
		log.Fatal(err)
	}

loop:
	for {
		n := askOptions("What would you like to do?", []string{"List Invoices", "New invoice", "New invoice from Base", "Edit Invoice", "Print Invoice", "quit"})

		switch n {
		case 0:
			f := askString("Filter?", "", nil)
			ivs := FilterInvoices(invs, f)
			for _, v := range ivs {
				fmt.Println(v.OneLine())
			}
		case 1:
			inv, err := BuildInvoice(*_pre, len(invs), Invoice{}, false)
			if err != nil {
				fmt.Println("Error :" + err.Error())
				continue
			}
			invs = append(invs, inv)
		default:
			break loop
		}
	}
	for _, v := range invs {
		fmt.Println(v)
	}
	SaveInvoices(invs, jsloc)
}

func main() {

	Interactive()
	//iv, _ := BuildInvoice("", 0, Invoice{}, false)
	//fmt.Println(iv)

}
