package main
import "fmt"

type Animal interface {
 Name() string
}

type PartyAnimal interface {
 Animal
 Party()
}

type Dog struct {}

func (d *Dog) Name() string {
 return "Dog"
}

func (d *Dog) Bark() {
 fmt.Println("Woof!")
}

type GuideDog struct {
 *Dog
}

 // removed Help(h *Human) ??
func (gd *GuideDog) Help(humanName string) {
 fmt.Printf("Hey human %s, grab %s’s leash!\n", humanName, gd.Name())
}

type Cat struct {}
func (c *Cat) Name() string {
 return "Cat"
}
type CatDog struct {
 *Cat
 *Dog
}

func (cd *CatDog) Name() string {
 return "CatDog"
}

func main() {
 var animal Animal
 animal = &Dog{} // returns a pointer to a new Dog
 fmt.Println(animal.Name()) // Dog
 // What is the difference between this and using interfaces?
 myDog := Dog{}
 fmt.Println(myDog.Name())

 gd := &GuideDog{}
 gd.Help("Bob") // prints “Hey human, grab Dog’s leash!”

 cd := &CatDog{}
 fmt.Printf("My favorite animal is the %s!\n", cd.Name())
}
