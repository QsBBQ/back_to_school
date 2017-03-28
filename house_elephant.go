// App to inventory a house
// Requirements
// Need to be able to import house stuff
// Need to be able to add a device
// Need to be able to remove a device
// Need to be able to report on total
// category?

package main
import "fmt"
import "time"

type House struct {
  Things []Thing
}

func (h *House) blah() int {
  return "blah"
}

type Thing struct {
  id int
  cost int64
  description string
  dateinservice time
  datedeposed time
  broken []
  tag string
  category
}

type Appliance interface {

}
type Appliance struct {
}

type Refrigerator struct {
  waterfilter
  Thing
}

func (d *Refrigerator) temperature() int {
  // get temperature
  return 35
}

func (d *Refrigerator) temperatureFreezer() int {
  // get temperature
  return 0
}

type Entertainment struct {}
type Tv struct {}
type SoundBar struct {}
type GameSystem struct {}
type HouseSensor struct {}
type DogFeeder struct {}
type Catfeeder struct {}

housethings = {"refrigerator": [{"price": 500}, {"brand": "samsung"}, {"type": "appliance"}],
               "dishwasher": [{"price": 300}, {"brand": "samsung"}, {"type": "appliance"}],
               "TV": [{"price": 800}, {"brand": "LG"}, {"type": "entertainment"}]}

func main() {


}
