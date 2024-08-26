package main

import (
	"fmt"
	"time"
)

// There are two categories of cars: SUV and hatchback.
// Maintain a count of how many SUV and hatchback cars enter the premise

// Calculate the payment each car has to make based upon the rates as hatchback parking as 10 rupees per hour and SUV being 20 rupees an hour.
// In case if hatchback occupancy is full then hatchback cars can occupy SUV spaces but with hatchback rates.
// During exit there needs to be the system to inform the user how much they have to pay
// Admin can see all the cars which are parked in the system

type Car struct {
	Category  string
	EntryTime time.Time
}

type Parking struct {
	SUVCount, HatchbackCount int
	SUVLimit, HatchbackLimit int
	Cars                     []Car
}

func NewParking(suvLimit, hatchbackLimit int) *Parking {
	return &Parking{SUVLimit: suvLimit, HatchbackLimit: hatchbackLimit}
}

func (p *Parking) AddCar(category string) error {
	if category == "SUV" {
		if p.SUVCount < p.SUVLimit {
			p.SUVCount++
		} else {
			return fmt.Errorf("SUV parking full")
		}
	} else if category == "Hatchback" {
		if p.HatchbackCount < p.HatchbackLimit {
			p.HatchbackCount++
		} else if p.SUVCount < p.SUVLimit {
			p.SUVCount++
		} else {
			return fmt.Errorf("Hatchback and SUV parking full")
		}
	}
	p.Cars = append(p.Cars, Car{Category: category, EntryTime: time.Now()})
	return nil
}

func (p *Parking) CalculatePayment(car Car, exitTime time.Time) int {
	duration := exitTime.Sub(car.EntryTime).Hours()
	if car.Category == "SUV" {
		return int(duration) * 20
	} else {
		return int(duration) * 10
	}
}

func (p *Parking) RemoveCar(index int) (int, error) {
	if index >= len(p.Cars) {
		return 0, fmt.Errorf("invalid index")
	}
	car := p.Cars[index]
	p.Cars = append(p.Cars[:index], p.Cars[index+1:]...)
	if car.Category == "SUV" {
		p.SUVCount--
	} else {
		p.HatchbackCount--
	}
	return p.CalculatePayment(car, time.Now()), nil
}

func (p *Parking) DisplayAllCars() {
	fmt.Println("Cars currently in parking:")
	for i, car := range p.Cars {
		fmt.Printf("Car %d: %s, Entered at: %v\n", i+1, car.Category, car.EntryTime)
	}
}

func main() {
	parking := NewParking(5, 5)
	parking.AddCar("SUV")
	parking.AddCar("Hatchback")
	parking.AddCar("Hatchback")
	parking.DisplayAllCars()
	payment, err := parking.RemoveCar(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Payment for removed car: %d\n", payment)
	}
	parking.DisplayAllCars()
}
