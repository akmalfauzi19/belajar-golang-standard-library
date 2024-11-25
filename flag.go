package main

import (
	"flag"
)

func main() {
	var username *string = flag.String("username", "admin", "The username of the user")
	var password *string = flag.String("password", "admin", "The password of the user")
	var host *string = flag.String("host", "localhost", "The host of the server")
	var port *int = flag.Int("port", 8080, "The port of the server")

	flag.Parse()

	println("Username:", *username)
	println("Password:", *password)
	println("Host:", *host)
	println("Port:", *port)
}
