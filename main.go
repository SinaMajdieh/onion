package main

import (
	"fmt"

	"github.com/SinaMajdieh/onion/egg"
)

func main() {
	var inp string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s", &inp)
		if inp == "create_layer" {
			create_layer()
		} else if inp == "layers" {
			egg.Print_layers()
		} else if inp == "do_magic" {
			fmt.Println("abracadabra b!tch")
		} else if inp == "exit" {
			return
		} else {
			fmt.Println("you onion! that not a command")
		}
	}

}

func create_layer() {
	var name, addr string
	var w, h, x, y int
	fmt.Print("layer name > ")
	fmt.Scanf("%s", &name)
	fmt.Print("size > ")
	fmt.Scanf("%d %d", &w, &h)
	fmt.Print("starting at > ")
	fmt.Scanf("%d %d", &x, &y)
	fmt.Print("address > ")
	fmt.Scanf("%s", &addr)
	layer := egg.New_layer(name, egg.Pair(w, h), egg.Pair(x, y), addr)
	egg.Add_layer(&layer)

}
