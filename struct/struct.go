package main

import "math"

type shape interface{
    Area() float64
}
type rec struct {
    w float64
    h float64
}

type circle struct{
    r float64
}

func perimeter(w float64, h float64) float64{
    return 2 * (w + h)
}


func (r rec) Area() float64{
    return r.w * r.h
}

func (c circle) Area() float64{
    return c.r * c.r * math.Pi
}
