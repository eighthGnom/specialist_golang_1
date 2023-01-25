package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	tracker := New()
	for scaner.Scan() {
		t := scaner.Text()
		arguments := strings.Split(t, " ")
		command := arguments[0]
		switch {
		case command == "Add" && tracker.onOff:
			tracker.Add(arguments[1], arguments[2])
		case command == "Del" && tracker.onOff:
			l := len(arguments)
			if l == 2 {
				fmt.Println(tracker.DelAll(arguments[1]))
			} else if l == 3 {
				fmt.Println(tracker.DelOne(arguments[1], arguments[2]))
			}
		case command == "Find" && tracker.onOff:
			fmt.Println(tracker.Find(arguments[1]))
		case command == "Print" && tracker.onOff:
			tracker.Print()
		case command == "StartApp":
			tracker.StartApp()
		case command == "Quit":
			tracker.Quit()

		}
	}
}
