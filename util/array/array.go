package array

import (
	"reflect"
)

/**
返回输入数组中某个单一列的值
*/
func ArrayColumn(slice interface{}, key string) []interface{} {

	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)

	column := make([]interface{}, 0, rt.Size())

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {

		for i := 0; i < rv.Len(); i++ {

			value := rv.Index(i)

			kind := value.Kind()

			var keyValue reflect.Value

			switch kind {
			case reflect.Map:

				keyValue = value.MapIndex(reflect.ValueOf(key))

			case reflect.Struct:

				keyValue = value.FieldByName(key)

			case reflect.Ptr:

				keyValue = value.Elem().FieldByName(key)

			}

			if keyValue.IsValid() {
				column = append(column, keyValue.Interface())
			}
		}
	}
	return column
}

func ArrayMap(slice interface{}, from, to string) map[interface{}]interface{} {
	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)

	column := make(map[interface{}]interface{})

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {

		for i := 0; i < rv.Len(); i++ {

			value := rv.Index(i)

			kind := value.Kind()

			var fromValue reflect.Value
			var toValue reflect.Value

			switch kind {
			case reflect.Map:

				fromValue = value.MapIndex(reflect.ValueOf(from))
				toValue = value.MapIndex(reflect.ValueOf(to))

			case reflect.Struct:

				fromValue = value.FieldByName(from)
				toValue = value.FieldByName(to)

			case reflect.Ptr:

				fromValue = value.Elem().FieldByName(from)
				toValue = value.Elem().FieldByName(to)

			}

			if fromValue.IsValid() && toValue.IsValid() {

				column[fromValue.Interface()] = toValue.Interface()

			}
		}
	}
	return column
}

func ArrayGroup(slice interface{}, key string) map[interface{}]interface{} {

	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)

	group := make(map[interface{}]interface{})

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {
		for i := 0; i < rv.Len(); i++ {

			value := rv.Index(i)

			kind := value.Kind()

			var keyValue reflect.Value

			switch kind {
			case reflect.Map:

				keyValue = value.MapIndex(reflect.ValueOf(key))

			case reflect.Struct:

				keyValue = value.FieldByName(key)

			case reflect.Ptr:

				keyValue = value.Elem().FieldByName(key)

			}

			if keyValue.IsValid() {

				if groupValue, ok := group[keyValue.Interface()]; !ok {

					slice := reflect.MakeSlice(rv.Type(), 0, 0)

					group[keyValue.Interface()] = reflect.Append(slice, value).Interface()

				} else {

					group[keyValue.Interface()] = reflect.Append(reflect.ValueOf(groupValue), value).Interface()

				}
			}
		}
	}

	return group

}

/**
移除数组中的重复的值
*/
func ArrayUnique(slice interface{}) []interface{} {

	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)

	m := make(map[interface{}]bool, rv.Len())

	arr := make([]interface{}, 0, rv.Len())

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {

		for i := 0; i < rv.Len(); i++ {

			value := rv.Index(i).Interface()

			if _, ok := m[value]; !ok {
				m[value] = true

				arr = append(arr, value)
			}
		}
	}
	return arr
}

/**
把一个或多个数组合并为一个数组
*/
func ArrayMerge(slices ...[]string) []string {
	s := make([]string, 0, len(slices)*5)

	for _, slice := range slices {
		s = append(s, slice...)
	}
	return s
}

/**
用于比较两个（或更多个）数组的值，并返回交集
*/
func ArrayIntersect(slices ...[]string) []interface{} {

	return nil
}
func ArrayIntersect2(slice1 []string, slice2 []string) []string {

	m1 := make(map[string]bool, len(slice1))
	m2 := make(map[string]bool, len(slice1))
	arr := make([]string, 0, len(slice1))

	for _, value1 := range slice1 {
		m1[value1] = true
	}

	for _, value2 := range slice2 {
		if _, ok := m1[value2]; ok {
			if _, ok := m2[value2]; !ok {
				m2[value2] = true
				arr = append(arr, value2)
			}
		}
	}
	return arr
}

/**
搜索数组中是否存在指定的值
*/
func InArray(search interface{}, slice interface{}) bool {
	return ArraySearch(search, slice) != -1
}

/**
在数组中搜索给定的值，如果成功则返回首个相应的index
*/
func ArraySearch(search interface{}, slice interface{}) int {

	rt := reflect.TypeOf(slice)
	rv := reflect.ValueOf(slice)

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {

		for i := 0; i < rv.Len(); i++ {

			value := rv.Index(i).Interface()

			if value == search {
				return i
			}
		}

		return -1
	} else {
		return -1
	}
}

func ArrayPush(slice []int, value int) {

	slice = append(slice, value)

}

func ArrayUnshift(slice []int, value int) {

	s := []int{}

	s = append(s, value)

	s = append(s, slice...)
}
