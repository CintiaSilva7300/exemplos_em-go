package main

import "fmt"

func main() {
	var minhaVar interface{} = "Cintia silva"
	fmt.Println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Println(res, ok)

	res2 := minhaVar.(int)
	fmt.Println(res2)
}
