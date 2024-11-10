package main

type Animals interface {
	makeSound() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) makeSound() string {
	return "Woff!"
}

func (c Cat) makeSound() string {
	return "Meow!!"
}

func main() {

}
