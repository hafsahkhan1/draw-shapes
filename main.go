// hafsah khan
// cs 341 hw 5
// practicing go interfaces by creating user-defined images
// main: contains control loop to direct user flow

package main

import (
	// "crypto/x509"
	"fmt"
)

var display Display

func main() {
	// Print intro message
	fmt.Println("Homework 5: Geometry Using Go Interfaces")
	fmt.Println("CS 341, Fall 2025")
	fmt.Println("")
	fmt.Println("This application allows you to draw various shapes")
	fmt.Println("of different colors via interfaces in Go.")
	fmt.Println("")

	//
	// SOME PRINT STATEMENTS YOU WILL NEED CAN BE FOUND BELOW
	//
	// Ask user for dimensions for display
	numRows := 0
	numCols := 0
	fmt.Print("Enter the number of rows (x) that you would like the display to have: ")
	fmt.Scanln(&numRows)
	fmt.Print("Enter the number of columns (y) that you would like the display to have: ")
	fmt.Scanln(&numCols)
	display.initialize(numRows, numCols)
	//
	// Menu options

	input := "A"
	for input != "X" {
		fmt.Println("")
		fmt.Println("Select a shape to draw: ")
		fmt.Println("	 R for a rectangle")
		fmt.Println("	 T for a triangle")
		fmt.Println("	 C for a circle")
		fmt.Println(" or X to stop drawing shapes.")
		fmt.Print("Your choice --> ")
		fmt.Scanln(&input)

		if input == "X" {
			break
		} else if input == "R" {
			//
			// Drawing a rectangle
			var xVal, yVal int
			color := ""

			fmt.Print("Enter the X and Y values of the lower left corner of the rectangle: ")
			fmt.Scanln(&xVal, &yVal)
			ll := Point{xVal, yVal}

			fmt.Print("Enter the X and Y values of the upper right corner of the rectangle: ")
			fmt.Scanln(&xVal, &yVal)
			ur := Point{xVal, yVal}

			fmt.Print("Enter the color of the rectangle: ")
			fmt.Scanln(&color)
			element, found := colors[color]
			if !found {
				element = Color{1, 2, 3}
			}

			rect := Rectangle{ll, ur, element}
			fmt.Println(rect.printShape())
			err := rect.draw(&display)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Rectangle drawn successfully.")
			}
		} else if input == "T" {
			//
			// Drawing a triangle
			var xVal, yVal int
			color := ""

			fmt.Print("Enter the X and Y values of the first point of the triangle: ")
			fmt.Scanln(&xVal, &yVal)
			p0 := Point{xVal, yVal}

			fmt.Print("Enter the X and Y values of the second point of the triangle: ")
			fmt.Scanln(&xVal, &yVal)
			p1 := Point{xVal, yVal}

			fmt.Print("Enter the X and Y values of the third point of the triangle: ")
			fmt.Scanln(&xVal, &yVal)
			p2 := Point{xVal, yVal}

			fmt.Print("Enter the color of the triangle: ")
			fmt.Scanln(&color)
			element, found := colors[color]
			if !found {
				element = Color{1, 2, 3}
			}

			tri := Triangle{p0, p1, p2, element}
			fmt.Println(tri.printShape())
			err := tri.draw(&display)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Triangle drawn successfully.")
			}
		} else if input == "C" {
			//
			// Drawing a circle
			var xVal, yVal, rad int
			color := ""

			fmt.Print("Enter the X and Y values of the center of the circle: ")
			fmt.Scanln(&xVal, &yVal)
			center := Point{xVal, yVal}

			fmt.Print("Enter the value of the radius of the circle: ")
			fmt.Scanln(&rad)

			fmt.Print("Enter the color of the circle: ")
			fmt.Scanln(&color)
			element, found := colors[color]
			if !found {
				element = Color{1, 2, 3}
			}

			circ := Circle{center, rad, element}
			fmt.Println(circ.printShape())
			err := circ.draw(&display)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Circle drawn successfully.")
			}
		} else {
			fmt.Println("**Error, unknown command. Try again.")
			continue
		}
	}
	//
	// Saving the results in a file
	filename := "drawing"
	fmt.Print("Enter the name of the .ppm file in which the results should be saved: ")
	fmt.Scanln(&filename)
	display.screenShot(filename)
	//
	// Exiting program
	fmt.Println("Done. Exiting program...")
}
