package main

import (
	"fmt"
	"math"
)

func main() {

	//membulatkan ke atas
	fmt.Println(math.Ceil(1.40))

	//membulatkan ke bawah
	fmt.Println(math.Floor(1.60))

	//membulatkan ke yang terdekat
	fmt.Println(math.Round(1.60))

	//mencari angka terbesar
	fmt.Println(math.Max(10, 11))

	//mencari angka terkecil
	fmt.Println(math.Min(10, 11))

}
