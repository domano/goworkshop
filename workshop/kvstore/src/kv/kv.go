package main

import (
	"fmt"
	"os"
	"sync"
	"net/http"
)

var mutex =&sync.Mutex{}
const filename="httpkvstore.txt"

func main() {
	//mainroutine(os.Args, "kvstore.db")
	waitGroup:=new(sync.WaitGroup)

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		key:=req.URL.Path[1:]
		switch req.Method {
		case "GET":
			args:=[]string{"",key}
			mainRoutineWithWaitGroup(args, waitGroup, filename)
		case "POST":
			post:=[]byte{}
			req.Body.Read(post)
			poststring:= string(post)
			fmt.Printf(post)
			fmt.Printf(poststring+"\n")
			args:=[]string{"",key+"="+poststring}
			mainRoutineWithWaitGroup(args, waitGroup, filename)

		}
	})
	http.ListenAndServe(":1337", nil)
}



func mainRoutineWithWaitGroup(args []string, waitGroup *sync.WaitGroup, filename string){
	waitGroup.Add(1)
	mainroutine(args, filename)
	waitGroup.Done()
}

func mainroutine(args []string, file string){
	mutex.Lock()
	kvstore:=NewStore()
	store, err := kvstore.ReadStore(file)
	if err != nil {
		fmt.Printf("error opening store: %v", err.Error())
		os.Exit(1)
	}

	if len(args) == 1 {
		kvstore.WriteTo(store, os.Stdout)
		os.Exit(0)
	}

	shouldSave := false
	for _, arg := range args[1:] {
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
	mutex.Unlock()
}



