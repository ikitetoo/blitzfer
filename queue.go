package main

import (
	"log"
)

func passiveQueue() {
	for {
		dir := <- pq
		log.Printf("Got [%v]", dir.path)
//		go scanDir(dir)
	}
}

/*
func activeQueue() {
	
}
*/
