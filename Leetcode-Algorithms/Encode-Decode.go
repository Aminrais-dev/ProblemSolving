package main

import (
	"fmt"
	"strings"
)

/*
Buatlah sebuah program untuk encode dan decode sebuah string dengan shift sebanyak 2 karakter
*/

func encodeDecode(word string) string {

	split := strings.Split(word, ">")
	if split == nil {
		return "mohon masukan perintah (<decode>/<encode>)"
	}

	var perintah = split[0]
	var str = split[1]
	var shift = 2

	var res string
	if perintah == "<encode" {
		for _, v := range str {
			encode := v + rune(shift)
			if encode > 122 {
				encode -= 26
				res += string(encode)
			} else {
				res += string(encode)
			}
		}
	} else if perintah == "<decode" {
		for _, v := range str {
			decode := v - rune(shift)
			if decode < 97 {
				decode += 26
				res += string(decode)
			} else {
				res += string(decode)
			}
		}
	} else {
		return "mohon format perintah sesuai dengan (<encode>/<decode>)"
	}

	return res

}

func main() {

	//encode
	fmt.Println(encodeDecode("<encode>abcd")) //cdef
	fmt.Println(encodeDecode("<encode>xyz"))  //zab

	//decode
	fmt.Println(encodeDecode("<decode>cdef")) //abcd
	fmt.Println(encodeDecode("<decode>zab"))  //xyz

}
