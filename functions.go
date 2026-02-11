// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"


func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
    res, err := divide(10, 2)
    if err != nil {
    	fmt.Println(err)
    	return
    }
    fmt.Println(res)
}
