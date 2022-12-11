package builder

type human struct {
	age       int
	height    float64
	eyeColour string
}

func NewHuman() human {
	return human{}
}

func NewHumanWithFields(age int, height float64, eyeColour string) human {
	return human{
		age:       age,
		height:    height,
		eyeColour: eyeColour,
	}
}

func (h human) WithEyeColour(colour string) human {
	h.eyeColour = colour

	return h
}

func (h human) WithAge(age int) human {
	h.age = age

	return h
}

func (h human) WithHeight(height float64) human {
	h.height = height

	return h
}

func (h human) Reset() human {
	return human{}
}

func Giant() human {
	return NewHuman().WithHeight(2.5).WithEyeColour("green")
}
