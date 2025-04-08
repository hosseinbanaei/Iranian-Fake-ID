package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function generates a fake 10-digit national ID
func generateFakeID() string {
	// Generate the first 9 digits randomly
	var digits [9]int
	for i := 0; i < 9; i++ {
		digits[i] = rand.Intn(10)
	}

	// Calculate the check digit (tenth digit)
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	control := 11 - (sum % 11)
	if control >= 10 {
		control = 0
	}

	// Create the final string (complete national ID)
	ID := ""
	for i := 0; i < 9; i++ {
		ID += fmt.Sprint(digits[i])
	}
	ID += fmt.Sprint(control)

	return ID
}

func main() {
	// Set randomness based on time

	/* "rand.Seed(time.Now().UnixNano())" is deprecated.
	As of Go 1.20 there is no reason to call Seed with a random value.
	Programs that call Seed with a known value to get a specific sequence of results should use
	New(NewSource(seed)) to obtain a local random generator. */

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var n int
	fmt.Print("How many fake national IDs should be generated? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("The number entered is invalid.")
		return
	}

	fmt.Println("National IDs generated:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, generateFakeID())
	}
}
