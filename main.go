package main

import (
	"fmt"

	"github.com/coderconvoy/lazyf"
	"github.com/pkg/errors"
)

func main() {

	_ = lazyf.FlagString("js", "invoices.json", "js-list", "The location of the json file for the list")

	_, _ = lazyf.FlagLoad("cf", "conf.lz", "{HOME}/.config/invoices/conf.lz")

	s := askString("Hello, What is your name", "Pete", func(s string) error {
		if s[0] == 'D' {
			return errors.Errorf("No D: %s", s)
		}
		return nil
	})

	fmt.Println("You said:" + s)
}
