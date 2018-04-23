package main

import (
	"fmt"
	"log"
	"path"

	"github.com/coderconvoy/lazyf"
)

func Interactive() {
	_jsloc := lazyf.FlagString("js", "invoices.json", "js-list", "The location of the json file for the list")

	_, cpath := lazyf.FlagLoad("cf", "conf.lz", "{HOME}/.config/invoices/conf.lz")

	jsloc := plusPath(path.Dir(cpath), *_jsloc)

	invs, err := LoadInvoices(jsloc)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(invs)
}

func main() {

	iv, _ := BuildInvoice("", 0, Invoice{})
	fmt.Println(iv)

}

func plusPath(parent, child string) string {
	if len(child) == 0 {
		return parent
	}
	if child[0] == '/' {
		return child
	}
	return path.Join(parent, child)
}
