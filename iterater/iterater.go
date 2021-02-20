package iteration

func Repeat(str string) string{
    var ret string
    for i := 0; i < 5; i++{
        ret = ret + str
    }
    return ret
}
