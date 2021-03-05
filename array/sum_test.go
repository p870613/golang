package main

import "testing"
import "reflect"

func Test_sum(t *testing.T){

    t.Run("test case 1", func(t *testing.T){
        number := []int{1, 2, 3, 4, 5}

        got := Sum(number)
        want := 15

        if(want != got){
            t.Errorf("got %d want %d given, %v", got, want, number)
        }

    })


    t.Run("test case 2", func(t *testing.T){
        number := []int{1, 2, 3}

        got := Sum(number)
        want := 6

        if(want != got){
            t.Errorf("got %d want %d given, %v", got, want, number)
        }

    })
}

func Test_SunAll(t *testing.T){
    
    check := func(t *testing.T, got []int, want []int){
        if !reflect.DeepEqual(got, want){
            t.Errorf("got %v want %v", got, want)
        }
    }
    
    t.Run("test case 1", func(t *testing.T){
        got := SumAll([]int{1, 2}, []int{0, 9})
        want := []int{3, 9}

        //if(!reflect.DeepEqual(got, want)){
            //t.Errorf("got %v want %v", got, want)
        //}
        check(t, got, want)
    })
        
    
    t.Run("test case 2", func(t *testing.T){
        got := SumAll([]int{1, 2}, []int{0, 9})
        want := []int{3, 9}

        check(t, got, want)
        //if(!reflect.DeepEqual(got, want)){
            //t.Errorf("got %v want %v", got, want)
        //}
    })

}
