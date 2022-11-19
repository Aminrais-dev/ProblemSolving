package leetcodealgorithms

/*
Given an array of integers, determine the minimum number of elements to delete to leave only elements of equal value.
*/
func equalizeArray(arr []int32) int32 {

	var max int32
	var in int32
	for i := 0; i < len(arr); i++ {
		var temp = int32(0)
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp++
			}
		}
		if temp >= max {
			max = temp
			in = arr[i]
		}
	}

	var hasil []int32
	for k := 0; k < len(arr); k++ {
		if arr[k] != in {
			hasil = append(hasil, arr[k])
		}
	}

	return int32(len(hasil))

}
