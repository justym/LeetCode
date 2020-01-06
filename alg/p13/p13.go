package p13
/*import "fmt"*/

func romanToInt(s string) int {
    //Init variable
    var rev,num int = 0,0
    var next int = 0
    var curSymbol,nextSymbol string
    RomanIntMap := make(map[string]int,7)
    RomanIntMap = InitMap()
    symbols := make([]string,len(s)) 
    
    //Split into Symbols.
    for i,v := range s{
        symbols[i] = string(v)
    }
    
    //IF A Roman Int symbol 
    if len(s) == 1 {
        return RomanIntMap[s]
    }    
    
    //Compare Current RomanInt condition with Next one.
    //(IF) Current one is less than next one, The rev += current RomanInt - next RomanInt
    //(ELSE) The rev += current RomanInt
    for current := 0; current < len(s); current++{
        curSymbol = symbols[current]
        next = current + 1     
        if next == len(s)  {
            nextSymbol = symbols[current]
        } else {
            nextSymbol = symbols[next]
        }
        
        if RomanIntMap[curSymbol] < RomanIntMap[nextSymbol]  {
            num = RomanIntMap[nextSymbol] - RomanIntMap[curSymbol]
            rev += num
            /*fmt.Printf("Num: %d, Rev: %d\n",num,rev)*/
            if next == len(s) {
                break       
            }
            current++
        }else{
            num = RomanIntMap[curSymbol]
            rev += num
            /*fmt.Printf("Num: %d, Rev: %d\n",num,rev)*/
        }
    }
    return rev
}

//Prepare Hash Map
func InitMap() map[string]int{
    RomanIntMap := make(map[string]int,7)
    RomanIntMap["I"] = 1
    RomanIntMap["V"] = 5    
    RomanIntMap["X"] = 10    
    RomanIntMap["L"] = 50    
    RomanIntMap["C"] = 100    
    RomanIntMap["D"] = 500    
    RomanIntMap["M"] = 1000
    return RomanIntMap
}



