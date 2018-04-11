package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputLine() string {
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	return text

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
