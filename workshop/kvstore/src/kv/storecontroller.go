package main
import (
	"strings"
	"fmt"
)


func handleCommand(store map[string]string, arg string) bool {
	if strings.Contains(arg, "=") {

		kv := strings.SplitN(arg, "=", 2)
		store[kv[0]] = kv[1]
		return true

	} else {

		if v, exist := store[arg]; exist {
			fmt.Printf("%v=%v\n", arg, v)
		} else {
			fmt.Printf("%v NOT FOUND\n", arg)
		}
		return false
	}
}