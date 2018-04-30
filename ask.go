package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/coderconvoy/money"
)

func inputLine() string {
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	return text

}

func parseWideBool(s string) (bool, error) {
	yesses := []string{"y", "yes", "true", "t"}
	nos := []string{"n", "no", "false", "f"}

	s = strings.ToLower(s)
	for _, v := range yesses {
		if s == v {
			return true, nil
		}
	}
	for _, v := range nos {
		if s == v {
			return false, nil
		}
	}
	return false, fmt.Errorf("Could not convert %s to boolean", s)
}

func truetoyes(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
func askBool(question string, def bool) bool {

	rs := askString(question, truetoyes(def), func(s string) error {
		_, err := parseWideBool(s)
		return err
	})
	if rs == "" {
		return def
	}

	res, _ := parseWideBool(rs)
	return res

}

func askInt(question string, def int) int {
	rs := askString(question, strconv.Itoa(def), func(s string) error {
		_, err := strconv.ParseInt(s, 10, 64)
		return err
	})

	res, _ := strconv.ParseInt(rs, 10, 64)
	return int(res)
}

func askIntRange(question string, def, min, max int) int {
	rs := askString(question, strconv.Itoa(def), func(s string) error {
		n64, err := strconv.ParseInt(s, 10, 64)
		n := int(n64)
		if err != nil {
			return err
		}
		if n < min {
			return fmt.Errorf("Too Low : %d. Must be between %d and %d inclusive", n, min, max)
		}
		if n > max {
			return fmt.Errorf("Too High : %d. Must be between %d and %d inclusive", n, min, max)
		}
		return nil
	})

	res, _ := strconv.ParseInt(rs, 10, 64)
	return int(res)
}

func askOptions(question string, opts []string) int {
	q := question + "\n"
	for k, v := range opts {
		q += fmt.Sprintf("- %d - %s\n", k, v)
	}
	return askIntRange(q, 0, 0, len(opts)-1)
}

func askString(question, def string, filter func(string) error) string {
	for i := 0; i < 10; i++ {
		fmt.Println(question)
		fmt.Printf("Leave blank for Default : [%s]\n>", def)
		s := inputLine()
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}

		if filter != nil {
			err := filter(s)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
		return s

	}
	fmt.Println("Using:" + def)
	return def
}

func askMoney(question string, def string) money.M {
	var res money.M
	askAny(question, def, func(s string) error {
		r, err := money.Parse(s)
		if err != nil {
			return err
		}
		res = r
		return nil
	})
	return res
}

func askDate(question string, def time.Time) time.Time {
	var res time.Time
	askAny(question, def, func(s string) error {
		r, err := time.Parse("02/01/06", s)
		if err != nil {
			return err
		}
		res = r
		return nil
	})
	return res
}

//the recieved func, should either accept the result and put it where needed, or return an error
func askAny(question string, def string, fputter func(string) error) {
	for i := 0; i < 10; i++ {
		fmt.Println(question)
		fmt.Printf("Leave blank for Default : [%s]\n>", def)
		s := inputLine()
		s = strings.TrimSpace(s)
		if s == "" {
			fputter(def)
			return
		}

		err := fputter(s)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		return

	}
	fmt.Println("Using:" + def)
	fputter(def)
}
