package controller

import (
	"calculate/service"

	"github.com/gin-gonic/gin"
)

func convertString(v1 []string, v2 []string) ([]string, []string) {
	result1 := unionString(v1, v2)
	result2 := intersectString(v1, v2)
	return result1, result2
}

type ArrayDataString struct {
	Value1 []string `json:"value1"`
	Value2 []string `json:"value2"`
}

func ConvertValueString() func(c *gin.Context) {
	return func(c *gin.Context) {
		var arrayData ArrayDataString

		if err := c.BindJSON(&arrayData); err != nil {
			return
		}

		resultType1, resultType2 := convertString(arrayData.Value1, arrayData.Value2)

		result := ArrayDataString{Value1: resultType1, Value2: resultType2}
		service.Resp(c, 1, "OK", result)
	}
}

func intersectString(first, second []string) []string {
	out := []string{}
	bucket := map[string]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
}

func unionString(myslice1, myslice2 []string) []string {
	// Create a map to store the elements of the union
	values := make(map[string]bool)
	for _, key := range myslice1 { // for loop used in slice1 to remove duplicates from the values
		values[key] = true
	}
	for _, key := range myslice2 { // for loop used in slice2 to remove duplicates from the values
		values[key] = true
	}
	// Convert the map keys to a sliceq5
	output := make([]string, 0, len(values)) //create slice output
	for val := range values {
		output = append(output, val) //append values in slice output
	}
	return output
}
