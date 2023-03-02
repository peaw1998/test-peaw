package controller

import (
	"calculate/service"

	"github.com/gin-gonic/gin"
)

func convert(v1 []int, v2 []int) ([]int, []int) {
	result1 := intersect(v1, v2)
	result2 := union(v1, v2)
	return result1, result2
}

type ArrayData struct {
	Value1 []int `json:"value1"`
	Value2 []int `json:"value2"`
}

func ConvertValue() func(c *gin.Context) {
	return func(c *gin.Context) {
		var arrayData ArrayData

		if err := c.BindJSON(&arrayData); err != nil {
			return
		}

		resultType1, resultType2 := convert(arrayData.Value1, arrayData.Value2)

		result := ArrayData{Value1: resultType1, Value2: resultType2}
		service.Resp(c, 1, "OK", result)
	}
}

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	// identify the smaller array
	smaller, larger := nums1, nums2
	if len(nums1) > len(nums2) {
		smaller, larger = nums2, nums1
	}

	counts := make(map[int]int, len(smaller))
	// iterate through the smaller array to store each value and its count
	for _, num := range smaller {
		_, ok := counts[num]
		if ok {
			counts[num] += 1
		} else {
			counts[num] = 1
		}
	}

	// iterate through the larger array, if the value's count is nonzero, decrement the count and add to result set
	for _, num := range larger {
		count, ok := counts[num]
		if ok && count != 0 {
			result = append(result, num)
			counts[num] -= 1
		}
	}
	return result
}

func union(myslice1, myslice2 []int) []int {
	// Create a map to store the elements of the union
	values := make(map[int]bool)
	for _, key := range myslice1 { // for loop used in slice1 to remove duplicates from the values
		values[key] = true
	}
	for _, key := range myslice2 { // for loop used in slice2 to remove duplicates from the values
		values[key] = true
	}
	// Convert the map keys to a sliceq5
	output := make([]int, 0, len(values)) //create slice output
	for val := range values {
		output = append(output, val) //append values in slice output
	}
	return output
}
