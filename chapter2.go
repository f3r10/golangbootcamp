package main

import "fmt"

const (
  Pi = 3.14
  Truth = false
  Big = 1 << 62 
  Small = Big >> 61
)

type Artist struct {
  Name, Genre string
  Songs       int
}

func newRelease(a Artist) int {
  a.Songs++
  return a.Songs
}

func main() {
  name, city := "Fernando Ledesma", "Quito"
  age := 26

  action := func () {
    fmt.Printf("%s", "test anonymous func")
  }

  fmt.Printf("%s (%d) of %s", name, age, city)

  action()

  fmt.Println(Pi)

  continent, region := location(name, city)
  fmt.Println(continent)
  fmt.Println(region)

  me := Artist{Name: "Matt", Genre: "Electro", Songs: 42}
  fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
  fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)

}

func location(name, city string) (continent, region string) {
  switch city {
    case "Quito":
      continent = "America"
      region = "SouthAmerica"
    default:
      continent = "Unknown"
      region = "Unknown"
  }
  return
}