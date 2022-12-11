package main

import (
	"log"
)

func main() {
	me := newHuman().withEyeColour("brown").withAge(18).withHeigth(1.90).reset()
	you := newHumanWithFields(30, 1.70, "blue")

	log.Printf("%+v", me)
	log.Printf("%+v", you)

	youngGiant := giant().withAge(3)
	oldGiant := giant().withAge(150)

	log.Printf("%+v", youngGiant)
	log.Printf("%+v", oldGiant)

}

type human struct {
	age       int
	height    float64
	eyeColour string
}

func newHuman() human {
	return human{}
}

func newHumanWithFields(age int, height float64, eyeColour string) human {
	return human{
		age:       age,
		height:    height,
		eyeColour: eyeColour,
	}
}

func (h human) withEyeColour(colour string) human {
	h.eyeColour = colour

	return h
}

func (h human) withAge(age int) human {
	h.age = age

	return h
}

func (h human) withHeigth(height float64) human {
	h.height = height

	return h
}

func (h human) reset() human {
	return human{}
}

func giant() human {
	return newHuman().withHeigth(2.5).withEyeColour("green")
}
