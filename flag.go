package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database usernmae")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	//jika file dijalankan tanpa flag, flagnya akan mengambil dari default value

	//contoh command
	// go run flag.go -username test
	//jika value labih dari satu kata, bisa menggunakan ""
	//contoh command
	// go run flag.go -username "test test"
}
