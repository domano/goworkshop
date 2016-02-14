package main

import (
	"fmt"
	"os"
)


func main() {
	kvstore:=NewStore()
	file := "store_simple.db"
	store, err := kvstore.ReadStore(file)
	if err != nil {
		fmt.Printf("error opening store: %v", err.Error())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		kvstore.WriteTo(store, os.Stdout)
		os.Exit(0)
	}

	shouldSave := false
	for _, arg := range os.Args[1:] {
		if modified := handleCommand(store, arg); modified {
			shouldSave = modified
		}
	}

	if shouldSave {
		if err := kvstore.WriteStore(store, file); err != nil {
			fmt.Printf("error writing store: %v", err.Error())
			os.Exit(2)
		}
	}
}


