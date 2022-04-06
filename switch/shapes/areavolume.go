package main

import (
	"fmt"
	"math"
)

func main() {

	x := 0
	for {
		fmt.Println("*********** Main Menu **********")
		fmt.Println("Enter 1 for Area and 2 for Volume (0 to exit)")
		if _, err := fmt.Scan(&x); err == nil {
			switch x {
			case 1:
				y := 0
				for {
					fmt.Println("********** Area **********")
					fmt.Println("Menu: [1] Rectangle, [2] Circle, [9] Main Menu, [0] Exit")
					if _, err := fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							areaRectangle()
						case 2:
							areaCircle()
						case 9:
							goto exitArea
						case 0:
							goto outerloop
						default:
							fmt.Println("Not a valid value. Enter again.")
						}
					}
				}
			exitArea:
				break
			case 2:
				y := 0
				for {
					fmt.Println("********** Volume **********")
					fmt.Println("Menu: [1] Cylinder, [2] Cube, [9] Main Menu, [0] Exit")
					if _, err := fmt.Scan(&y); err == nil {
						switch y {
						case 1:
							volumeCylinder()
						case 2:
							volumeCube()
						case 9:
							goto exitVolume
						case 0:
							goto outerloop
						default:
							fmt.Println("Not a valid value. Enter again.")
						}
					}
				}
			exitVolume:
				break
			case 0:
				goto outerloop
			default:
				fmt.Println("Not a valid value. Enter again.")
			}
		}
	}
outerloop:
}

func areaCircle() {
	var r float32
	fmt.Println("Enter the radius of the circle:")
	if _, err := fmt.Scan(&r); err == nil {
		fmt.Printf("The area of the circle of radius %5.2f is: %5.2f\n", r, math.Pi*r*r)
	}
}

func areaRectangle() {
	var l, w float32
	fmt.Println("Enter the length and width of the rectangle:")
	if _, err := fmt.Scan(&l, &w); err == nil {
		fmt.Printf("The area of a %5.2f x %5.2f rectangle is: %5.2f\n", l, w, 1*w)
	}
}

func volumeCylinder() {
	var r, h float32
	fmt.Println("Enter the radius and height of the cylinder:")
	if _, err := fmt.Scan(&r, &h); err == nil {
		fmt.Printf("The volume of a %5.2f x %5.2f cylinder is: %5.2f\n", r, h, math.Pi*r*r*h)
	}
}

func volumeCube() {
	var e float32
	fmt.Println("Enter the length of a side of the cube:")
	if _, err := fmt.Scan(&e); err == nil {
		fmt.Printf("The volume of a %5.2f x %5.2f cube is: %5.2f\n", e, e, e*e*e)
	}
}
