package main

import "testing"
import "math"

func Test_Area(t *testing.T){
    check := func(t *testing.T, s shape, want float64){
        got := s.Area()
        if got != want {
            t.Errorf("got %f want %f", got, want)
        }
    }

    t.Run("rec", func (t *testing.T){
        r := rec{12.0, 6.0}
        want := 72.0
        
        check(t, r, want)
    })

    t.Run("circle", func (t *testing.T){
        c := circle{2.0}
        want := 2.0 * 2.0 * math.Pi

        check(t, c, want)
    })
}
func Test_per(t *testing.T){
    got := perimeter(10.0, 10.0)
    want := 40.0

    if( got != want){
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
