package maps

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
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

func Values2Map(values url.Values) map[string]interface{} {
	mapData := make(map[string]interface{})

	for k, v := range values {
		if len(v) > 0 {
			mapData[k] = v[0]
		} else {
			mapData[k] = ""
		}
	}

	return mapData
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

	for k := range *inputs {
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
			case reflect.Int64:
				buf.WriteString(strconv.FormatInt(v.(int64), 10))
			case reflect.Float64:
				buf.WriteString(strconv.Itoa(int(v.(float64))))
			case reflect.String:
				buf.WriteString(v.(string))
			}
		}
	}
	return buf.String()
}

func Map2Values(inputs *map[string]interface{}) url.Values {

	values := url.Values{}

	for k, v := range *inputs {
		rt := reflect.TypeOf(v)

		var value string

		switch rt.Kind() {
		case reflect.Int:
			value = strconv.Itoa(v.(int))
		case reflect.Int64:
			value = strconv.FormatInt(v.(int64), 10)
		case reflect.Float64:
			value = strconv.Itoa(int(v.(float64)))
		case reflect.String:
			value = v.(string)
		}
		values.Add(k, value)
	}

	return values
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
func (s *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*s = StringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*s)[e.XMLName.Local] = e.Value
	}
	return nil
}
