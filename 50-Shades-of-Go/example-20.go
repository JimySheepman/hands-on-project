package main

import "log"

func main() {
	log.Fatalln("Fatal Level: log entry") //app exits here
	log.Println("Normal Level: log entry")
}
