package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Count() int {
	return len(d)
}

func main() {

	var dict = Dictionary{
		"keyOne":   "value",
		"keyTwo":   "anotherValue",
		"keyThree": "yetAnotherValue",
	}
	fmt.Printf("There are %d elements in dictionary \n", dict.Count())
	searchKey := "nonExistingKey"
	found, err := dict.Search(searchKey)

	if errors.Is(err, ErrNotFound) {
		fmt.Printf("There is no such key : %s \n", searchKey)
	} else {
		fmt.Printf("The key was found %s \n", found)
	}
}
