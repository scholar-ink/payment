package array

import (
	"log"
	"strconv"
	"testing"
)

func TestArrayColumn(t *testing.T) {

	slice := [5]map[string]string{
		{
			"Name": "111",
		},
	}

	//slice = append(slice, map[string]string{"name":"1111"})

	type User struct {
		Name string
	}

	//slice := make([]*User,0)
	//
	//slice = append(slice, &User{Name:"1111"})

	ArrayColumn(slice, "Name")

	//log.Println(arr)
}

func TestArrayUnique(t *testing.T) {

	arr := []string{
		"zhouchao",
		"zhouchao2",
		"zhouchao",
	}

	arr2 := []string{
		"zhouchao",
	}
	arr3 := []string{
		"zhouchao2",
	}
	arr4 := []string{
		"zhouchao4",
	}
	//
	//arr := []int{
	//	1,2,2,4,5,3,2,2,1,3,6,
	//}

	log.Println(ArrayIntersect(arr, arr2, arr3, arr4))
}

func TestArraySearch(t *testing.T) {

	log.Println(ArraySearch(5, []int{1, 2, 3, 4, 5}))
	log.Println(ArraySearch("zhouchao3", []string{
		"zhouchao",
		"zhouchao2",
		"zhouchao",
	}))

}

func TestInArray(t *testing.T) {
	log.Println(InArray(5, []int{1, 2, 3, 4, 5}))
}

func TestArrayPush(t *testing.T) {

	ArrayPush([]int{1, 2, 3, 4}, 1)

}

func TestArrayUnshift(t *testing.T) {
	ArrayUnshift([]int{1, 2, 3, 4}, 1)
}

func TestArrayMap(t *testing.T) {

	slice := [5]map[string]string{
		{
			"name": "111",
			"age":  "1",
		},
		{
			"name": "2222",
			"age":  "2",
		},
	}
	mapData := ArrayMap(slice, "name", "age")

	log.Println(mapData)
}

func TestArrayGroup(t *testing.T) {

	//DeliveryRange := 10
	StartFeeFloat := 10.02222

	log.Println(strconv.FormatFloat(StartFeeFloat, 'G', -2, 64))

}
