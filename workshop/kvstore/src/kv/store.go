package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
)

type Store struct {

}

func NewStore() *Store {
	return &Store{}
}

func (s Store)ReadStore(file string) (map[string]string, error) {
	store := make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		handleCommand(store, scanner.Text())
	}

	return store, scanner.Err()
}

func (s Store)WriteStore(store map[string]string, file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return s.WriteTo(store, f)
}

func (s Store)WriteTo(store map[string]string, w io.Writer) error {
	for k, v := range store {
		if _, err := fmt.Fprintf(w, "%v=%v\n", k, v); err != nil {
			return err
		}
	}
	return nil
}