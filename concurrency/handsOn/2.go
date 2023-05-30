package main

type person struct{}

func (p *person) speak() {
	println("I am a person")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	saySomething(&person{})
	saySomething(person{})
}
