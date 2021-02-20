package main

import "testing"
//import "fmt"

func Test_hello(t *testing.T){

    t.Run("1", func(t *testing.T){
        got := hello("12")
        want := "hello, 12"

        if(got != want){
            t.Errorf("got '%q' want '%q'", got, want)
        }
    })

    t.Run("2", func(t *testing.T){
        got := hello("")
        want := "hello, world"

        if(got != want){
            t.Errorf("got '%q' want '%q'", got, want)
        }
    })
}
