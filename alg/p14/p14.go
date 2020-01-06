package p14

import (
    "strings"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 || strs == nil{
        return ""
    }
    var count int = 0
    minLen,index := GetMinLen(strs) 
    mlstr := strs[index]
    for i := minLen; 0 < i; i--{
        for i,v := range strs{
            if i == index{ continue }
            if strings.HasPrefix(string(v),mlstr) {
                count++
            } else { break }
        }
        if count == len(strs) - 1{
            break
        } else {
            count = 0
            mlstr = mlstr[0:i-1]
        }
    }
    return mlstr
}

func GetMinLen(strs []string) (min,index int){
    min = len(strs[0])
    for i,v := range strs{
        if len(v) < min{
            min = len(v)
            index = i
        }
    }
    return min,index
}


