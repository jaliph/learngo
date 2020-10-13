package main

import "fmt"

var privatemap map[string]string

func init() {
	fmt.Println("Always called...")
	privatemap = make(map[string]string)
}

func addtoMap(k, v string) {
	if _, ok := privatemap[k]; ok {
		fmt.Println("\nKey already exists, can't add")
	} else {
		fmt.Println("\nKey doesn't exists, adding it")
		privatemap[k] = v
	}
}

func deletefromMap(k string) {
	if _, ok := privatemap[k]; ok {
		fmt.Println("\nKey exists.. deleting it")
		delete(privatemap, k)
	} else {
		fmt.Println("\nKey doesn't exists, adding it")
	}
}

func main() {
	addtoMap("foo", "bar")
	addtoMap("foo", "bar")
	addtoMap("foo1", "bar1")
	deletefromMap("foo1")
	for k, v := range privatemap {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}
