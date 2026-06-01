package main

import "flag"

func main() {
	max := flag.Int("max",255,"max value")
	name := flag.String("name","something","myname")

	flag.Parse()

	println(*name,*max)

}