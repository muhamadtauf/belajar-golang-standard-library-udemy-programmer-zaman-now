package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	var data *ring.Ring = ring.New(5)

	data.Value = "Value 1"

	data = data.Next()
	data.Value = "Value 2"

	data = data.Next()
	data.Value = "Value 3"

	data = data.Next()
	data.Value = "Value 4"

	data = data.Next()
	data.Value = "Value 5"

	data.Do(func(value any) {
		fmt.Println(value)
	})

	//dengan perulangan
	var data2 *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data2.Value = "Value For " + strconv.Itoa(i)
		data2 = data2.Next()
	}
	data2.Do(func(value any) {
		fmt.Println(value)
	})

}
