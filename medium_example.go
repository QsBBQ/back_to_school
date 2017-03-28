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

type Human struct {}
func (h *Human) Name(humanName string) string {
 return humanName
}

 // removed Help(h *Human) ??
 // For some reason I can declare the method a string?
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

func PrintName(a Animal) {
  fmt.Printf("%s\n", a.Name())
}

func main() {
 var animal Animal
 animal = &Dog{} // returns a pointer to a new Dog
 fmt.Println(animal.Name()) // Dog
 // What is the difference between this and using interfaces?
 animal = &Cat{}
 fmt.Println(animal.Name())

 gd := &GuideDog{}
 gd.Help("Bob") // prints “Hey human, grab Dog’s leash!”

 cd := &CatDog{}
 fmt.Printf("My favorite animal is the %s!\n", cd.Name())

 PrintName(&Dog{})
 PrintName(&CatDog{})

 var animals []Animal
 animals = append(animals,&Cat{})
 animals = append(animals,&Dog{})
 animals = append(animals,&CatDog{})
 for _,animal:=range animals{
   PrintName(animal)
 }

}
