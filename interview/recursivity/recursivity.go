package main

import (
	"fmt"
)
//printNoDesc prints digits in descendant order from 10 to given id. 
func printNoDesc(id int) {
	if id > 10 {
		return
	}
	printNo(id+1)
	fmt.Println(id)
}

//printNoDesc prints digits in ascendant order from 10 to given id.
func printNoAsc(id int) {
	if id > 10 {
		return
	}
	fmt.Println(id)
	printNo(id+1)
}

func main() {
	printNoDesc(0)
  printNoAsc(0)
}
