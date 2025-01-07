package main

import (
	"log"

	"simplemath.com/utils"
	"test.com/pckg/mathutils"
)

func main() {
	res := mathutils.Add(1, 2)
	log.Println("res is ", res)

	sq := utils.Square(1)
	log.Println("sq is ", sq)

}
