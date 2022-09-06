package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Printf("%-16v %10v %-11v %5v\n", "Spaceline", "Days", "Trip type", "Price")
    fmt.Println("=============================================")

    numOfTickets := 10
    for numOfTickets > 0 {
        // Random spaceline
        companyName := "Virgin Galactic"
        switch companyIndex := rand.Intn(3); companyIndex {
        case 0:
            companyName = "Space Adventures"
        case 1:
            companyName = "SpaceX"
        }

        // Random trip duration
        minSpeed := 16
        maxSpeed := 30
        speed := rand.Intn(maxSpeed-minSpeed+1) + minSpeed
        marsDistance := 62_100_000
        days := marsDistance / speed / 60 / 60 / 24

        // Random trip type
        tripType := "One-way"
        if rand.Intn(2) == 1 {
            tripType = "Round-trip"
        }

        // Calculate price
        price := speed + 20
        if tripType == "Round-trip" {
            price *= 2
        }

        fmt.Printf("%-16v %10v %-11v $ %3v\n", companyName, days, tripType, price)
        numOfTickets--
    }
}
