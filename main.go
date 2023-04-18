package main

import "fmt"

func howManyElements(data []any) int {
	//your code here!
	return len(data)
}

func main() {
	data1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(howManyElements(data1))

	data2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(data2))

	data3 := []interface{}{1, 1, 1, 5, 5, 5}
	fmt.Println(howManyElements(data3))

	data4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(data4))

	data5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(data5))

	data6 := []interface{}{true, false, true, false, true, false}
	fmt.Println(howManyElements(data6))
}
