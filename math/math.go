package math
import "fmt"


func Add(a,b int) int {
	return a+b
}


func Subtract(a, b int) int {
	return a-b
}


func Divide(a, b int)( int , error){
	if b == 0{
		return 0, fmt.Errorf("division by zero")
	}

	return a / b , nil
}


func IsPositive(a int ) bool{
	 return a > 0
}

func Multiply(a, b int) int{
	return a * b
}