package main

import "fmt"

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	HatchCount    int
	CreateReptile ReptileCreator
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.HatchCount == 0 {
		egg.HatchCount++
		return egg.CreateReptile()
	}
	return nil
}

type FireDragon struct {
}

func (f FireDragon) Lay() ReptileEgg {
	reptileEgg := ReptileEgg{
		CreateReptile: func() Reptile {
			fireDragon := FireDragon{}
			return fireDragon
		},
		HatchCount: 0,
	}
	return reptileEgg
}

func main() {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch()
	fmt.Printf("%+v", childDragon)
}
