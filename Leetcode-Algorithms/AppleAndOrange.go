package leetcodealgorithms

import "fmt"

/*
Sam's house has an apple tree and an orange tree that yield an abundance of fruit.
Using the information given below, determine the number of apples and oranges that land on Sam's house.
*/

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var ressApple int
	var ressOrange int

	for i := 0; i < len(apples); i++ {
		if apples[i]+a >= s && apples[i]+a <= t {
			ressApple++
		}
	}

	for k := 0; k < len(oranges); k++ {
		if oranges[k]+b >= s && oranges[k]+b <= t {
			ressOrange++
		}
	}

	fmt.Println(ressApple)
	fmt.Println(ressOrange)
}
