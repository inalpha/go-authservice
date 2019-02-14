package main

import "github.com/gofrs/uuid"

func main() {
	var u1 = uuid.Must(uuid.NewV4())
	println(u1.String())
}
