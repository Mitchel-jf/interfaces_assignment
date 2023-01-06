package main

import (
	"github.com/mitchel-jf/interfaces_assignment/shapes"
)

func main() {
	square := shapes.Square{SideLength: 23}
	triangle := shapes.Triangle{Base: 10, Height: 20}
	shapes.PrintArea(&square)
	shapes.PrintArea(&triangle)
}
