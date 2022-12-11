package main

import (
	"design-patterns/builder"
	"log"
)

func main() {
	//builder examples
	me := builder.NewHuman().WithEyeColour("brown").WithAge(18).WithHeight(1.90).Reset()
	you := builder.NewHumanWithFields(30, 1.70, "blue")
	youngGiant := builder.Giant().WithAge(3)
	oldGiant := builder.Giant().WithAge(150)

	log.Printf("%+v", me)
	log.Printf("%+v", you)
	log.Printf("%+v", youngGiant)
	log.Printf("%+v", oldGiant)

	//factory examples
}
