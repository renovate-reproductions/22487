package main

import "github.com/fifsky/goconf"

func main() {
	_, err := goconf.NewConfig("")
	panic(err)
}
