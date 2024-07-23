package main

import (
	"fmt"
	"privLessons/db"
)

func main() {
    fmt.Println("PrivLessons")
    db.ConnectDB()
}
