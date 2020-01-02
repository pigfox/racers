package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"time"
)

var racers = make(map[int]string)
var distance = 200
var tracks = 10
var maxHops = 8

func main() {
	for d := 1; d <= distance; d++ {
		for i := 1; i <= tracks; i++ {
			racers[i] += run(maxHops)
			callClear()
			display()
			if distance <= len(racers[i]){
				fmt.Println("\nNumber [" + strconv.Itoa(i) + "] wins!!!\n")
				os.Exit(1)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func display(){
	for i := 1; i <= tracks; i++  {
		fmt.Println(racers[i] + "[" + strconv.Itoa(i) + "]")
	}
}

func run(i int) string{
	var padding string
	rand.Seed(time.Now().UnixNano())
	pad := rand.Intn(i)
	for k := 1; k < pad; k++  {
		padding += "X"
	}

	return padding
}