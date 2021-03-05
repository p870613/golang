package main

func Sum(number []int) int{
    sum := 0
    for _, i := range number{
        sum = i + sum
    }
    return sum
}

//func SumAll(nums ...[]int)(ret []int){
    //length := len(nums)
    //ret = make([]int, length)

    //for i, num := range nums {
        //ret[i] = Sum(num)
    //}

    //return
//}
func SumAll(nums ...[]int)([]int){
    var ret []int

    for _, num := range nums {
        ret = append(ret, Sum(num))
    }

    return ret
}
