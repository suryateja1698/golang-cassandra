package main

import (
	"github.com/suryateja1698/golang-cassandra/db"
)

func main() {
	conn := db.Init()
	defer conn.Close()
}
