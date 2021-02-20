package iteration

import "testing"

func Test_iteration(t *testing.T){
    get := Repeat("a")
    expected  := "aaaaa"

    if get != expected {
        t.Errorf("expext '%q' but got '%q'", expected, get)
    }
}

func Benchmark_iterator(b *testing.B){
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
