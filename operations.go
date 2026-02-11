// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import (
    "fmt"
    "errors"
)

func stat(nums []int)(int, float64, error){
    if len(nums) == 0{
        return 0,0, errors.New("empty array")
    }
    sum := 0
    for _, value := range nums {
        sum += value
    }
    avg := float64(sum) / float64(len(nums))
    return sum, avg, nil
}
 
func main(){
    numbers := []int{1,2,3,4,5,6,7,8,9}
    
    sum, avg, err := stat(numbers)
    if err != nil{
        fmt.Println("Error: ",err)
        return
    }
    fmt.Println("Sum: ", sum)
    fmt.Println("Average: ", avg)
}
