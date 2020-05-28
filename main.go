package main

import (
	"dictionary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()
	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknow action : %v\n", *action)
	}
}


func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("%v was removed from the dictionary", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	get, err := d.Get(word)
	handleErr(err)
	fmt.Println(get)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to dictionnary \n", word)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error :%v\n", err)
		os.Exit(1)
	}
}
