package main

import "fmt"

func main() {

	var s [5]string = [5]string{"i", "am", "stupid", "and", "weak"}

	for index := range s {
		if s[index] == "stupid" {
			s[index] = "smart"
		}
		if s[index] == "weak" {
			s[index] = "strong"
		}
	}
	fmt.Println(s)

}
