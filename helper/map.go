package helper

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func RemoveKeys(inputs *map[string]string, keys ...string) {
}

func KSort(inputs *map[string]interface{}) {

	var keys []string

	values := make(map[string]interface{})

	for k, value := range *inputs {
		keys = append(keys, k)
		values[k] = value
		delete(*inputs, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		(*inputs)[k] = values[k]
	}

	return
}

func CreateLinkString(inputs *map[string]interface{}) string {

	var buf bytes.Buffer

	var keys []string

	for k, _ := range *inputs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {

		if k != "sign" && k != "paySign" {

			prefix := k + "="

			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			v := (*inputs)[k]

			rt := reflect.TypeOf(v)

			buf.WriteString(prefix)

			switch rt.Kind() {
			case reflect.Int:
				buf.WriteString(strconv.Itoa(v.(int)))
			case reflect.Float64:
				buf.WriteString(strconv.Itoa(int(v.(float64))))
			case reflect.String:
				buf.WriteString(v.(string))
			}
		}
	}

	return buf.String()
}

func ToXml(values interface{}) string {

	b, err := xml.Marshal(values)

	fmt.Println(err)

	fmt.Println(string(b))

	//xml := "<xml>"
	//
	//for k,v := range *values{
	//	if reflect.TypeOf(v).Kind() == reflect.Int{
	//		xml+="/<"+k+"/>"+v+"/<"+k+"/>"
	//	}
	//}

	return ""

}

func MarshalXML(values *map[string]interface{}, b *bytes.Buffer, start xml.StartElement) error {

	e := xml.NewEncoder(b)

	tokens := []xml.Token{start}

	for key, value := range *values {
		t := xml.StartElement{Name: xml.Name{"", key}}

		tokens = append(tokens, t, xml.CharData(value.(string)), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

type StringMap map[string]string

// StringMap marshals into XML.
func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}

	for key, value := range s {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func Struct2Map(obj interface{}) map[string]interface{} {

	b, err := json.Marshal(obj)

	m := make(map[string]interface{})

	if err != nil {
		return m
	}

	err = json.Unmarshal(b, &m)

	return m
}
