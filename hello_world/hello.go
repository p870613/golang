package main

import "fmt"

func hello(name string) string{
    if(name == ""){
        return "hello, world"
    }
    return "hello, " + name
}
func main(){
    fmt.Println(hello("123"))
}
