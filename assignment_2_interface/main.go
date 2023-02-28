package main

import "log"

type shape interface {
	getArea() (float64, error)
}

type triangle struct { 
	height float64
	base float64
}
type square struct { 
	sideLength float64
}

func main()  {
    sq := square{sideLength: 10}
    tr := triangle{base: 10, height: 10}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() (float64, error) {
    return 0.5 * t.base * t.height, nil
}

func (s square) getArea() (float64, error) {
    return s.sideLength * s.sideLength, nil
}

func printArea(s shape)  {
	log.Println(s.getArea())
}