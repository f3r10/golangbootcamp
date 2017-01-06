package main

import (
  "fmt"
  "math/cmplx"
  "time"
)

type Bootcamp struct {
  // Latitude of the event
  Lat float64
  // Longitude of the event
  Lon float64
  // Date of the event
  Date time.Time
}

var (
  goIsFun bool       = true
  maxInt  uint64     = 1<<64 - 1
  complex complex128 = cmplx.Sqrt(-5 + 12i)
)

type User struct {
  Id             int
  Name, Location string
}

func (u *User) Greetings() string {
  return fmt.Sprintf("Hi %s from %s",
    u.Name, u.Location)
}

type Player struct {
  *User
  GameId int
}

func NewPlayer(id int, name, location string, gameId int) *Player {
  return &Player{
    User: &User{id, name, location},
    GameId: gameId,
  }
}

func main() {
  // const f = "%T(%v)\n"
  // fmt.Printf(f, goIsFun, goIsFun)
  // fmt.Printf(f, maxInt, maxInt)
  // fmt.Printf(f, complex, complex)

  // event := Bootcamp{
  //   Lat: 34.012836,
  //   Lon: -118.495338,
  // }
  // event.Date = time.Now()
  // fmt.Printf("Event on %s, location (%f, %f)",
  //   event.Date, event.Lat, event.Lon)

  player := NewPlayer(42, "f3r10", "Quito", 12234234)
  fmt.Println(player.Greetings())


}