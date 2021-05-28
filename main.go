// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that returns what you should do

package main

import (
	"fmt"
)

// This main function will ask user for age if it is the weekday
func main() {
	// Defining variables
	var age int
	var weekday bool

	fmt.Println("Should you go to the movies?")
	fmt.Println()

	// User Input
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	fmt.Println()

	fmt.Print("Is it a weekday(true or false): ")
	fmt.Scanln(&weekday)
	fmt.Println()

	// Return what the user should do
	if weekday {
		if age < 18 {
			fmt.Println("Time to go to School!")
		} else if age >= 18 {
			fmt.Println("Time to go to work")
		}
	} else {
		fmt.Println("Its the weekened just relax!")
	}
}
