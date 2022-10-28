package main

import "fmt"

/*
Temukan number yang terduplikasi dalam array! Bilangan dianggap terduplikasi jika muncul >1
*/
func findDuplicationNumber(numbers []int) {
	maps := map[int]int{}

	for _, v := range numbers {
		maps[v]++
	}

	var res []int
	for key, value := range maps {
		if value > 1 {
			res = append(res, key)
		}
	}

	fmt.Println(res)

}

func main() {
	findDuplicationNumber([]int{1, 1})                            //[1]
	findDuplicationNumber([]int{1, 2, 3, 4})                      //[]
	findDuplicationNumber([]int{1, 2, 1, 2, 2, 3, 4, 5, 5, 5, 5}) //[1,2,5]
}
