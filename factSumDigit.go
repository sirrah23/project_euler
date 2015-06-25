package main

import "fmt"
import "math/big"

func factSumDigit(n int64) *big.Int{
	factN := new(big.Int)
	/*Compute n!*/
	factN = factN.MulRange(1,n)
	/*Sum the digits*/
	ret := sumDigit(factN)
	/*Return the answer*/
	return ret
}

func sumDigit(n *big.Int) *big.Int{
	result := big.NewInt(0)
	zero := big.NewInt(0)
	for n.Cmp(zero) > 0{
		temp := new(big.Int)
		result = result.Add(result,temp.Mod(n,big.NewInt(10)))
		n = n.Div(n,big.NewInt(10))
	}
	return result
}



func main() {
	y := factSumDigit(100)
	fmt.Println(y)
}